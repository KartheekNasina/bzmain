package models

import (
	"time"
)

type UserExclusiveOfferDTO struct {
	UpdatedById string `json:"updated_by_id" db:"updated_by_id"`
	UserId string `json:"user_id" db:"user_id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedById string `json:"created_by_id" db:"created_by_id"`
	Id string `json:"id" db:"id"`
	StartDate string `json:"start_date" db:"start_date"`
	EndDate string `json:"end_date" db:"end_date"`
	StartTime string `json:"start_time" db:"start_time"`
	EndTime string `json:"end_time" db:"end_time"`
	Repeats string `json:"repeats" db:"repeats"`
	Status string `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Title string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Thumbnail string `json:"thumbnail" db:"thumbnail"`
}
