// brewery_basic_info.go
package models

import (
	"time"
)

type BreweryBasicInfo struct {
	ID               string    `db:"id"`
	Name             string    `db:"name"`
	LogoURL          string    `db:"logo_url"`
	ThumbnailURL     string    `db:"thumbnail_url"`
	ShortDescription string    `db:"short_description"`
	Description      string    `db:"description"`
	Images           []string  `db:"images"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}
