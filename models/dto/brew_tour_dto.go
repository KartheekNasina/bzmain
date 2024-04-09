package models

import "time"

type BrewTourDTO struct {
	ID               string    `json:"tour_id"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"short_description"`
	BreweryID        string    `json:"brewery_id"`
	Thumbnail        string    `json:"thumbnail"`
	IsPublished      bool      `json:"is_published"`
	PublishedAt      time.Time `json:"published_at"`
	Images           []string  `json:"images"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedByID      string    `json:"created_by_id"`
	UpdatedByID      string    `json:"updated_by_id"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	Repeats          string    `json:"repeats"`
	MaxPeopleCount   int       `json:"max_people_count"`
}
