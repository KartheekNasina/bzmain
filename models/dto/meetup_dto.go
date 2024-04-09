package models

import (
	"time"
)

type MeetupDTO struct {
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	CreatedById string `json:"created_by_id" db:"created_by_id"`
	UpdatedById string `json:"updated_by_id" db:"updated_by_id"`
	IsPublished bool `json:"is_published" db:"is_published"`
	Id string `json:"id" db:"id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	StartTime time.Time `json:"start_time" db:"start_time"`
	MaxPeopleCount int `json:"max_people_count" db:"max_people_count"`
	JoinedPeopleCount int `json:"joined_people_count" db:"joined_people_count"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name string `json:"name" db:"name"`
	OrganizerName string `json:"organizer_name" db:"organizer_name"`
	Thumbnail string `json:"thumbnail" db:"thumbnail"`
	Description string `json:"description" db:"description"`
}
