package models

type FoodDrinkItemDTO struct {
	Id string `json:"id" db:"id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	Type string `json:"type" db:"type"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	ImageUrl string `json:"image_url" db:"image_url"`
}
