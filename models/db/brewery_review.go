package models

import "time"

type BreweryReview struct {
	ID          string    `db:"id"`
	BreweryID   string    `db:"brewery_id"`
	Rating      int       `db:"rating"`
	UserID      string    `db:"user_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
