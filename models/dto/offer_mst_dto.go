package models

type OfferMstDTO struct {
	Type string `json:"type" db:"type"`
	Description string `json:"description" db:"description"`
}
