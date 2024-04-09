// user_history.go
package models

import (
	"time"
)

type UserHistory struct {
	ID              string          `db:"id"`
	UserID          string          `db:"user_id"`
	InteractionID   string          `db:"interaction_id"`
	InteractionType InteractionType `db:"interaction_type"`
	InteractionDate time.Time       `db:"interaction_date"`
	Notes           string          `db:"notes"`
}

type InteractionType string

const (
	InteractionTypeBreweryVisit    InteractionType = "brewery"
	InteractionTypeEventAttendance InteractionType = "event"
	InteractionTypeTourAttendance  InteractionType = "tour"
	InteractionTypeClassAttendance InteractionType = "class"
)
