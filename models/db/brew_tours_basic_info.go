// brewTourBasicInfo.go
package models

import "time"

type BrewTourBasicInfo struct {
	ID               string    `db:"id"`
	Title            string    `db:"title"`
	ShortDescription string    `db:"short_description"`
	BreweryID        string    `db:"brewery_id"`
	Thumbnail        string    `db:"thumbnail"`
	IsPublished      bool      `db:"is_published"`
	PublishedAt      time.Time `db:"published_at"`
	Images           []string  `db:"images"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
	CreatedByID      string    `db:"created_by_id"`
	UpdatedByID      string    `db:"updated_by_id"`
}
