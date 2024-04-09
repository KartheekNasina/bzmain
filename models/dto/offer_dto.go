package models

import (
	"time"
)

type OfferDTO struct {
	UpdatedById string `json:"updated_by_id" db:"updated_by_id"`
	IsPublished bool `json:"is_published" db:"is_published"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedById string `json:"created_by_id" db:"created_by_id"`
	Id string `json:"id" db:"id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	StartDate string `json:"start_date" db:"start_date"`
	EndDate string `json:"end_date" db:"end_date"`
	Count int `json:"count" db:"count"`
	AvailedCount int `json:"availed_count" db:"availed_count"`
	StartTime string `json:"start_time" db:"start_time"`
	EndTime string `json:"end_time" db:"end_time"`
	DiscountPerc int `json:"discount_perc" db:"discount_perc"`
	Title string `json:"title" db:"title"`
	Type string `json:"type" db:"type"`
	Day string `json:"day" db:"day"`
	Description string `json:"description" db:"description"`
	Thumbnail string `json:"thumbnail" db:"thumbnail"`
	Images []string `json:"images" db:"images"`
}
