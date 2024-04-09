package models

type BeerMstDTO struct {
	Type string `json:"type" db:"type"`
	Description string `json:"description" db:"description"`
}
