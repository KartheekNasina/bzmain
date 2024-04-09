package models

import (
	"time"
)

type TourDTO struct {
	UpdatedById string `json:"updated_by_id" db:"updated_by_id"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedById string `json:"created_by_id" db:"created_by_id"`
	Id string `json:"id" db:"id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	StartDate string `json:"start_date" db:"start_date"`
	StartTime string `json:"start_time" db:"start_time"`
	EndTime string `json:"end_time" db:"end_time"`
	EndDate string `json:"end_date" db:"end_date"`
	MaxPeopleCount int `json:"max_people_count" db:"max_people_count"`
	IsPublished bool `json:"is_published" db:"is_published"`
	Title string `json:"title" db:"title"`
	ShortDescription string `json:"short_description" db:"short_description"`
	Images []string `json:"images" db:"images"`
	Repeats string `json:"repeats" db:"repeats"`
	Thumbnail string `json:"thumbnail" db:"thumbnail"`
}
