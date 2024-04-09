package models

import (
	"time"
)

type UserCommunityDriveRegistrationDTO struct {
	Id string `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	DriveId string `json:"drive_id" db:"drive_id"`
	Status string `json:"status" db:"status"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	Notes string `json:"notes" db:"notes"`
}
