package models

type EventScheduleDTO struct {
	EventId string `json:"event_id" db:"event_id"`
	StartDate string `json:"start_date" db:"start_date"`
	EndDate string `json:"end_date" db:"end_date"`
	StartTime string `json:"start_time" db:"start_time"`
	EndTime string `json:"end_time" db:"end_time"`
	Repeats string `json:"repeats" db:"repeats"`
	MaxPeopleCount int `json:"max_people_count" db:"max_people_count"`
}