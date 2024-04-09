package models

type BreweryLocationDTO struct {
	Lat string `json:"lat" db:"lat"`
	Geolocation string `json:"geolocation" db:"geolocation"`
	CityId string `json:"city_id" db:"city_id"`
	BreweryId string `json:"brewery_id" db:"brewery_id"`
	Lng string `json:"lng" db:"lng"`
	CountryId string `json:"country_id" db:"country_id"`
	Address string `json:"address" db:"address"`
}
