package models

import (
	"time"
)

type BreweryBasicInfoDTO struct {
	Id string `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	ThumbnailUrl string `json:"thumbnail_url" db:"thumbnail_url"`
	ShortDescription string `json:"short_description" db:"short_description"`
	Description string `json:"description" db:"description"`
	Images []string `json:"images" db:"images"`
	Name string `json:"name" db:"name"`
	LogoUrl string `json:"logo_url" db:"logo_url"`
}
