// meetups.go
package models

import (
	"time"
)

type Meetup struct {
	ID                 string       `db:"id"`
	BreweryID          string       `db:"brewery_id"`
	OrganizerID        string       `db:"organizer_id"`
	Title              string       `db:"title"`
	Description        string       `db:"description"`
	MeetupDate         time.Time    `db:"meetup_date"`
	StartTime          time.Time    `db:"start_time"`
	EndTime            time.Time    `db:"end_time"`
	MaxAttendees       *int         `db:"max_attendees"`
	CreatedAt          time.Time    `db:"created_at"`
	UpdatedAt          time.Time    `db:"updated_at"`
	Status             MeetupStatus `db:"status"`
	ApprovedAt         time.Time    `db:"approved_at"`
	RejectedAt         time.Time    `db:"rejected_at"`
	CancellationReason string       `db:"cancellation_reason"`
}

type MeetupStatus string

const (
	MeetupStatusPendingApproval MeetupStatus = "pending_approval"
	MeetupStatusApproved        MeetupStatus = "approved"
	MeetupStatusRejected        MeetupStatus = "rejected"
	MeetupStatusCancelled       MeetupStatus = "cancelled"
)
