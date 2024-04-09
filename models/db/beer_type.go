// beerType.go
package models

type BeerType struct {
	Type        string `db:"type"`
	Description string `db:"description"`
}
