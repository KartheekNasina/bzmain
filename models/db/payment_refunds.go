package models

import "time"

type PaymentRefund struct {
	ID               string    `db:"id"`
	TransactionID    string    `db:"payment_transaction_id"`
	RazorpayRefundID string    `db:"razorpay_refund_id"`
	Amount           float64   `db:"amount"`
	Reason           string    `db:"reason"`
	RefundDate       time.Time `db:"refund_date"`
}
