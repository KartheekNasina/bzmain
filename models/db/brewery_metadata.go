package models

type BreweryMetadata struct {
	BreweryID        string `db:"brewery_id"`
	FeaturedPriority int    `db:"featured_priority"`
	TrendingPriority int    `db:"trending_priority"`
	PetFriendly      bool   `db:"pet_friendly"`
	IsNew            bool   `db:"is_new"`
	ComingSoon       bool   `db:"coming_soon"`
	WorkFromBrewery  bool   `db:"work_from_brewery"`
	IsPublished      bool   `db:"is_published"`
	PublishedAt      string `db:"published_at"`
	SundayBrunch     bool   `db:"sunday_brunch"`
}
