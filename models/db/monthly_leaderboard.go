// monthlyLeaderboard.go
package models

import "time"

type MonthlyLeaderboard struct {
	MonthYear   time.Time `db:"month_year"`
	UserID      string    `db:"user_id"`
	VisitsCount int       `db:"visits_count"`
	Rank        int       `db:"rank"`
	Title       string    `db:"title"`
}
