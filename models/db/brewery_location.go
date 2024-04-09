package models

import (
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
)

type BreweryLocation struct {
	BreweryID   string  `db:"brewery_id"`
	Geolocation geom.T  `db:"geolocation"`
	Address     string  `db:"address"`
	Lat         float64 `db:"lat"`
	Lng         float64 `db:"lng"`
	CountryID   string  `db:"country_id"`
	CityID      string  `db:"city_id"`
}

// MarshalGeolocation marshals the Geolocation field to Well-Known Binary (WKB) format.
func (bl *BreweryLocation) MarshalGeolocation() ([]byte, error) {
	return wkb.Marshal(bl.Geolocation, wkb.NDR)
}

// UnmarshalGeolocation unmarshals Well-Known Binary (WKB) data to the Geolocation field.
func (bl *BreweryLocation) UnmarshalGeolocation(data []byte) error {
	geom, err := wkb.Unmarshal(data)
	if err != nil {
		return err
	}
	bl.Geolocation = geom
	return nil
}
