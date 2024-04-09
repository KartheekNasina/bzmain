// user_brew_class_registrations.go
package models

import (
	"time"
)

type UserBrewClassRegistration struct {
	ID                string                  `db:"id"`
	UserID            string                  `db:"user_id"`
	ClassID           string                  `db:"class_id"`
	Status            ClassRegistrationStatus `db:"status"`
	RegistrationDate  time.Time               `db:"registration_date"`
	CreatedAt         time.Time               `db:"created_at"`
	UpdatedAt         time.Time               `db:"updated_at"`
	NumberOfAttendees int                     `db:"number_of_attendees"`
	Notes             string                  `db:"notes"`
}

type ClassRegistrationStatus string

const (
	ClassRegistrationPending    ClassRegistrationStatus = "pending"
	ClassRegistrationConfirmed  ClassRegistrationStatus = "confirmed"
	ClassRegistrationWaitlisted ClassRegistrationStatus = "waitlisted"
	ClassRegistrationCanceled   ClassRegistrationStatus = "canceled"
)
