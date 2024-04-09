// beer.go
package models

import "time"

type Beer struct {
	ID          string     `db:"id"`
	Name        string     `db:"name"`
	Type        string     `db:"type"`
	ImageURL    string     `db:"image_url"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	IsPublished bool       `db:"is_published"`
	PublishedAt *time.Time `db:"published_at"`
	BreweryID   string     `db:"brewery_id"`
	ABV         float64    `db:"abv"`
	IBU         float64    `db:"ibu"`
	Description string     `db:"description"`
}
