package models

import (
	"time"
)

type BrewToursBasicInfoDTO struct {
	UpdatedById string `json:"updated_by_id" db:"updated_by_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedById string `json:"created_by_id" db:"created_by_id"`
	Id string `json:"id" db:"id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	IsPublished bool `json:"is_published" db:"is_published"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	Title string `json:"title" db:"title"`
	ShortDescription string `json:"short_description" db:"short_description"`
	Images []string `json:"images" db:"images"`
	Thumbnail string `json:"thumbnail" db:"thumbnail"`
}
