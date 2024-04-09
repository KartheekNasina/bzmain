// eventJoinRequest.go
package models

import "time"

type EventJoinRequest struct {
	ID          string        `db:"id"`
	EventID     string        `db:"event_id"`
	UserID      string        `db:"user_id"`
	Status      RequestStatus `db:"status"`
	RequestedAt time.Time     `db:"requested_at"`
	ApprovedAt  time.Time     `db:"approved_at"`
	DeclinedAt  time.Time     `db:"declined_at"`
	ConfirmedAt time.Time     `db:"confirmed_at"`
	CanceledAt  time.Time     `db:"canceled_at"`
	CreatedAt   time.Time     `db:"created_at"`
	UpdatedAt   time.Time     `db:"updated_at"`
	Notes       string        `db:"notes"`
}
