package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type UserBreweryCheckinsService struct {
	userBreweryCheckinsRepo *repository.UserBreweryCheckinRepository
}

func NewUserBreweryCheckinsService(r *repository.UserBreweryCheckinRepository) *UserBreweryCheckinsService {
	return &UserBreweryCheckinsService{userBreweryCheckinsRepo: r}
}

func (ubs *UserBreweryCheckinsService) CreateUserBreweryCheckin(checkinDTO *dto.UserBreweryCheckinDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserBreweryCheckin",
	}).Debug("Create UserBreweryCheckin - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserBreweryCheckin",
	}).Debug("Create UserBreweryCheckin - End")

	var checkin models.UserBreweryCheckin
	if err := copier.Copy(&checkin, checkinDTO); err != nil {
		return err
	}

	err := ubs.userBreweryCheckinsRepo.CreateUserBreweryCheckin(checkin)
	if err != nil {
		return err
	}

	return nil
}

func (ubs *UserBreweryCheckinsService) UpdateUserBreweryCheckin(id string, checkinDTO *dto.UserBreweryCheckinDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserBreweryCheckin",
	}).Debug("Update UserBreweryCheckin - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserBreweryCheckin",
	}).Debug("Update UserBreweryCheckin - End")

	var checkin models.UserBreweryCheckin
	if err := copier.Copy(&checkin, checkinDTO); err != nil {
		return err
	}

	err := ubs.userBreweryCheckinsRepo.UpdateUserBreweryCheckin(id, checkin)
	if err != nil {
		return err
	}

	return nil
}

func (ubs *UserBreweryCheckinsService) DeleteUserBreweryCheckin(id string) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserBreweryCheckin",
	}).Debug("Delete UserBreweryCheckin - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserBreweryCheckin",
	}).Debug("Delete UserBreweryCheckin - End")

	err := ubs.userBreweryCheckinsRepo.DeleteUserBreweryCheckin(id)
	if err != nil {
		return err
	}

	return nil
}

func (ubs *UserBreweryCheckinsService) ListUserBreweryCheckins() ([]dto.UserBreweryCheckinDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "ListUserBreweryCheckins",
	}).Debug("List UserBreweryCheckins - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "ListUserBreweryCheckins",
	}).Debug("List UserBreweryCheckins - End")

	checkins, err := ubs.userBreweryCheckinsRepo.ListUserBreweryCheckins()
	if err != nil {
		return nil, err
	}

	var checkinDTOs []dto.UserBreweryCheckinDTO
	copier.Copy(&checkinDTOs, &checkins)

	return checkinDTOs, nil
}
