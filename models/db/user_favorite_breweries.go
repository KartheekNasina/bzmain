// user_favorite_breweries.go
package models

import (
	"time"
)

type UserFavoriteBrewery struct {
	UserID    string    `db:"user_id"`
	BreweryID string    `db:"brewery_id"`
	AddedDate time.Time `db:"added_date"`
}
