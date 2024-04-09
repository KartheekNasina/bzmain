package models

import (
	"time"
)

type ChatRoomDTO struct {
	Id string `json:"id" db:"id"`
	OneUserId string `json:"one_user_id" db:"one_user_id"`
	TwoUserId string `json:"two_user_id" db:"two_user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
