package models

import (
	"time"
)

type DrinkPurchaseRequestDTO struct {
	Id string `json:"id" db:"id"`
	BuyerId string `json:"buyer_id" db:"buyer_id"`
	RecipientId string `json:"recipient_id" db:"recipient_id"`
	MessageId string `json:"message_id" db:"message_id"`
	RequestDate time.Time `json:"request_date" db:"request_date"`
	Status string `json:"status" db:"status"`
}
