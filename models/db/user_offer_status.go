// userOfferStatus.go
package models

type UserOfferStatus string

const (
	Active   UserOfferStatus = "active"
	Redeemed UserOfferStatus = "redeemed"
	Expired  UserOfferStatus = "expired"
)
