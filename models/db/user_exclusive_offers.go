// userExclusiveOffer.go
package models

import "time"

type UserExclusiveOffer struct {
	ID          string          `db:"id"`
	UserID      string          `db:"user_id"`
	BreweryID   string          `db:"brewery_id"`
	Title       string          `db:"title"`
	Description string          `db:"description"`
	Thumbnail   string          `db:"thumbnail"`
	StartDate   time.Time       `db:"start_date"`
	EndDate     *time.Time      `db:"end_date"`
	StartTime   time.Time       `db:"start_time"`
	EndTime     time.Time       `db:"end_time"`
	Repeats     RepeatPattern   `db:"repeats"`
	Status      UserOfferStatus `db:"status"`
	CreatedAt   time.Time       `db:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at"`
	CreatedByID string          `db:"created_by_id"`
	UpdatedByID string          `db:"updated_by_id"`
}
