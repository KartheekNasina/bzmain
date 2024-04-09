package models

type SubscriptionTierDTO struct {
	Id string `json:"id" db:"id"`
	Type string `json:"type" db:"type"`
	DailyMessageLimit int `json:"daily_message_limit" db:"daily_message_limit"`
	UnlimitedSwipes bool `json:"unlimited_swipes" db:"unlimited_swipes"`
	SeeReadReceipts bool `json:"see_read_receipts" db:"see_read_receipts"`
	Price string `json:"price" db:"price"`
	DurationDays int `json:"duration_days" db:"duration_days"`
}
