package models

type BreweryDTO struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	LogoURL          string   `json:"logo_url"`
	ThumbnailURL     string   `json:"thumbnail_url"`
	ShortDescription string   `json:"short_description"`
	Description      string   `json:"description"`
	Images           []string `json:"images"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
	// Geolocation      string   `json:"geolocation"`
	Address   string  `json:"address"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	CountryID string  `json:"country_id"`
	CityID    string  `json:"city_id"`
	// PrimaryEmail     string  `json:"primary_email"`
	// SecondaryEmail   string  `json:"secondary_email"`
	PublicPhone string `json:"public_phone"`
	// SecondaryPhone   string  `json:"secondary_phone"`
	// OwnerName        string  `json:"owner_name"`
	// OwnerPhone       string  `json:"owner_phone"`
	FeaturedPriority int    `json:"featured_priority"`
	TrendingPriority int    `json:"trending_priority"`
	PetFriendly      bool   `json:"pet_friendly"`
	IsNew            bool   `json:"is_new"`
	ComingSoon       bool   `json:"coming_soon"`
	WorkFromBrewery  bool   `json:"work_from_brewery"`
	IsPublished      bool   `json:"is_published"`
	PublishedAt      string `json:"published_at"`
	SundayBrunch     bool   `json:"sunday_brunch"`
}

type BreweryLandingDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	LogoURL      string `json:"logo_url"`
	ThumbnailURL string `json:"thumbnail_url"`
	Description  string `json:"description"`
	PetFriendly  bool   `json:"pet_friendly"`
	IsNew        bool   `json:"is_new"`
	ComingSoon   bool   `json:"coming_soon"`
}
