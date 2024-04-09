package models

type BreweryContactDTO struct {
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	PrimaryEmail string `json:"primary_email" db:"primary_email"`
	SecondaryEmail string `json:"secondary_email" db:"secondary_email"`
	PublicPhone string `json:"public_phone" db:"public_phone"`
	SecondaryPhone string `json:"secondary_phone" db:"secondary_phone"`
	OwnerName string `json:"owner_name" db:"owner_name"`
	OwnerPhone string `json:"owner_phone" db:"owner_phone"`
}
