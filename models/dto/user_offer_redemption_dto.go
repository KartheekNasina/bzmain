package models

import (
	"time"
)

type UserOfferRedemptionDTO struct {
	RedemptionDate time.Time `json:"redemption_date" db:"redemption_date"`
	OfferId string `json:"offer_id" db:"offer_id"`
	UserId string `json:"user_id" db:"user_id"`
	Id string `json:"id" db:"id"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Notes string `json:"notes" db:"notes"`
}
