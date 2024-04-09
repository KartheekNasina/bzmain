package models

import (
	"time"
)

type UserReferralDTO struct {
	RewardClaimed bool `json:"reward_claimed" db:"reward_claimed"`
	ReferrerId string `json:"referrer_id" db:"referrer_id"`
	RefereeId string `json:"referee_id" db:"referee_id"`
	Id string `json:"id" db:"id"`
	Status string `json:"status" db:"status"`
	ReferredAt time.Time `json:"referred_at" db:"referred_at"`
	CompletedAt time.Time `json:"completed_at" db:"completed_at"`
	ReferralCode string `json:"referral_code" db:"referral_code"`
}
