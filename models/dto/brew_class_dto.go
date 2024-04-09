package models

import (
	"time"
)

type BrewClassDTO struct {
	DiscountPerc     int       `json:"discount_perc" db:"discount_perc"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	Lat              string    `json:"lat" db:"lat"`
	Lng              string    `json:"lng" db:"lng"`
	Cost             string    `json:"cost" db:"cost"`
	Id               string    `json:"id" db:"id"`
	BreweryId        string    `json:"brewery_id" db:"brewery_id"`
	StartDate        string    `json:"start_date" db:"start_date"`
	StartTime        string    `json:"start_time" db:"start_time"`
	EndTime          string    `json:"end_time" db:"end_time"`
	EndDate          string    `json:"end_date" db:"end_date"`
	MaxPeopleCount   int       `json:"max_people_count" db:"max_people_count"`
	IsPublished      bool      `json:"is_published" db:"is_published"`
	PublishedAt      time.Time `json:"published_at" db:"published_at"`
	CreatedById      string    `json:"created_by_id" db:"created_by_id"`
	UpdatedById      string    `json:"updated_by_id" db:"updated_by_id"`
	Title            string    `json:"title" db:"title"`
	ShortDescription string    `json:"short_description" db:"short_description"`
	Repeats          string    `json:"repeats" db:"repeats"`
	Images           []string  `json:"images" db:"images"`
	Address          string    `json:"address" db:"address"`
	Thumbnail        string    `json:"thumbnail" db:"thumbnail"`
	Schedules        BrewClassesScheduleDTO
}
