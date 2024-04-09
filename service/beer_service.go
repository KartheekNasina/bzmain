package service

import (
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type BeerService struct {
	beerRepo *repository.BeerRepository
}

func NewBeerService(r *repository.BeerRepository) *BeerService {
	return &BeerService{beerRepo: r}
}

func (us *BeerService) GetBeer(id string) (*models.Beer, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetBeer",
	}).Debug("Get Beer - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetBeer",
	}).Debug("Get Beer - End")

	beer, err := us.beerRepo.GetBeer(id)
	if err != nil {
		return nil, err
	}

	return beer, nil
}

func (us *BeerService) GetBeersBasedOnType(beerType string, limit, offset int) ([]models.Beer, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetBeersBasedOnType",
	}).Debug("Get Beers Based On Type - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetBeersBasedOnType",
	}).Debug("Get Beers Based On Type - End")

	beers, err := us.beerRepo.GetBeersBasedOnType(beerType, limit, offset)
	if err != nil {
		return nil, err
	}

	return beers, nil
}

func (us *BeerService) GetBeersBasedOnBreweryID(breweryID string, limit, offset int) ([]models.Beer, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetBeersBasedOnBreweryID",
	}).Debug("Get Beers Based On Brewery ID - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetBeersBasedOnBreweryID",
	}).Debug("Get Beers Based On Brewery ID - End")

	beers, err := us.beerRepo.GetBeersBasedOnBreweryID(breweryID, limit, offset)
	if err != nil {
		return nil, err
	}

	return beers, nil
}
