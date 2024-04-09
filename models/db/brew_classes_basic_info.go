package models

type BrewClassBasicInfo struct {
	ID               string   `db:"id"`
	Title            string   `db:"title"`
	ShortDescription string   `db:"short_description"`
	BreweryID        string   `db:"brewery_id"`
	Thumbnail        string   `db:"thumbnail"`
	IsPublished      bool     `db:"is_published"`
	PublishedAt      string   `db:"published_at"`
	Images           []string `db:"images"`
	CreatedAt        string   `db:"created_at"`
	UpdatedAt        string   `db:"updated_at"`
	CreatedByID      string   `db:"created_by_id"`
	UpdatedByID      string   `db:"updated_by_id"`
	Address          string   `db:"address"`
	Lat              float64  `db:"lat"`
	Lng              float64  `db:"lng"`
	Cost             float64  `db:"cost"`
	DiscountPerc     int      `db:"discount_perc"`
}
