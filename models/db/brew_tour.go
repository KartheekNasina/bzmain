package models

import (
	"time"
)

type BrewTour struct {
	TourID           string    `db:"tour_id"`
	Title            string    `db:"title"`
	ShortDescription string    `db:"short_description"`
	BreweryID        string    `db:"brewery_id"`
	Thumbnail        string    `db:"thumbnail"`
	IsPublished      bool      `db:"is_published"`
	PublishedAt      time.Time `db:"published_at"`
	Images           []string  `db:"images"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
	CreatedByID      string    `db:"created_by_id"`
	UpdatedByID      string    `db:"updated_by_id"`
	StartDate        time.Time `db:"start_date"`
	EndDate          time.Time `db:"end_date"`
	StartTime        time.Time `db:"start_time"`
	EndTime          time.Time `db:"end_time"`
	Repeats          string    `db:"repeats"`
	MaxPeopleCount   int       `db:"max_people_count"`
}
