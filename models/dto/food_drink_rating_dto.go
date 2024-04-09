package models

import (
	"time"
)

type FoodDrinkRatingDTO struct {
	Id string `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	ItemId string `json:"item_id" db:"item_id"`
	Rating int `json:"rating" db:"rating"`
	CommentId string `json:"comment_id" db:"comment_id"`
	RatedAt time.Time `json:"rated_at" db:"rated_at"`
}
