package models

type CommunityDrive struct {
	ID                 string   `db:"id"`
	Title              string   `db:"title"`
	CommunityType      string   `db:"community_type"`
	Description        string   `db:"description"`
	Address            string   `db:"address"`
	Lat                float64  `db:"lat"`
	Lng                float64  `db:"lng"`
	ContactName        string   `db:"contact_name"`
	ContactPhoneNumber int64    `db:"contact_phone_number"`
	StartTime          string   `db:"start_time"`
	EndTime            string   `db:"end_time"`
	IsPublished        bool     `db:"is_published"`
	CreatedAt          string   `db:"created_at"`
	UpdatedAt          string   `db:"updated_at"`
	PublishedAt        string   `db:"published_at"`
	Thumbnail          string   `db:"thumbnail"`
	Images             []string `db:"images"`
}
