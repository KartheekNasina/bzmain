package models

import (
	"time"
)

type UserDTO struct {
	Id                          string     `json:"id" db:"id"`
	IsOnline                    bool       `json:"is_online" db:"is_online"`
	AllowNotifications          bool       `json:"allow_notifications" db:"allow_notifications"`
	AllowLocation               bool       `json:"allow_location" db:"allow_location"`
	IsLegalAge                  bool       `json:"is_legal_age" db:"is_legal_age"`
	PhoneNumberVerified         *bool      `json:"phone_number_verified" db:"phone_number_verified"`
	PhoneNumberVerificationDate time.Time  `json:"phone_number_verification_date" db:"phone_number_verification_date"`
	OnboardingStatus            *string    `json:"onboarding_status" db:"onboarding_status"`
	OtpVerifiedAt               *time.Time `json:"otp_verified_at" db:"otp_verified_at"`
	ProfileCompletedAt          *time.Time `json:"profile_completed_at" db:"profile_completed_at"`
	CreatedAt                   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt                   time.Time  `json:"updated_at" db:"updated_at"`
	Dob                         string     `json:"dob" db:"dob"`
	UserId                      string     `json:"user_id" db:"user_id"`
	Name                        string     `json:"name" db:"name"`
	AboutMe                     string     `json:"about_me" db:"about_me"`
	Gender                      *string    `json:"gender" db:"gender"`
	Email                       string     `json:"email" db:"email"`
	ProfileUrl                  string     `json:"profile_url" db:"profile_url"`
	Images                      []string   `json:"images" db:"images"`
	PhoneNumber                 *string    `json:"phone_number" db:"phone_number"`
	PersonalInterests           []string   `json:"personal_interests" db:"personal_interests"`
	Provider                    string     `json:"provider" db:"provider"`
	BrewInterests               []string   `json:"brew_interests" db:"brew_interests"`
}
