package models

import (
	"time"
)

type UserBreweryVisitDTO struct {
	Id string `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	VisitDate string `json:"visit_date" db:"visit_date"`
	VisitTime time.Time `json:"visit_time" db:"visit_time"`
}
