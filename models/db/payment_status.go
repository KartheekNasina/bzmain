// paymentStatus.go
package models

type PaymentStatus string

const (
	Pending   PaymentStatus = "pending"
	Completed PaymentStatus = "completed"
	Failed    PaymentStatus = "failed"
	Refunded  PaymentStatus = "refunded"
)
