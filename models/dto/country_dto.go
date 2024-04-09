package models

import (
	"time"
)

type CountryDTO struct {
	Id string `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name string `json:"name" db:"name"`
	Code string `json:"code" db:"code"`
}
