package models

type ReferralStatus string

const (
	ReferralStatusPending    ReferralStatus = "pending"
	ReferralStatusSuccessful ReferralStatus = "successful"
	ReferralStatusExpired    ReferralStatus = "expired"
)
