package models

import (
	"time"
)

type CommunityDriveDTO struct {
	Id string `json:"id" db:"id"`
	ContactPhoneNumber string `json:"contact_phone_number" db:"contact_phone_number"`
	StartTime time.Time `json:"start_time" db:"start_time"`
	EndTime time.Time `json:"end_time" db:"end_time"`
	IsPublished bool `json:"is_published" db:"is_published"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	Lat string `json:"lat" db:"lat"`
	Lng string `json:"lng" db:"lng"`
	Title string `json:"title" db:"title"`
	CommunityType string `json:"community_type" db:"community_type"`
	Description string `json:"description" db:"description"`
	Address string `json:"address" db:"address"`
	Images []string `json:"images" db:"images"`
	Thumbnail string `json:"thumbnail" db:"thumbnail"`
	ContactName string `json:"contact_name" db:"contact_name"`
}
