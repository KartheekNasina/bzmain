package models

import "time"

type PaymentLog struct {
	ID                   string                 `db:"id"`
	PaymentTransactionID string                 `db:"payment_transaction_id"`
	Interaction          PaymentInteractionType `db:"interaction"`
	Endpoint             string                 `db:"endpoint"`
	Payload              map[string]interface{} `db:"payload"`
	Timestamp            time.Time              `db:"timestamp"`
}
