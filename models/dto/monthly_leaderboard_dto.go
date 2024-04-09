package models

type MonthlyLeaderboardDTO struct {
	MonthYear string `json:"month_year" db:"month_year"`
	UserId string `json:"user_id" db:"user_id"`
	VisitsCount int `json:"visits_count" db:"visits_count"`
	Rank int `json:"rank" db:"rank"`
	Title string `json:"title" db:"title"`
}
