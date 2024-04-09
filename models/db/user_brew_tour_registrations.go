package models

import "time"

type UserBrewTourRegistration struct {
	ID               string    `db:"id"`
	UserID           string    `db:"user_id"`
	TourID           string    `db:"tour_id"`
	Status           string    `db:"status"`
	RegistrationDate string    `db:"registration_date"`
	NumberOfPeople   int       `db:"number_of_people"`
	Notes            string    `db:"notes"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}
