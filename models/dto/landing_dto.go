package models

type LandingDTO struct {
	TrendingBreweries []BreweryLandingDTO `json:"trending_breweries"`
	Events            []EventDTO          `json:"events"`
	BrewClasses       []BrewClassDTO      `json:"brew_classes"`
	BrewTours         []BrewTourDTO       `json:"brew_tours"`
}
