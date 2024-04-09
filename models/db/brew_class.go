package models

import "time"

// BrewClass represents the structure for the brew_classes table
type BrewClass struct {
	Id               string    `json:"id" db:"id"`
	Title            string    `json:"title" db:"title"`
	Thumbnail        string    `json:"thumbnail" db:"thumbnail"`
	ShortDescription string    `json:"short_description" db:"short_description"`
	BreweryId        *string   `json:"brewery_id" db:"brewery_id"`
	StartDate        time.Time `json:"start_date" db:"start_date"`
	EndDate          time.Time `json:"end_date" db:"end_date"`
	// StartTime        time.Time  `json:"start_time" db:"start_time"`
	// EndTime          time.Time  `json:"end_time" db:"end_time"`
	// Repeats          *string    `json:"repeats" db:"repeats"`
	// MaxPeopleCount   int        `json:"max_people_count" db:"max_people_count"`
	// IsPublished      bool       `json:"is_published" db:"is_published"`
	// PublishedAt      *time.Time `json:"published_at" db:"published_at"`
	// CreatedById      *string    `json:"created_by_id" db:"created_by_id"`
	// UpdatedById      *string    `json:"updated_by_id" db:"updated_by_id"`
	// Images           []string   `json:"images" db:"images"`
	// CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	// UpdatedAt        time.Time  `json:"updated_at" db:"updated_at"`
	// Address          string     `json:"address" db:"address"`
	// Lat              *float64   `json:"lat" db:"lat"`
	// Lng              *float64   `json:"lng" db:"lng"`
	// Cost             float64    `json:"cost" db:"cost"`
	// DiscountPerc     int        `json:"discount_perc" db:"discount_perc"`
}
