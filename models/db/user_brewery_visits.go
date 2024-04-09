// userBreweryVisit.go
package models

import "time"

type UserBreweryVisit struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	BreweryID string    `db:"brewery_id"`
	VisitDate time.Time `db:"visit_date"`
	VisitTime time.Time `db:"visit_time"`
}
