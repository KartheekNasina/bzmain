// event_schedule.go
package models

import (
	"time"
)

type EventSchedule struct {
	EventID        string    `db:"event_id"`
	StartDate      time.Time `db:"start_date"`
	EndDate        time.Time `db:"end_date"`
	StartTime      time.Time `db:"start_time"`
	EndTime        time.Time `db:"end_time"`
	Repeats        string    `db:"repeats"`
	MaxPeopleCount int       `db:"max_people_count"`
}
