package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type BrewClassService struct {
	brewClassRepo *repository.BrewClassRepository
}

func NewBrewClassService(r *repository.BrewClassRepository) *BrewClassService {
	return &BrewClassService{brewClassRepo: r}
}

func (us *BrewClassService) GetBrewClassWithSchedule(id string) (*dto.BrewClassDTO, error) {
	// logrus.WithFields(logrus.Fields{
	// 	"service": "bz-main",
	// 	"event":   "GetBrewClassWithSchedule",
	// }).Debug("Get Brew Class With Schedule - Start")

	// defer logrus.WithFields(logrus.Fields{
	// 	"service": "bz-main",
	// 	"event":   "GetBrewClassWithSchedule",
	// }).Debug("Get Brew Class With Schedule - End")

	// basicInfo, schedules, err := us.brewClassRepo.GetBrewClassWithSchedule(id)
	// if err != nil {
	// 	return nil, err
	// }

	// var dtoBrewClass dto.BrewClassDTO
	// copier.Copy(&dtoBrewClass, basicInfo)
	// // You can also copy other fields as needed.

	// dtoBrewClass.Schedules = make([]*dto.BrewClassesScheduleDTO, len(schedules))
	// for i, schedule := range schedules {
	// 	var dtoSchedule dto.BrewClassesScheduleDTO
	// 	copier.Copy(&dtoSchedule, schedule)
	// 	// You can also copy other fields as needed.
	// 	dtoBrewClass.Schedules[i] = &dtoSchedule
	// }

	// return &dtoBrewClass, nil
	return nil, nil
}

func (us *BrewClassService) UpdateUserBrewClassRegistration(id string, registrationDTO *dto.UserBrewClassRegistrationDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserBrewClassRegistration",
	}).Debug("Update User Brew Class Registration - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserBrewClassRegistration",
	}).Debug("Update User Brew Class Registration - End")

	// Copy data from registrationDTO to the registration model
	var registration models.UserBrewClassRegistration
	copier.Copy(&registration, registrationDTO)

	err := us.brewClassRepo.UpdateUserBrewClassRegistration(id, registration)
	if err != nil {
		return err
	}

	return nil
}

func (us *BrewClassService) DeleteUserBrewClassRegistration(registrationID string) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserBrewClassRegistration",
	}).Debug("Delete User Brew Class Registration - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserBrewClassRegistration",
	}).Debug("Delete User Brew Class Registration - End")

	err := us.brewClassRepo.DeleteUserBrewClassRegistration(registrationID)
	if err != nil {
		return err
	}

	return nil
}

func (us *BrewClassService) CreateUserBrewClassRegistration(registrationDTO *dto.UserBrewClassRegistrationDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserBrewClassRegistration",
	}).Debug("Create User Brew Class Registration - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserBrewClassRegistration",
	}).Debug("Create User Brew Class Registration - End")

	// Copy data from registrationDTO to the registration model
	var registration models.UserBrewClassRegistration
	copier.Copy(&registration, registrationDTO)

	err := us.brewClassRepo.CreateUserBrewClassRegistration(registration)
	if err != nil {
		return err
	}

	return nil
}

func (us *BrewClassService) GetUserBrewClassRegistration(userID, classID string) (*dto.UserBrewClassRegistrationDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetUserBrewClassRegistration",
	}).Debug("Get User Brew Class Registration - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetUserBrewClassRegistration",
	}).Debug("Get User Brew Class Registration - End")

	registration, err := us.brewClassRepo.GetUserBrewClassRegistration(userID, classID)
	if err != nil {
		return nil, err
	}

	var registrationDTO dto.UserBrewClassRegistrationDTO
	copier.Copy(&registrationDTO, registration)
	// You can also copy other fields as needed.

	return &registrationDTO, nil
}
