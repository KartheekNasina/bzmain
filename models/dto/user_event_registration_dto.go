package models

import (
	"time"
)

type UserEventRegistrationDTO struct {
	Id string `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	EventId string `json:"event_id" db:"event_id"`
	Status string `json:"status" db:"status"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	NumberOfPeople int `json:"number_of_people" db:"number_of_people"`
	Notes string `json:"notes" db:"notes"`
}
