package models

import (
	"time"
)

type CommunityCategoryDTO struct {
	UpdatedById string `json:"updated_by_id" db:"updated_by_id"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedById string `json:"created_by_id" db:"created_by_id"`
	Id string `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Type string `json:"type" db:"type"`
	Thumbnail string `json:"thumbnail" db:"thumbnail"`
	Title string `json:"title" db:"title"`
}
