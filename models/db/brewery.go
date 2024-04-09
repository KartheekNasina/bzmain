package models

import "time"

// Breweries represents the structure for breweries table
type Brewery struct {
	ID               string     `db:"id"`
	Name             string     `db:"name"`
	LogoURL          string     `db:"logo_url"`
	ThumbnailURL     string     `db:"thumbnail_url"`
	ShortDescription string     `db:"short_description"`
	Description      string     `db:"description"`
	Images           []string   `db:"images"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at"`
	Geolocation      string     `db:"geolocation"`
	Address          string     `db:"address"`
	Lat              float64    `db:"lat"`
	Lng              float64    `db:"lng"`
	CountryID        string     `db:"country_id"`
	CityID           string     `db:"city_id"`
	PrimaryEmail     string     `db:"primary_email"`
	SecondaryEmail   string     `db:"secondary_email"`
	PublicPhone      string     `db:"public_phone"`
	SecondaryPhone   string     `db:"secondary_phone"`
	OwnerName        string     `db:"owner_name"`
	OwnerPhone       string     `db:"owner_phone"`
	FeaturedPriority int        `db:"featured_priority"`
	TrendingPriority int        `db:"trending_priority"`
	PetFriendly      bool       `db:"pet_friendly"`
	IsNew            bool       `db:"is_new"`
	ComingSoon       bool       `db:"coming_soon"`
	WorkFromBrewery  bool       `db:"work_from_brewery"`
	IsPublished      bool       `db:"is_published"`
	PublishedAt      *time.Time `db:"updated_at"`
	SundayBrunch     bool       `db:"sunday_brunch"`
}
