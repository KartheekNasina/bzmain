package models

type DailyMessageCountDTO struct {
	UserId string `json:"user_id" db:"user_id"`
	Date string `json:"date" db:"date"`
	MessageCount int `json:"message_count" db:"message_count"`
}
