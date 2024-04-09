package models

type BreweryOffer struct {
	ID          string   `db:"id"`
	BreweryID   string   `db:"brewery_id"`
	OfferType   string   `db:"offer_type"`
	Title       string   `db:"title"`
	Description string   `db:"description"`
	Thumbnail   string   `db:"thumbnail"`
	Images      []string `db:"images"`
	StartDate   string   `db:"start_date"`
	EndDate     string   `db:"end_date"`
	StartTime   string   `db:"start_time"`
	EndTime     string   `db:"end_time"`
	Repeats     string   `db:"repeats"`
	Status      string   `db:"status"`
	CreatedAt   string   `db:"created_at"`
	UpdatedAt   string   `db:"updated_at"`
	CreatedByID string   `db:"created_by_id"`
	UpdatedByID string   `db:"updated_by_id"`
}
