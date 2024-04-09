package models

import "time"

type FoodDrinkRating struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	ItemID    string    `db:"item_id"`
	Rating    int       `db:"rating"`
	CommentID string    `db:"comment_id"`
	RatedAt   time.Time `db:"rated_at"`
}
