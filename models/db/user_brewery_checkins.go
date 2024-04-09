// user_brewery_checkins.go
package models

import (
	"time"
)

type UserBreweryCheckin struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	BreweryID   string    `db:"brewery_id"`
	CheckinDate time.Time `db:"checkin_date"`
	CheckinTime time.Time `db:"checkin_time"`
	Notes       string    `db:"notes"`
}
