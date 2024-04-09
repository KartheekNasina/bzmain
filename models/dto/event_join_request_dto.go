package models

import (
	"time"
)

type EventJoinRequestDTO struct {
	Id string `json:"id" db:"id"`
	EventId string `json:"event_id" db:"event_id"`
	UserId string `json:"user_id" db:"user_id"`
	Status string `json:"status" db:"status"`
	RequestedAt time.Time `json:"requested_at" db:"requested_at"`
	ApprovedAt time.Time `json:"approved_at" db:"approved_at"`
	DeclinedAt time.Time `json:"declined_at" db:"declined_at"`
	ConfirmedAt time.Time `json:"confirmed_at" db:"confirmed_at"`
	CanceledAt time.Time `json:"canceled_at" db:"canceled_at"`
	Notes string `json:"notes" db:"notes"`
}
