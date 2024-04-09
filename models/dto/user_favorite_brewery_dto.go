package models

import (
	"time"
)

type UserFavoriteBreweryDTO struct {
	UserId string `json:"user_id" db:"user_id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	AddedDate time.Time `json:"added_date" db:"added_date"`
}
