package models

import (
	"time"
)

type BreweryFollowDTO struct {
	UserId string `json:"user_id" db:"user_id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	FollowedAt time.Time `json:"followed_at" db:"followed_at"`
}
