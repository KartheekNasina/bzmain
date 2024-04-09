// foodDrinkItem.go
package models

type FoodDrinkItem struct {
	ID          string   `db:"id"`
	BreweryID   string   `db:"brewery_id"`
	Name        string   `db:"name"`
	Type        ItemType `db:"type"`
	Description string   `db:"description"`
	ImageURL    string   `db:"image_url"`
}
