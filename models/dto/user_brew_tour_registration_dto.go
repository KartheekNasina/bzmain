package models

import (
	"time"
)

type UserBrewTourRegistrationDTO struct {
	Id               string    `json:"id" db:"id"`
	UserId           string    `json:"user_id" db:"user_id"`
	TourId           string    `json:"tour_id" db:"tour_id"`
	Status           string    `json:"status" db:"status"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	NumberOfPeople   int       `json:"number_of_people" db:"number_of_people"`
	Notes            string    `json:"notes" db:"notes"`
}
