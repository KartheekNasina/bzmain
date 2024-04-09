package models

import (
	"time"
)

type MeetupAttendee struct {
	MeetupID string    `db:"meetup_id"`
	UserID   string    `db:"user_id"`
	RSVPDate time.Time `db:"rsvp_date"`
	Attended bool      `db:"attended"`
}
