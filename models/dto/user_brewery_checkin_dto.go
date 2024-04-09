package models

import (
	"time"
)

type UserBreweryCheckinDTO struct {
	Id string `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	CheckinDate string `json:"checkin_date" db:"checkin_date"`
	CheckinTime time.Time `json:"checkin_time" db:"checkin_time"`
	Notes string `json:"notes" db:"notes"`
}
