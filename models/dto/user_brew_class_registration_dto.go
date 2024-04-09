package models

import (
	"time"
)

type UserBrewClassRegistrationDTO struct {
	Id string `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	ClassId string `json:"class_id" db:"class_id"`
	Status string `json:"status" db:"status"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	NumberOfAttendees int `json:"number_of_attendees" db:"number_of_attendees"`
	Notes string `json:"notes" db:"notes"`
}
