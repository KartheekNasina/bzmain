package models

import (
	"time"
)

type CommunityCenterDTO struct {
	ContactPhoneNumber string `json:"contact_phone_number" db:"contact_phone_number"`
	IsPublished bool `json:"is_published" db:"is_published"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Lng string `json:"lng" db:"lng"`
	Id string `json:"id" db:"id"`
	Lat string `json:"lat" db:"lat"`
	Images []string `json:"images" db:"images"`
	Title string `json:"title" db:"title"`
	CommunityType string `json:"community_type" db:"community_type"`
	Description string `json:"description" db:"description"`
	Address string `json:"address" db:"address"`
	ContactName string `json:"contact_name" db:"contact_name"`
	Thumbnail string `json:"thumbnail" db:"thumbnail"`
}
