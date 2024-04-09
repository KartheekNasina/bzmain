package models

import (
	"time"
)

type BrewClassesBasicInfoDTO struct {
	DiscountPerc int `json:"discount_perc" db:"discount_perc"`
	UpdatedById string `json:"updated_by_id" db:"updated_by_id"`
	Lat string `json:"lat" db:"lat"`
	Lng string `json:"lng" db:"lng"`
	Cost string `json:"cost" db:"cost"`
	Id string `json:"id" db:"id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	IsPublished bool `json:"is_published" db:"is_published"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedById string `json:"created_by_id" db:"created_by_id"`
	Title string `json:"title" db:"title"`
	ShortDescription string `json:"short_description" db:"short_description"`
	Address string `json:"address" db:"address"`
	Thumbnail string `json:"thumbnail" db:"thumbnail"`
	Images []string `json:"images" db:"images"`
}
