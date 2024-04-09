package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type UserBlockedListService struct {
	userBlockedListRepo *repository.UserBlockedListRepository
}

func NewUserBlockedListService(r *repository.UserBlockedListRepository) *UserBlockedListService {
	return &UserBlockedListService{userBlockedListRepo: r}
}

func (us *UserBlockedListService) CreateUserBlocked(userBlockedDTO *dto.UserBlockedListDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserBlocked",
	}).Debug("Create UserBlocked - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserBlocked",
	}).Debug("Create UserBlocked - End")

	var userBlocked models.UserBlockedList
	if err := copier.Copy(&userBlocked, userBlockedDTO); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "bz-main",
			"event":   "CreateUserBlocked",
		}).Errorf("Failed to copy data from userBlockedDTO to userBlocked: %v", err)
		return err
	}

	err := us.userBlockedListRepo.CreateUserBlocked(userBlocked)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserBlockedListService) GetUserBlockedByUserID(userID string) ([]dto.UserBlockedListDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetUserBlockedByUserID",
	}).Debug("Get UserBlocked by User ID - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetUserBlockedByUserID",
	}).Debug("Get UserBlocked by User ID - End")

	userBlockeds, err := us.userBlockedListRepo.GetUserBlockedByUserID(userID)
	if err != nil {
		return nil, err
	}

	var userBlockedDtos []dto.UserBlockedListDTO
	copier.Copy(&userBlockedDtos, &userBlockeds)

	return userBlockedDtos, nil
}

func (us *UserBlockedListService) DeleteUserBlocked(userID, blockedUserID string) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserBlocked",
	}).Debug("Delete UserBlocked - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserBlocked",
	}).Debug("Delete UserBlocked - End")

	err := us.userBlockedListRepo.DeleteUserBlocked(userID, blockedUserID)
	if err != nil {
		return err
	}

	return nil
}
