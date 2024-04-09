package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type BreweryService struct {
	BreweryRepo *repository.BreweryRepository
}

// NewBreweryService creates a new BreweryService with the given BreweryRepository.
func NewBreweryService(r *repository.BreweryRepository) *BreweryService {
	return &BreweryService{BreweryRepo: r}
}

// GetBrewery returns the details of a specific brewery by ID.
func (s *BreweryService) GetBrewery(breweryID string) (dto.BreweryDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetBrewery",
	}).Debug("Get Brewery Details - Start")

	// Using defer to ensure the log is printed regardless of where we exit the function.
	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetBrewery",
	}).Debug("Get Brewery Details - End")

	Brewery, err := s.BreweryRepo.GetBrewery(breweryID)
	if err != nil {
		return dto.BreweryDTO{}, err
	}

	var BreweryDTO dto.BreweryDTO
	copier.Copy(&BreweryDTO, &Brewery)

	return BreweryDTO, nil
}

// GetBreweries returns a list of breweries with pagination support.
func (s *BreweryService) GetBreweries(limit, offset int) ([]dto.BreweryDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetBreweries",
	}).Debug("Get Breweries - Start")

	// Using defer to ensure the log is printed regardless of where we exit the function.
	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetBreweries",
	}).Debug("Get Breweries - End")

	Brewery, err := s.BreweryRepo.GetBreweries(limit, offset)
	if err != nil {
		return nil, err
	}

	var BreweryDTOs []dto.BreweryDTO
	copier.Copy(&BreweryDTOs, &Brewery)

	return BreweryDTOs, nil

}
