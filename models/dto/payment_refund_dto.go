package models

import (
	"time"
)

type PaymentRefundDTO struct {
	Id string `json:"id" db:"id"`
	PaymentTransactionId string `json:"payment_transaction_id" db:"payment_transaction_id"`
	Amount string `json:"amount" db:"amount"`
	RefundDate time.Time `json:"refund_date" db:"refund_date"`
	RazorpayRefundId string `json:"razorpay_refund_id" db:"razorpay_refund_id"`
	Reason string `json:"reason" db:"reason"`
}
