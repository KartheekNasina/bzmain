package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type UserConnectionService struct {
	userConnectionsRepo *repository.UserConnectionRepository
}

func NewUserConnectionsService(r *repository.UserConnectionRepository) *UserConnectionService {
	return &UserConnectionService{userConnectionsRepo: r}
}

func (ucs *UserConnectionService) CreateUserConnection(userConnectionDTO *dto.UserConnectionDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserConnection",
	}).Debug("Create UserConnection - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserConnection",
	}).Debug("Create UserConnection - End")

	var userConnection models.UserConnection
	if err := copier.Copy(&userConnection, userConnectionDTO); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "bz-main",
			"event":   "CreateUserConnection",
		}).Errorf("Failed to copy data from userConnectionDTO to userConnection: %v", err)
		return err
	}

	return ucs.userConnectionsRepo.CreateUserConnection(userConnection)
}

func (ucs *UserConnectionService) UpdateUserConnection(id string, userConnectionDTO *dto.UserConnectionDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserConnection",
	}).Debug("Update UserConnection - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserConnection",
	}).Debug("Update UserConnection - End")

	var userConnection models.UserConnection
	if err := copier.Copy(&userConnection, userConnectionDTO); err != nil {
		return err
	}

	return ucs.userConnectionsRepo.UpdateUserConnection(id, userConnection)
}

func (ucs *UserConnectionService) DeleteUserConnection(id string) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserConnection",
	}).Debug("Delete UserConnection - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserConnection",
	}).Debug("Delete UserConnection - End")

	return ucs.userConnectionsRepo.DeleteUserConnection(id)
}
