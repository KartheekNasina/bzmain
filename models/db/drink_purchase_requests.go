// drinkRequestStatus.go
package models

import "time"

type DrinkRequestStatus string

const (
	DrinkRequestPending  DrinkRequestStatus = "pending"
	DrinkRequestAccepted DrinkRequestStatus = "accepted"
	DrinkRequestDeclined DrinkRequestStatus = "declined"
)

type DrinkPurchaseRequest struct {
	ID          string             `db:"id"`
	BuyerID     string             `db:"buyer_id"`
	RecipientID string             `db:"recipient_id"`
	MessageID   string             `db:"message_id"`
	RequestDate time.Time          `db:"request_date"`
	Status      DrinkRequestStatus `db:"status"`
}
