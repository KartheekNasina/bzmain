// user_connections.go
package models

import (
	"time"
)

type UserConnection struct {
	ID             string    `db:"id"`
	User1ID        string    `db:"user1_id"`
	User2ID        string    `db:"user2_id"`
	ConnectionDate time.Time `db:"connection_date"`
}
