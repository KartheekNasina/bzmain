package models

type BrewClassSchedule struct {
	ClassID        string `db:"class_id"`
	StartDate      string `db:"start_date"`
	EndDate        string `db:"end_date"`
	StartTime      string `db:"start_time"`
	EndTime        string `db:"end_time"`
	Repeats        string `db:"repeats"`
	MaxPeopleCount int    `db:"max_people_count"`
}
