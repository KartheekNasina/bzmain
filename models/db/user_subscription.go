// userSubscription.go
package models

import "time"

type UserSubscription struct {
	UserID             string    `db:"user_id"`
	SubscriptionTierID string    `db:"subscription_tier_id"`
	StartDate          time.Time `db:"start_date"`
	EndDate            time.Time `db:"end_date"`
	IsActive           bool      `db:"is_active"`
}
