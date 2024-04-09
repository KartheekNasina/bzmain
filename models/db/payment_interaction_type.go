// paymentInteractionType.go
package models

type PaymentInteractionType string

const (
	Request  PaymentInteractionType = "request"
	Response PaymentInteractionType = "response"
)
