package models

import (
	"time"
)

type UserSubscriptionDTO struct {
	UserId string `json:"user_id" db:"user_id"`
	SubscriptionTierId string `json:"subscription_tier_id" db:"subscription_tier_id"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate time.Time `json:"end_date" db:"end_date"`
	IsActive bool `json:"is_active" db:"is_active"`
}
