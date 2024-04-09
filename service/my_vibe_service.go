package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type MyVibeService struct {
	myVibeRepo *repository.MyVibeRepository
}

func NewMyVibeService(r *repository.MyVibeRepository) *MyVibeService {
	return &MyVibeService{myVibeRepo: r}
}

func (vs *MyVibeService) CreateFoodDrinkRating(foodDrinkRatingDTO *dto.FoodDrinkRatingDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateFoodDrinkRating",
	}).Debug("Create FoodDrinkRating - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateFoodDrinkRating",
	}).Debug("Create FoodDrinkRating - End")

	var rating models.FoodDrinkRating
	if err := copier.Copy(&rating, foodDrinkRatingDTO); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "bz-main",
			"event":   "CreateFoodDrinkRating",
		}).Errorf("Failed to copy data from DTO to model: %v", err)
		return err
	}

	err := vs.myVibeRepo.CreateFoodDrinkRating(rating)
	if err != nil {
		return err
	}

	return nil
}

func (vs *MyVibeService) GetFoodDrinkRatingByID(id string) (dto.FoodDrinkRatingDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetFoodDrinkRatingByID",
	}).Debug("Get FoodDrinkRating by ID - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetFoodDrinkRatingByID",
	}).Debug("Get FoodDrinkRating by ID - End")

	rating, err := vs.myVibeRepo.GetFoodDrinkRatingByID(id)
	if err != nil {
		return dto.FoodDrinkRatingDTO{}, err
	}

	var ratingDTO dto.FoodDrinkRatingDTO
	copier.Copy(&ratingDTO, &rating)

	return ratingDTO, nil
}

// ... (repeat similar structure for Update, Delete, and GetFoodDrinkItems functions)
