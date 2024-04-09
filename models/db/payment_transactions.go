package models

import "time"

type PaymentTransaction struct {
	ID                string        `db:"id"`
	UserID            string        `db:"user_id"`
	Amount            float64       `db:"amount"`
	CurrencyCode      string        `db:"currency_code"`
	RazorpayPaymentID string        `db:"razorpay_payment_id"`
	Status            PaymentStatus `db:"status"`
	OrderDescription  string        `db:"order_description"`
	CreatedAt         time.Time     `db:"created_at"`
	UpdatedAt         time.Time     `db:"updated_at"`
}
