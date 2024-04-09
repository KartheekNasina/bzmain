// userNotification.go
package models

import "time"

type UserNotification struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	IsRead    bool      `db:"is_read"`
	ReadAt    time.Time `db:"read_at"`
}
