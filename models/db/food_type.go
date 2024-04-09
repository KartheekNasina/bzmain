// foodType.go
package models

type FoodType struct {
	Type  string `db:"type"`
	Title string `db:"title"`
}
