package models

import (
	"time"
)

type UserBlockedListDTO struct {
	UserId string `json:"user_id" db:"user_id"`
	BlockedUserId string `json:"blocked_user_id" db:"blocked_user_id"`
	BlockedAt time.Time `json:"blocked_at" db:"blocked_at"`
}
