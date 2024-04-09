package models

import "time"

type UserConnectionDTO struct {
	ID             string    `json:"id"`
	User1ID        string    `json:"user1_id"`
	User2ID        string    `json:"user2_id"`
	ConnectionDate time.Time `json:"connection_date"`
}
