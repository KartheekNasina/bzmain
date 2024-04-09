package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type CityService struct {
	cityRepo *repository.CityRepository
}

// NewCityService creates a new CityService with the given CityRepository.
func NewCityService(r *repository.CityRepository) *CityService {
	return &CityService{cityRepo: r}
}

// ListCities returns a list of all cities.
func (cs *CityService) ListCities() []models.CityDTO {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "ListCities",
	}).Debug("List Cities - Start")

	// Using defer to ensure the log is printed regardless of where we exit the function.
	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "ListCities",
	}).Debug("List Cities - End")

	cities, _ := cs.cityRepo.ListCities()

	// Convert cities to CityDTOs
	var cityDtos []models.CityDTO
	copier.Copy(&cityDtos, &cities)

	return cityDtos
}
