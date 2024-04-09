// user_feedback.go
package models

import (
	"time"
)

type UserFeedback struct {
	ID              string       `db:"id"`
	UserID          string       `db:"user_id"`
	FeedbackDate    time.Time    `db:"feedback_date"`
	Type            FeedbackType `db:"type"`
	Subject         string       `db:"subject"`
	Description     string       `db:"description"`
	IsResolved      bool         `db:"is_resolved"`
	ResolvedAt      time.Time    `db:"resolved_at"`
	ResolutionNotes string       `db:"resolution_notes"`
}

type FeedbackType string

const (
	FeedbackTypeSuggestion FeedbackType = "suggestion"
	FeedbackTypeIssue      FeedbackType = "issue"
	FeedbackTypeCompliment FeedbackType = "compliment"
	FeedbackTypeComplaint  FeedbackType = "complaint"
	FeedbackTypeOther      FeedbackType = "other"
)
