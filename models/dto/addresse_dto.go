package models

type AddresseDTO struct {
	Id string `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	Street string `json:"street" db:"street"`
	City string `json:"city" db:"city"`
	State string `json:"state" db:"state"`
	PostalCode string `json:"postal_code" db:"postal_code"`
}
