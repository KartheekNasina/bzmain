// user_community_drive_registrations.go
package models

import (
	"time"
)

type UserCommunityDriveRegistration struct {
	ID               string             `db:"id"`
	UserID           string             `db:"user_id"`
	DriveID          string             `db:"drive_id"`
	Status           RegistrationStatus `db:"status"`
	RegistrationDate time.Time          `db:"registration_date"`
	CreatedAt        time.Time          `db:"created_at"`
	UpdatedAt        time.Time          `db:"updated_at"`
	Notes            string             `db:"notes"`
}

type RegistrationStatus string

const (
	RegistrationPending    RegistrationStatus = "pending"
	RegistrationConfirmed  RegistrationStatus = "confirmed"
	RegistrationWaitlisted RegistrationStatus = "waitlisted"
	RegistrationCanceled   RegistrationStatus = "canceled"
)
