package models

import (
	"time"
)

type PaymentTransactionDTO struct {
	Id string `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	Amount string `json:"amount" db:"amount"`
	Status string `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CurrencyCode string `json:"currency_code" db:"currency_code"`
	RazorpayPaymentId string `json:"razorpay_payment_id" db:"razorpay_payment_id"`
	OrderDescription string `json:"order_description" db:"order_description"`
}
