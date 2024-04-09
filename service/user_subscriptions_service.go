package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type UserSubscriptionService struct {
	userSubscriptionRepo *repository.UserSubscriptionRepository
}

func NewUserSubscriptionService(r *repository.UserSubscriptionRepository) *UserSubscriptionService {
	return &UserSubscriptionService{userSubscriptionRepo: r}
}

func (us *UserSubscriptionService) CreateUserSubscription(subscriptionDTO *dto.UserSubscriptionDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserSubscription",
	}).Debug("Create User Subscription - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserSubscription",
	}).Debug("Create User Subscription - End")

	// Copy data from subscriptionDTO to the UserSubscription model
	var userSubscription models.UserSubscription
	copier.Copy(&userSubscription, subscriptionDTO)

	err := us.userSubscriptionRepo.CreateUserSubscription(userSubscription)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserSubscriptionService) GetUserSubscriptionByUserID(userID string) (*dto.UserSubscriptionDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetUserSubscriptionByUserID",
	}).Debug("Get User Subscription By UserID - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetUserSubscriptionByUserID",
	}).Debug("Get User Subscription By UserID - End")

	userSubscription, err := us.userSubscriptionRepo.GetUserSubscriptionByUserID(userID)
	if err != nil {
		return nil, err
	}

	var subscriptionDTO dto.UserSubscriptionDTO
	copier.Copy(&subscriptionDTO, userSubscription)
	// You can also copy other fields as needed.

	return &subscriptionDTO, nil
}

func (us *UserSubscriptionService) UpdateUserSubscription(userID string, subscriptionDTO *dto.UserSubscriptionDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserSubscription",
	}).Debug("Update User Subscription - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserSubscription",
	}).Debug("Update User Subscription - End")

	// Copy data from subscriptionDTO to the UserSubscription model
	var userSubscription models.UserSubscription
	copier.Copy(&userSubscription, subscriptionDTO)

	err := us.userSubscriptionRepo.UpdateUserSubscription(userID, userSubscription)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserSubscriptionService) DeleteUserSubscription(userID string) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserSubscription",
	}).Debug("Delete User Subscription - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserSubscription",
	}).Debug("Delete User Subscription - End")

	err := us.userSubscriptionRepo.DeleteUserSubscription(userID)
	if err != nil {
		return err
	}

	return nil
}
