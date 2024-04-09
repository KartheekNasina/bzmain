// user_connection_requests.go
package models

import (
	"time"
)

type UserConnectionRequest struct {
	ID          string                  `db:"id"`
	SenderID    string                  `db:"sender_id"`
	ReceiverID  string                  `db:"receiver_id"`
	RequestDate time.Time               `db:"request_date"`
	Status      ConnectionRequestStatus `db:"status"`
}

type ConnectionRequestStatus string

const (
	ConnectionRequestPending  ConnectionRequestStatus = "pending"
	ConnectionRequestAccepted ConnectionRequestStatus = "accepted"
	ConnectionRequestRejected ConnectionRequestStatus = "rejected"
)
