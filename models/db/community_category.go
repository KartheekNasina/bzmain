package models

type CommunityCategory struct {
	ID          string `db:"id"`
	Type        string `db:"type"`
	Thumbnail   string `db:"thumbnail"`
	Title       string `db:"title"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
	CreatedByID string `db:"created_by_id"`
	UpdatedByID string `db:"updated_by_id"`
}
