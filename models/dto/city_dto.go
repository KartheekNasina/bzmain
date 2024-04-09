package models

import (
	"time"
)

type CityDTO struct {
	Id string `json:"id" db:"id"`
	CountryId string `json:"country_id" db:"country_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name string `json:"name" db:"name"`
}
