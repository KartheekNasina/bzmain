// user_referrals.go
package models

import (
	"time"
)

type UserReferral struct {
	ID            string         `db:"id"`
	ReferrerID    string         `db:"referrer_id"`
	RefereeID     string         `db:"referee_id"`
	ReferralCode  string         `db:"referral_code"`
	Status        ReferralStatus `db:"status"`
	ReferredAt    time.Time      `db:"referred_at"`
	CompletedAt   time.Time      `db:"completed_at"`
	RewardClaimed bool           `db:"reward_claimed"`
}
