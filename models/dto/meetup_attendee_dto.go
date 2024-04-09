package models

import (
	"time"
)

type MeetupAttendeeDTO struct {
	MeetupId string `json:"meetup_id" db:"meetup_id"`
	UserId string `json:"user_id" db:"user_id"`
	RsvpDate time.Time `json:"rsvp_date" db:"rsvp_date"`
	Attended bool `json:"attended" db:"attended"`
}
