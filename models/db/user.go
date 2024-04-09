package models

import (
	"time"
)

type User struct {
	ID                          string     `db:"id"`
	UserID                      string     `db:"user_id"`
	Name                        string     `db:"name"`
	AboutMe                     string     `db:"about_me"`
	Gender                      *string    `db:"gender"`
	Email                       string     `db:"email"`
	ProfileURL                  string     `db:"profile_url"`
	Images                      []string   `db:"images"`
	DOB                         time.Time  `db:"dob"`
	IsOnline                    bool       `db:"is_online"`
	AllowNotifications          *bool      `db:"allow_notifications"`
	AllowLocation               *bool      `db:"allow_location"`
	PhoneNumber                 string     `db:"phone_number"`
	IsLegalAge                  bool       `db:"is_legal_age"`
	Provider                    string     `db:"provider"`
	BrewInterests               []string   `db:"brew_interests"`
	PersonalInterests           []string   `db:"personal_interests"`
	PhoneNumberVerified         *bool      `db:"phone_number_verified"`
	PhoneNumberVerificationDate time.Time  `db:"phone_number_verification_date"`
	OnboardingStatus            *string    `db:"onboarding_status"`
	OTPVerifiedAt               *time.Time `db:"otp_verified_at"`
	ProfileCompletedAt          *time.Time `db:"profile_completed_at"`
	CreatedAt                   time.Time  `db:"created_at"`
	UpdatedAt                   time.Time  `db:"updated_at"`
}
