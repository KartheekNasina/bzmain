package models

import (
	"time"
)

type ChatMessageDTO struct {
	Id string `json:"id" db:"id"`
	RoomId string `json:"room_id" db:"room_id"`
	SenderId string `json:"sender_id" db:"sender_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedById string `json:"created_by_id" db:"created_by_id"`
	UpdatedById string `json:"updated_by_id" db:"updated_by_id"`
	Status string `json:"status" db:"status"`
	Type string `json:"type" db:"type"`
	Content string `json:"content" db:"content"`
	MediaUrl string `json:"media_url" db:"media_url"`
	MediaMimeType string `json:"media_mime_type" db:"media_mime_type"`
}
