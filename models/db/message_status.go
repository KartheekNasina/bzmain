// message_status.go
package models

import (
	"time"
)

type MessageStatus struct {
	MessageID   string    `db:"message_id"`
	UserID      string    `db:"user_id"`
	DeliveredAt time.Time `db:"delivered_at"`
	ReadAt      time.Time `db:"read_at"`
}
