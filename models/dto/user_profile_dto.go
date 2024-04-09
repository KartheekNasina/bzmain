package models

import (
	"time"
)

type UserProfileDTO struct {
	Id string `json:"id" db:"id"`
	IsOnline bool `json:"is_online" db:"is_online"`
	BlockUser bool `json:"block_user" db:"block_user"`
	IsLegalAge bool `json:"is_legal_age" db:"is_legal_age"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	ProfileUrl string `json:"profile_url" db:"profile_url"`
	Interests []string `json:"interests" db:"interests"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Provider string `json:"provider" db:"provider"`
	UserId string `json:"user_id" db:"user_id"`
	Name string `json:"name" db:"name"`
	AboutMe string `json:"about_me" db:"about_me"`
	Gender string `json:"gender" db:"gender"`
	Email string `json:"email" db:"email"`
}
