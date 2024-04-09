package models

import (
	"time"
)

type RatingReactionDTO struct {
	UserId string `json:"user_id" db:"user_id"`
	RatingId string `json:"rating_id" db:"rating_id"`
	Reaction string `json:"reaction" db:"reaction"`
	ReactedAt time.Time `json:"reacted_at" db:"reacted_at"`
}
