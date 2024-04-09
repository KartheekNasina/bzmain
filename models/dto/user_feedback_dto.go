package models

import (
	"time"
)

type UserFeedbackDTO struct {
	IsResolved bool `json:"is_resolved" db:"is_resolved"`
	ResolvedAt time.Time `json:"resolved_at" db:"resolved_at"`
	Id string `json:"id" db:"id"`
	Type string `json:"type" db:"type"`
	UserId string `json:"user_id" db:"user_id"`
	FeedbackDate time.Time `json:"feedback_date" db:"feedback_date"`
	ResolutionNotes string `json:"resolution_notes" db:"resolution_notes"`
	Subject string `json:"subject" db:"subject"`
	Description string `json:"description" db:"description"`
}
