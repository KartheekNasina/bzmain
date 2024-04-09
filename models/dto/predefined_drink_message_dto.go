package models

type PredefinedDrinkMessageDTO struct {
	Id string `json:"id" db:"id"`
	MessageText string `json:"message_text" db:"message_text"`
}
