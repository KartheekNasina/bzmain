package models

import (
	"time"
)

type BeerDTO struct {
	Id string `json:"id" db:"id"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	Abv string `json:"abv" db:"abv"`
	Ibu string `json:"ibu" db:"ibu"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	IsPublished bool `json:"is_published" db:"is_published"`
	Name string `json:"name" db:"name"`
	Type string `json:"type" db:"type"`
	ImageUrl string `json:"image_url" db:"image_url"`
	Description string `json:"description" db:"description"`
}
