// brewTourSchedule.go
package models

import "time"

type BrewTourSchedule struct {
	TourID         string        `db:"tour_id"`
	StartDate      time.Time     `db:"start_date"`
	EndDate        *time.Time    `db:"end_date"`
	StartTime      time.Time     `db:"start_time"`
	EndTime        time.Time     `db:"end_time"`
	Repeats        RepeatPattern `db:"repeats"`
	MaxPeopleCount int           `db:"max_people_count"`
}
