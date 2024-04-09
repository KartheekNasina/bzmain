package models

import "time"

type Payment struct {
	ID            string        `db:"id"`
	EntityType    string        `db:"entity_type"`
	EntityID      string        `db:"entity_id"`
	UserID        string        `db:"user_id"`
	Amount        float64       `db:"amount"`
	CurrencyCode  string        `db:"currency_code"`
	PaymentStatus PaymentStatus `db:"payment_status"`
	PaymentMethod string        `db:"payment_method"`
	TransactionID string        `db:"transaction_id"`
	PaymentDate   time.Time     `db:"payment_date"`
	RefundDate    time.Time     `db:"refund_date"`
	FailureReason string        `db:"failure_reason"`
	Notes         string        `db:"notes"`
}
