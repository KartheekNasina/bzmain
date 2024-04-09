// community_centers.go
package models

import (
	"time"
)

type CommunityCenter struct {
	ID                 string    `db:"id"`
	Title              string    `db:"title"`
	CommunityType      string    `db:"community_type"`
	Description        string    `db:"description"`
	Address            string    `db:"address"`
	Lat                float64   `db:"lat"`
	Lng                float64   `db:"lng"`
	ContactName        string    `db:"contact_name"`
	ContactPhoneNumber int64     `db:"contact_phone_number"`
	IsPublished        bool      `db:"is_published"`
	PublishedAt        time.Time `db:"published_at"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
	Thumbnail          string    `db:"thumbnail"`
	Images             []string  `db:"images"`
}
