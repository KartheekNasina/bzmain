package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type BrewTourService struct {
	brewTourRepo *repository.BrewTourRepository
}

func NewBrewTourService(r *repository.BrewTourRepository) *BrewTourService {
	return &BrewTourService{brewTourRepo: r}
}

// BrewTourService.go
func (bts *BrewTourService) GetTours(limit, offset int) ([]dto.BrewTourDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetTours",
	}).Debug("Get Tours - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetTours",
	}).Debug("Get Tours - End")

	tours, err := bts.brewTourRepo.GetTours(limit, offset)
	if err != nil {
		return nil, err
	}

	var tourDtos []dto.BrewTourDTO
	copier.Copy(&tourDtos, &tours)

	return tourDtos, nil
}

// ... (The previous service functions you provided for Update, Delete, and Create are also included here.)

// BrewTourService
func (bts *BrewTourService) UpdateUserBrewTourRegistration(id string, registrationDTO *dto.UserBrewTourRegistrationDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserBrewTourRegistration",
	}).Debug("Update User Brew Tour Registration - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserBrewTourRegistration",
	}).Debug("Update User Brew Tour Registration - End")

	var registration models.UserBrewTourRegistration
	copier.Copy(&registration, registrationDTO)

	return bts.brewTourRepo.UpdateUserBrewTourRegistration(id, registration)
}

func (bts *BrewTourService) DeleteUserBrewTourRegistration(registrationID string) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserBrewTourRegistration",
	}).Debug("Delete User Brew Tour Registration - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserBrewTourRegistration",
	}).Debug("Delete User Brew Tour Registration - End")

	return bts.brewTourRepo.DeleteUserBrewTourRegistration(registrationID)
}

func (bts *BrewTourService) CreateUserBrewTourRegistration(registrationDTO *dto.UserBrewTourRegistrationDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserBrewTourRegistration",
	}).Debug("Create User Brew Tour Registration - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserBrewTourRegistration",
	}).Debug("Create User Brew Tour Registration - End")

	var registration models.UserBrewTourRegistration
	copier.Copy(&registration, registrationDTO)

	return bts.brewTourRepo.CreateUserBrewTourRegistration(registration)
}
