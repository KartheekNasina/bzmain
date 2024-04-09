package models

type FoodTypeDTO struct {
	Type string `json:"type" db:"type"`
	Title string `json:"title" db:"title"`
}
