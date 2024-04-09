package models

import (
	"time"
)

type PaymentLogDTO struct {
	Id string `json:"id" db:"id"`
	PaymentTransactionId string `json:"payment_transaction_id" db:"payment_transaction_id"`
	Interaction string `json:"interaction" db:"interaction"`
	Payload string `json:"payload" db:"payload"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Endpoint string `json:"endpoint" db:"endpoint"`
}
