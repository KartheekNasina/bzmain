
CREATE TYPE onboarding_status_type AS ENUM ('started', 'otp_verified', 'profile_completed', 'onboarded');


CREATE TABLE users (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id text NOT NULL UNIQUE,
    name text NOT NULL,
    about_me text,
    gender text NOT NULL REFERENCES gender(type) ON DELETE RESTRICT ON UPDATE RESTRICT,
    email text NOT NULL UNIQUE,
    profile_url text,
	images text[],
    dob date, -- Date of birth
    is_online boolean NOT NULL DEFAULT false,
    allow_notifications boolean NOT NULL DEFAULT true, -- Allow notifications check
    allow_location boolean NOT NULL DEFAULT true, -- Allow location check
    phone_number text,
    is_legal_age boolean NOT NULL DEFAULT false,
    provider text NOT NULL REFERENCES auth_providers(provider) ON DELETE RESTRICT ON UPDATE RESTRICT,
    brew_interests text[], -- Brewing related interests
    personal_interests text[], -- Personal interests
	phone_number_verified boolean DEFAULT false,
	phone_number_verification_date timestamp with time zone,
	onboarding_status onboarding_status_type DEFAULT 'started',
    otp_verified_at timestamp with time zone,
    profile_completed_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

-- Basic Brewery Information
CREATE TABLE brewery_basic_info (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name text NOT NULL UNIQUE,
    logo_url text NOT NULL,
    thumbnail_url text NOT NULL,
    short_description text NOT NULL,
    description text NOT NULL,
    images text[],
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

-- Brewery Location
CREATE TABLE brewery_location (
    brewery_id uuid REFERENCES brewery_basic_info(id) ON DELETE CASCADE,
    geolocation geography(Point,4326) NOT NULL,
    address text NOT NULL,
	lat double precision NOT NULL CHECK (lat >= -90 AND lat <= 90),
    lng double precision NOT NULL CHECK (lng >= -180 AND lng <= 180),
    country_id uuid REFERENCES countries(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    city_id uuid REFERENCES cities(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    PRIMARY KEY (brewery_id)
);

-- Brewery Contacts
CREATE TABLE brewery_contacts (
    brewery_id uuid REFERENCES brewery_basic_info(id) ON DELETE CASCADE,
    primary_email text NOT NULL,
    secondary_email text,
    public_phone text,
    secondary_phone text,
    owner_name text,
    owner_phone text,
    PRIMARY KEY (brewery_id)
);

-- Brewery Metadata
CREATE TABLE brewery_metadata (
    brewery_id uuid REFERENCES brewery_basic_info(id) ON DELETE CASCADE,
    featured_priority integer DEFAULT 0,
    trending_priority integer DEFAULT 0,
    pet_friendly boolean DEFAULT false,
    is_new boolean DEFAULT false,
    coming_soon boolean DEFAULT false,
    work_from_brewery boolean DEFAULT false,
    is_published boolean NOT NULL DEFAULT false,
    published_at timestamp with time zone,
    sunday_brunch boolean NOT NULL DEFAULT false,
    PRIMARY KEY (brewery_id)
);


-- Countries Table
CREATE TABLE countries (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name text NOT NULL UNIQUE,
    code text NOT NULL UNIQUE, -- ISO country code, e.g., 'US' for the United States
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

-- Cities Table
CREATE TABLE cities (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name text NOT NULL,
    country_id uuid REFERENCES countries(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    UNIQUE(name, country_id) -- Ensures city names are unique within a country
);


-- Beer Types Table
CREATE TABLE beer_types (
    type text PRIMARY KEY,
    description text
    -- Add comments if needed
);

-- Beers Table
CREATE TABLE beers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name text NOT NULL,
    type text NOT NULL REFERENCES beer_types(type) ON DELETE RESTRICT ON UPDATE RESTRICT,
    image_url text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    is_published boolean NOT NULL DEFAULT false,
    published_at timestamp with time zone,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    abv double precision CHECK (abv >= 0 AND abv <= 100), -- Assuming a percentage representation
    ibu double precision CHECK (ibu >= 0),
    description text,
    UNIQUE(name, brewery_id) -- Ensuring beer names are unique within a brewery
    -- Add comments if needed
);


CREATE TABLE brewery_review (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    rating integer NOT NULL CHECK (rating >= 1 AND rating <= 5), -- Assuming a 1-5 rating scale
    user_id uuid REFERENCES users(id) ON DELETE SET NULL, -- Assuming a review can exist even if a user is deleted
    title text NOT NULL,
    description text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    UNIQUE(brewery_id, user_id) -- Ensuring one review per brewery per user
    -- Add comments if needed
);

-- Index for faster lookups based on brewery and user
CREATE INDEX idx_brewery_review_brewery_id ON brewery_review(brewery_id);
CREATE INDEX idx_brewery_review_user_id ON brewery_review(user_id);


-- Event Basics Table
CREATE TABLE event_basic_info (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    title text NOT NULL,
    short_description text NOT NULL,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    thumbnail text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    is_published boolean NOT NULL DEFAULT false,
    published_at timestamp with time zone,
    created_by_id uuid REFERENCES users(id) ON DELETE SET NULL,
    updated_by_id uuid REFERENCES users(id) ON DELETE SET NULL,
    images text[] NOT NULL
);

CREATE TYPE repeat_pattern AS ENUM (
    'none', 
    'daily', 
    'weekday', 
    'weekend', 
    'weekly_mon', 
    'weekly_tue',
    'weekly_wed',
    'weekly_thu',
    'weekly_fri',
    'weekly_sat',
    'weekly_sun',
    'monthly',
    'yearly',
    'custom'
);

CREATE TABLE event_schedule (
    event_id uuid REFERENCES event_basic_info(id) ON DELETE CASCADE,
    start_date date NOT NULL,
    end_date date,
    start_time time with time zone NOT NULL,
    end_time time with time zone NOT NULL,
    repeats repeat_pattern NOT NULL DEFAULT 'none',
    max_people_count integer NOT NULL CHECK (max_people_count > 0),
    PRIMARY KEY (event_id)
);

-- Event Join Requests Table
CREATE TYPE request_status AS ENUM ('requested', 'approved', 'declined', 'confirmed', 'canceled');

CREATE TABLE event_join_requests (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    event_id uuid NOT NULL REFERENCES event_basic_info(id) ON DELETE CASCADE,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status request_status NOT NULL DEFAULT 'requested',
    requested_at timestamp with time zone DEFAULT now(),
    approved_at timestamp with time zone,
    declined_at timestamp with time zone,
    confirmed_at timestamp with time zone,
    canceled_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    -- Additional fields for any notes or reasons for decisions can be added
    notes text,
    -- Ensuring a user can't have multiple requests for the same event
    UNIQUE(event_id, user_id)
);

-- Brew Tours Basic Information
CREATE TABLE brew_tours_basic_info (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    title text NOT NULL,
    short_description text NOT NULL,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    thumbnail text NOT NULL,
    is_published boolean NOT NULL DEFAULT false,
    published_at timestamp with time zone,
    images text[],
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    created_by_id uuid REFERENCES users(id) ON DELETE SET NULL,
    updated_by_id uuid REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE brew_tours_schedule (
    tour_id uuid REFERENCES brew_tours_basic_info(id) ON DELETE CASCADE,
    start_date date NOT NULL,
    end_date date,
    start_time time with time zone NOT NULL,
    end_time time with time zone NOT NULL,
    repeats repeat_pattern NOT NULL DEFAULT 'none',
    max_people_count integer NOT NULL CHECK (max_people_count > 0),
    PRIMARY KEY (tour_id)
);

-- -- Payments Table
-- CREATE TYPE payment_status AS ENUM ('pending', 'completed', 'failed', 'refunded');

-- CREATE TABLE payments (
--     id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
--     entity_type text NOT NULL, -- 'event' or 'brew_tour'
--     entity_id uuid NOT NULL, -- Can be event_id or brew_tour_id
--     user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
--     amount numeric(10,2) NOT NULL,
--     currency_code text NOT NULL,
--     payment_status payment_status NOT NULL DEFAULT 'pending',
--     payment_method text NOT NULL,
--     transaction_id text,
--     payment_date timestamp with time zone,
--     refund_date timestamp with time zone,
--     failure_reason text,
--     notes text,
--     UNIQUE(entity_type, entity_id, user_id) -- Ensuring unique payment per entity per user
-- );

-- Brew Classes Basic Information
CREATE TABLE brew_classes_basic_info (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    title text NOT NULL,
    short_description text NOT NULL,
    brewery_id uuid REFERENCES breweries(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    thumbnail text NOT NULL,
    is_published boolean NOT NULL DEFAULT false,
    published_at timestamp with time zone,
    images text[],
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    created_by_id uuid REFERENCES users(id) ON DELETE SET NULL,
    updated_by_id uuid REFERENCES users(id) ON DELETE SET NULL,
    address text NOT NULL,
    lat double precision,
    lng double precision,
    cost double precision NOT NULL DEFAULT '0'::double precision,
    discount_perc integer NOT NULL DEFAULT 0
);


CREATE TABLE brew_classes_schedule (
    class_id uuid REFERENCES brew_classes_basic_info(id) ON DELETE CASCADE,
    start_date date NOT NULL,
    end_date date,
    start_time time with time zone NOT NULL,
    end_time time with time zone NOT NULL,
    repeats repeat_pattern NOT NULL DEFAULT 'none',
    max_people_count integer NOT NULL CHECK (max_people_count > 0),
    PRIMARY KEY (class_id)
);



-- Community Categories Table
CREATE TABLE community_categories (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    type text NOT NULL UNIQUE,
    thumbnail text NOT NULL,
    title text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    created_by_id uuid REFERENCES users(id) ON DELETE SET NULL,
    updated_by_id uuid REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE community_drives (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    title text NOT NULL,
    community_type text NOT NULL REFERENCES community_categories(type) ON DELETE RESTRICT ON UPDATE RESTRICT,
    description text NOT NULL,
    address text NOT NULL,
    lat double precision NOT NULL CHECK (lat >= -90 AND lat <= 90),
    lng double precision NOT NULL CHECK (lng >= -180 AND lng <= 180),
    contact_name text NOT NULL,
    contact_phone_number bigint NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    is_published boolean NOT NULL DEFAULT true,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    published_at timestamp with time zone,
    thumbnail text NOT NULL,
    images text[]
);

CREATE TABLE community_centers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    title text NOT NULL,
    community_type text NOT NULL REFERENCES community_categories(type) ON DELETE RESTRICT ON UPDATE RESTRICT,
    description text NOT NULL,
    address text NOT NULL,
    lat double precision NOT NULL CHECK (lat >= -90 AND lat <= 90),
    lng double precision NOT NULL CHECK (lng >= -180 AND lng <= 180),
    contact_name text NOT NULL,
    contact_phone_number bigint NOT NULL,
    is_published boolean NOT NULL DEFAULT false,
    published_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    thumbnail text NOT NULL,
    images text[]
);

CREATE TYPE registration_status AS ENUM ('pending', 'confirmed', 'waitlisted', 'canceled');

CREATE TABLE user_community_drive_registrations (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    drive_id uuid NOT NULL REFERENCES community_drives(id) ON DELETE CASCADE,
    status registration_status NOT NULL DEFAULT 'pending',
    registration_date timestamp with time zone DEFAULT now(),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    notes text, -- Any additional notes or details about the registration
    UNIQUE(user_id, drive_id) -- Ensuring a user can't register for the same drive multiple times
);


CREATE TYPE class_registration_status AS ENUM ('pending', 'confirmed', 'waitlisted', 'canceled');

CREATE TABLE user_brew_class_registrations (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    class_id uuid NOT NULL REFERENCES brew_classes_basic_info(id) ON DELETE CASCADE,
    status class_registration_status NOT NULL DEFAULT 'pending',
    registration_date timestamp with time zone DEFAULT now(),
    number_of_attendees integer NOT NULL CHECK (number_of_attendees > 0),
    notes text, -- Any additional notes or details about the registration
    UNIQUE(user_id, class_id) -- Ensuring a user can't register for the same class multiple times
);



-- For Event Registrations
CREATE TYPE event_registration_status AS ENUM ('pending', 'confirmed', 'waitlisted', 'canceled');

CREATE TABLE user_event_registrations (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    event_id uuid NOT NULL REFERENCES event_basic_info(id) ON DELETE CASCADE,
    status event_registration_status NOT NULL DEFAULT 'pending',
    registration_date timestamp with time zone DEFAULT now(),
    number_of_people integer NOT NULL CHECK (number_of_people > 0),
    notes text, -- Additional notes or details about the registration
    UNIQUE(user_id, event_id) -- Ensuring a user can't register for the same event multiple times
);

-- For Brew Tour Registrations
CREATE TYPE tour_registration_status AS ENUM ('pending', 'confirmed', 'waitlisted', 'canceled');

CREATE TABLE user_brew_tour_registrations (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    tour_id uuid NOT NULL REFERENCES brew_tours_basic_info(id) ON DELETE CASCADE,
    status tour_registration_status NOT NULL DEFAULT 'pending',
    registration_date timestamp with time zone DEFAULT now(),
    number_of_people integer NOT NULL CHECK (number_of_people > 0),
    notes text, -- Additional notes or details about the registration
    UNIQUE(user_id, tour_id) -- Ensuring a user can't register for the same tour multiple times
);


CREATE TYPE offer_type AS ENUM ('featured', 'exclusive', 'top', 'best');
CREATE TYPE offer_status AS ENUM ('active', 'inactive', 'expired');

CREATE TABLE brewery_offers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE CASCADE,
    offer_type offer_type NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    thumbnail text,
    images text[],
    start_date date NOT NULL,
    end_date date,
    start_time time with time zone NOT NULL,
    end_time time with time zone NOT NULL,
    repeats repeat_pattern NOT NULL DEFAULT 'none',
    status offer_status NOT NULL DEFAULT 'active',
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    created_by_id uuid REFERENCES users(id) ON DELETE SET NULL,
    updated_by_id uuid REFERENCES users(id) ON DELETE SET NULL
);

CREATE TYPE user_offer_status AS ENUM ('active', 'redeemed', 'expired');

CREATE TABLE user_exclusive_offers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE CASCADE,
    title text NOT NULL,
    description text NOT NULL,
    thumbnail text,
    start_date date NOT NULL,
    end_date date,
    start_time time with time zone NOT NULL,
    end_time time with time zone NOT NULL,
    repeats repeat_pattern NOT NULL DEFAULT 'none',
    status user_offer_status NOT NULL DEFAULT 'active',
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    created_by_id uuid REFERENCES users(id) ON DELETE SET NULL,
    updated_by_id uuid REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE user_offer_redemptions (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    offer_id uuid NOT NULL REFERENCES user_exclusive_offers(id) ON DELETE CASCADE,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    redemption_date timestamp with time zone DEFAULT now(),
    notes text, 
	created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

CREATE TABLE user_brewery_visits (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE CASCADE,
    visit_date date NOT NULL DEFAULT CURRENT_DATE,
    visit_time timestamp with time zone DEFAULT now()
);

CREATE TABLE monthly_leaderboard (
    month_year date NOT NULL, -- Represents the month and year of the leaderboard
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    visits_count integer NOT NULL,
    rank integer NOT NULL,
    title text, -- "King", "Queen", etc.
    UNIQUE(month_year, user_id)
);

-- Optional: Indexes for faster lookups
CREATE INDEX idx_monthly_leaderboard_month ON monthly_leaderboard(month_year);
CREATE INDEX idx_monthly_leaderboard_user ON monthly_leaderboard(user_id);





WITH MonthlyVisits AS (
    SELECT 
        u.user_id, 
        u.gender,
        COUNT(*) as visits_count
    FROM 
        user_brewery_visits ubv
    JOIN 
        users u ON ubv.user_id = u.id
    WHERE 
        EXTRACT(MONTH FROM ubv.visit_date) = EXTRACT(MONTH FROM CURRENT_DATE)
        AND EXTRACT(YEAR FROM ubv.visit_date) = EXTRACT(YEAR FROM CURRENT_DATE)
    GROUP BY 
        u.user_id, u.gender
    ORDER BY 
        visits_count DESC
)

, GenderRanks AS (
    SELECT
        user_id,
        gender,
        visits_count,
        ROW_NUMBER() OVER (PARTITION BY gender ORDER BY visits_count DESC) as gender_rank
    FROM
        MonthlyVisits
)

INSERT INTO monthly_leaderboard (month_year, user_id, visits_count, rank, title)
SELECT 
    CURRENT_DATE - INTERVAL '1 month' + INTERVAL '1 day', -- Start of the month
    user_id, 
    visits_count, 
    gender_rank,
    CASE 
        WHEN gender = 'male' AND gender_rank = 1 THEN 'King'
        WHEN gender = 'female' AND gender_rank = 1 THEN 'Queen'
        ELSE NULL
    END as title
FROM 
    GenderRanks
WHERE
    gender_rank = 1;


CREATE TYPE referral_status AS ENUM ('pending', 'successful', 'expired');

CREATE TABLE user_referrals (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    referrer_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- the one who referred
    referee_id uuid REFERENCES users(id) ON DELETE SET NULL, -- the one who was referred
    referral_code text NOT NULL UNIQUE, -- unique code associated with this referral
    status referral_status NOT NULL DEFAULT 'pending',
    referred_at timestamp with time zone DEFAULT now(),
    completed_at timestamp with time zone, -- date when the referral was successful
    reward_claimed boolean DEFAULT false
);

CREATE TYPE interaction_type AS ENUM ('brewery_visit', 'event_attendance', 'tour_attendance', 'class_attendance');

CREATE TABLE user_history (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    interaction_id uuid NOT NULL, -- This will store the ID of brewery/event/tour/class
    interaction_type interaction_type NOT NULL,
    interaction_date timestamp with time zone DEFAULT now(),
    notes text,
    UNIQUE(user_id, interaction_id, interaction_type) -- To ensure no duplicate entries
);

CREATE TABLE user_favorite_breweries (
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE CASCADE,
    added_date timestamp with time zone DEFAULT now(),
    PRIMARY KEY(user_id, brewery_id) -- Composite primary key to ensure a unique combination of user and brewery
);

CREATE TABLE user_brewery_checkins (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE CASCADE,
    checkin_date date NOT NULL DEFAULT CURRENT_DATE,
    checkin_time timestamp with time zone DEFAULT now(),
    notes text,
    UNIQUE(user_id, brewery_id, checkin_date) -- To ensure no duplicate check-ins on the same day
);

CREATE TYPE feedback_type AS ENUM ('suggestion', 'issue', 'compliment', 'complaint', 'other');

CREATE TABLE user_feedback (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feedback_date timestamp with time zone DEFAULT now(),
    type feedback_type NOT NULL,
    subject text NOT NULL, -- Short summary of feedback
    description text, -- Detailed feedback description or comment
    is_resolved boolean DEFAULT false, -- To track if the feedback has been addressed or resolved
    resolved_at timestamp with time zone,
    resolution_notes text -- Notes or comments on resolution
);


CREATE TYPE meetup_status AS ENUM ('pending_approval', 'approved', 'rejected', 'cancelled');

CREATE TABLE meetups (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE CASCADE,
    organizer_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title text NOT NULL,
    description text,
    meetup_date date NOT NULL,
    start_time time with time zone NOT NULL,
    end_time time with time zone NOT NULL,
    max_attendees integer, -- Maximum number of attendees allowed; NULL if no limit
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    status meetup_status DEFAULT 'pending_approval',
    approved_at timestamp with time zone,
    rejected_at timestamp with time zone,
    cancellation_reason text
);


CREATE TABLE meetup_attendees (
    meetup_id uuid NOT NULL REFERENCES meetups(id) ON DELETE CASCADE,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    rsvp_date timestamp with time zone DEFAULT now(),
    attended boolean DEFAULT false, -- Mark if the user attended the meetup
    PRIMARY KEY(meetup_id, user_id) -- Composite primary key to ensure unique combination
);


CREATE TABLE user_connection_requests (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    sender_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    request_date timestamp with time zone DEFAULT now(),
    status ENUM('pending', 'accepted', 'rejected') DEFAULT 'pending',
    UNIQUE(sender_id, receiver_id) -- Ensure one request between two users
);


CREATE TABLE user_connections (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user1_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user2_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    connection_date timestamp with time zone DEFAULT now(),
    UNIQUE(user1_id, user2_id) -- Ensure unique connection pair
);

CREATE TABLE chat_conversations (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user1_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user2_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    last_message_id uuid REFERENCES chat_messages(id) ON DELETE SET NULL,
    UNIQUE(user1_id, user2_id) -- Ensure unique conversation pair
);

CREATE TABLE chat_messages (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    conversation_id uuid NOT NULL REFERENCES chat_conversations(id) ON DELETE CASCADE,
    sender_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message_text text NOT NULL,
    message_type ENUM('text', 'image', 'video', 'audio', 'file'),
    timestamp timestamp with time zone DEFAULT now()
);

CREATE TABLE message_status (
    message_id uuid NOT NULL REFERENCES chat_messages(id) ON DELETE CASCADE,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    delivered_at timestamp with time zone,
    read_at timestamp with time zone,
    PRIMARY KEY(message_id, user_id) -- Ensure unique status entry per message and user
);

CREATE TYPE payment_status AS ENUM ('created', 'captured', 'failed', 'refunded');

CREATE TABLE payment_transactions (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount numeric(10,2) NOT NULL, -- Amount in the smallest currency unit (e.g., cents, paise)
    currency_code char(3) NOT NULL, -- INR for Indian Rupee, USD for US Dollar, etc.
    razorpay_payment_id text NOT NULL, -- ID returned by Razorpay for the payment
    status payment_status DEFAULT 'created',
    order_description text, -- Description of what the payment was for
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    UNIQUE(razorpay_payment_id) -- Ensure uniqueness for each Razorpay payment ID
);

CREATE TABLE payment_refunds (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    payment_transaction_id uuid NOT NULL REFERENCES payment_transactions(id) ON DELETE CASCADE,
    razorpay_refund_id text NOT NULL, -- ID returned by Razorpay for the refund
    amount numeric(10,2) NOT NULL, -- Amount refunded
    reason text, -- Reason for the refund
    refund_date timestamp with time zone DEFAULT now(),
    UNIQUE(razorpay_refund_id) -- Ensure uniqueness for each Razorpay refund ID
);

-- Create ENUM types
CREATE TYPE payment_interaction_type AS ENUM ('request', 'response');
CREATE TYPE entity_association_type AS ENUM ('brew_class', 'event', 'tour');

-- Create the payment_logs table
CREATE TABLE payment_logs (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    payment_transaction_id uuid REFERENCES payment_transactions(id) ON DELETE SET NULL,
    interaction payment_interaction_type NOT NULL,
    endpoint text, -- Razorpay API endpoint interacted with
    payload jsonb, -- JSON payload sent or received
    timestamp timestamp with time zone DEFAULT now()
);

-- Create the payment_associations table
CREATE TABLE payment_associations (
    payment_transaction_id uuid NOT NULL REFERENCES payment_transactions(id) ON DELETE CASCADE,
    entity_type entity_association_type NOT NULL,
    entity_id uuid NOT NULL,
    PRIMARY KEY(payment_transaction_id, entity_type, entity_id)
);


CREATE TYPE subscription_type AS ENUM ('basic', 'premium', 'gold');

CREATE TABLE subscription_tiers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    type subscription_type NOT NULL UNIQUE,
    daily_message_limit integer, -- Number of messages a user can send daily
    unlimited_swipes boolean DEFAULT false, -- Unlimited likes/swipes feature
    see_read_receipts boolean DEFAULT false, -- Ability to see if messages are read
    price numeric(10,2), -- Price for this subscription tier, NULL for the basic tier
    duration_days integer -- Duration of the subscription (e.g., 30 for a month). NULL for basic
);

CREATE TABLE user_subscriptions (
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE PRIMARY KEY,
    subscription_tier_id uuid NOT NULL REFERENCES subscription_tiers(id) ON DELETE CASCADE,
    start_date timestamp with time zone DEFAULT now(),
    end_date timestamp with time zone,
    is_active boolean DEFAULT true
);

CREATE TABLE daily_message_counts (
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    date date NOT NULL DEFAULT current_date,
    message_count integer DEFAULT 0,
    PRIMARY KEY(user_id, date)
);

CREATE TABLE user_notifications (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    is_read boolean DEFAULT false,
    read_at timestamp with time zone
);

-- Create the ENUM type for roles
CREATE TYPE admin_role AS ENUM ('admin', 'moderator');

-- Create the admins table
CREATE TABLE admins (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name text NOT NULL,
    email text NOT NULL UNIQUE,
    password text NOT NULL, -- Added password column (consider hashing the passwords before storing)
    role admin_role NOT NULL,
    created_at timestamp with time zone DEFAULT now()
);


CREATE TABLE user_blocked_lists (
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    blocked_user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    blocked_at timestamp with time zone DEFAULT now(),
    PRIMARY KEY(user_id, blocked_user_id)
);

CREATE TABLE brewery_follows (
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE CASCADE,
    followed_at timestamp with time zone DEFAULT now(),
    PRIMARY KEY(user_id, brewery_id)
);

-- Create the ENUM type for item type
CREATE TYPE item_type AS ENUM ('food', 'drink');

-- Create the food_drink_items table
CREATE TABLE food_drink_items (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    brewery_id uuid NOT NULL REFERENCES breweries(id) ON DELETE CASCADE,
    name text NOT NULL,
    type item_type NOT NULL,
    description text,
    image_url text
);


CREATE TABLE predefined_comments (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    comment_text text NOT NULL
);

CREATE TABLE food_drink_ratings (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    item_id uuid NOT NULL REFERENCES food_drink_items(id) ON DELETE CASCADE,
    rating integer NOT NULL CHECK (rating BETWEEN 1 AND 5), -- Assuming a 5-star rating system
    comment_id uuid REFERENCES predefined_comments(id) ON DELETE SET NULL,
    rated_at timestamp with time zone DEFAULT now(),
    UNIQUE(user_id, item_id) -- Ensure a user can rate an item only once
);

CREATE TYPE reaction_type AS ENUM ('like', 'dislike', 'love', 'wow');

CREATE TABLE rating_reactions (
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    rating_id uuid NOT NULL REFERENCES food_drink_ratings(id) ON DELETE CASCADE,
    reaction reaction_type NOT NULL,
    reacted_at timestamp with time zone DEFAULT now(),
    PRIMARY KEY(user_id, rating_id) -- Ensure a user can react to a rating only once
);

CREATE TABLE predefined_drink_messages (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    message_text text NOT NULL UNIQUE
);

-- Create the ENUM type for request status
CREATE TYPE drink_request_status AS ENUM ('pending', 'accepted', 'declined');

-- Create the drink_purchase_requests table
CREATE TABLE drink_purchase_requests (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    buyer_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    recipient_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message_id uuid REFERENCES predefined_drink_messages(id) ON DELETE SET NULL,
    request_date timestamp with time zone DEFAULT now(),
    status drink_request_status DEFAULT 'pending',
    UNIQUE(buyer_id, recipient_id, request_date) -- Ensure a unique request per day
);


------- 

CREATE OR REPLACE VIEW brewery_details AS
SELECT
    bbi.id,
    bbi.name,
    bbi.logo_url,
    bbi.thumbnail_url,
    bbi.short_description,
    bbi.description,
    bbi.images,
    bbi.created_at,
    bbi.updated_at,
    bl.geolocation,
    bl.address,
    bl.lat,
    bl.lng,
    bl.country_id,
    bl.city_id,
    bc.primary_email,
    bc.secondary_email,
    bc.public_phone,
    bc.secondary_phone,
    bc.owner_name,
    bc.owner_phone,
    bm.featured_priority,
    bm.trending_priority,
    bm.pet_friendly,
    bm.is_new,
    bm.coming_soon,
    bm.work_from_brewery,
    bm.is_published,
    bm.published_at,
    bm.sunday_brunch
FROM
    brewery_basic_info bbi
JOIN brewery_location bl ON bbi.id = bl.brewery_id
JOIN brewery_contacts bc ON bbi.id = bc.brewery_id
JOIN brewery_metadata bm ON bbi.id = bm.brewery_id;
