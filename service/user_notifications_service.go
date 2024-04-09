package service

import (
	"github.com/jinzhu/copier"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type UserNotificationService struct {
	userNotificationRepo *repository.UserNotificationRepository
}

func NewUserNotificationService(r *repository.UserNotificationRepository) *UserNotificationService {
	return &UserNotificationService{userNotificationRepo: r}
}

func (us *UserNotificationService) CreateUserNotification(notificationDTO *dto.UserNotificationDTO) error {
	// Copy data from notificationDTO to the UserNotification model
	var userNotification models.UserNotification
	copier.Copy(&userNotification, notificationDTO)

	err := us.userNotificationRepo.CreateUserNotification(userNotification)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserNotificationService) GetUserNotificationByID(notificationID string) (*dto.UserNotificationDTO, error) {
	userNotification, err := us.userNotificationRepo.GetUserNotificationByID(notificationID)
	if err != nil {
		return nil, err
	}

	var userNotificationDTO dto.UserNotificationDTO
	copier.Copy(&userNotificationDTO, userNotification)
	// You can also copy other fields as needed.

	return &userNotificationDTO, nil
}

func (us *UserNotificationService) UpdateUserNotification(notificationDTO *dto.UserNotificationDTO) error {
	// Copy data from notificationDTO to the UserNotification model
	var userNotification models.UserNotification
	copier.Copy(&userNotification, notificationDTO)

	err := us.userNotificationRepo.UpdateUserNotification(userNotification)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserNotificationService) DeleteUserNotification(notificationID string) error {
	err := us.userNotificationRepo.DeleteUserNotification(notificationID)
	if err != nil {
		return err
	}

	return nil
}
