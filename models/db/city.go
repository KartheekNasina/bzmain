// city.go
package models

import "time"

type City struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CountryID string    `db:"country_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
