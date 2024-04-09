package models

import (
	"time"
)

type EventDTO struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"short_description"`
	BreweryID        string    `json:"brewery_id"`
	Thumbnail        string    `json:"thumbnail"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	IsPublished      bool      `json:"is_published"`
	PublishedAt      time.Time `json:"published_at"`
	CreatedByID      string    `json:"created_by_id"`
	UpdatedByID      string    `json:"updated_by_id"`
	Images           []string  `json:"images"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	StartTime        string    `json:"start_time"`
	EndTime          string    `json:"end_time"`
	Repeats          string    `json:"repeats"`
	MaxPeopleCount   int       `json:"max_people_count"`
}
