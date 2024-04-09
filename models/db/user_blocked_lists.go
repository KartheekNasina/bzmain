// userBlockedList.go
package models

import "time"

type UserBlockedList struct {
	UserId        string    `db:"user_id"`
	BlockedUserId string    `db:"blocked_user_id"`
	BlockedAt     time.Time `db:"blocked_at"`
}
