package models

import (
	"time"
)

type UserHistoryDTO struct {
	Id string `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	InteractionId string `json:"interaction_id" db:"interaction_id"`
	InteractionType string `json:"interaction_type" db:"interaction_type"`
	InteractionDate time.Time `json:"interaction_date" db:"interaction_date"`
	Notes string `json:"notes" db:"notes"`
}
