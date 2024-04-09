// userOfferRedemption.go
package models

import "time"

type UserOfferRedemption struct {
	ID             string    `db:"id"`
	OfferID        string    `db:"offer_id"`
	UserID         string    `db:"user_id"`
	RedemptionDate time.Time `db:"redemption_date"`
	Notes          string    `db:"notes"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
