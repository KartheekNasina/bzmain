package models

type BreweryContact struct {
	BreweryID      string `db:"brewery_id"`
	PrimaryEmail   string `db:"primary_email"`
	SecondaryEmail string `db:"secondary_email"`
	PublicPhone    string `db:"public_phone"`
	SecondaryPhone string `db:"secondary_phone"`
	OwnerName      string `db:"owner_name"`
	OwnerPhone     string `db:"owner_phone"`
}
