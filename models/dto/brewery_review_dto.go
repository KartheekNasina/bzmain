package models

import (
	"time"
)

type BreweryReviewDTO struct {
	Id string `json:"id" db:"id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	Rating int `json:"rating" db:"rating"`
	UserId string `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Title string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}
