package models

import (
	"time"
)

type UserNotificationDTO struct {
	Id string `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	IsRead bool `json:"is_read" db:"is_read"`
	ReadAt time.Time `json:"read_at" db:"read_at"`
	Content string `json:"content" db:"content"`
}
