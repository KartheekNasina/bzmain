package models

import (
	"time"
)

type BreweryMetadataDTO struct {
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	FeaturedPriority int `json:"featured_priority" db:"featured_priority"`
	TrendingPriority int `json:"trending_priority" db:"trending_priority"`
	PetFriendly bool `json:"pet_friendly" db:"pet_friendly"`
	IsNew bool `json:"is_new" db:"is_new"`
	ComingSoon bool `json:"coming_soon" db:"coming_soon"`
	WorkFromBrewery bool `json:"work_from_brewery" db:"work_from_brewery"`
	IsPublished bool `json:"is_published" db:"is_published"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	SundayBrunch bool `json:"sunday_brunch" db:"sunday_brunch"`
}
