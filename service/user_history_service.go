package service

import (
	"github.com/jinzhu/copier"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type UserHistoryService struct {
	userHistoryRepo *repository.UserHistoryRepository
}

func NewUserHistoryService(r *repository.UserHistoryRepository) *UserHistoryService {
	return &UserHistoryService{userHistoryRepo: r}
}

func (us *UserHistoryService) GetUserHistory(limit, offset int) ([]*dto.UserHistoryDTO, error) {
	userHistories, err := us.userHistoryRepo.GetUserHistory(limit, offset)
	if err != nil {
		return nil, err
	}

	var userHistoryDTOs []*dto.UserHistoryDTO
	for _, userHistory := range userHistories {
		var userHistoryDTO dto.UserHistoryDTO
		copier.Copy(&userHistoryDTO, userHistory)
		// You can also copy other fields as needed.
		userHistoryDTOs = append(userHistoryDTOs, &userHistoryDTO)
	}

	return userHistoryDTOs, nil
}

func (us *UserHistoryService) GetUserHistoryByID(id string) (*dto.UserHistoryDTO, error) {
	userHistory, err := us.userHistoryRepo.GetUserHistoryByID(id)
	if err != nil {
		return nil, err
	}

	var userHistoryDTO dto.UserHistoryDTO
	copier.Copy(&userHistoryDTO, userHistory)
	// You can also copy other fields as needed.

	return &userHistoryDTO, nil
}

func (us *UserHistoryService) CreateUserHistory(historyDTO *dto.UserHistoryDTO) error {
	// Copy data from historyDTO to the UserHistory model
	var userHistory models.UserHistory
	copier.Copy(&userHistory, historyDTO)

	err := us.userHistoryRepo.CreateUserHistory(userHistory)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserHistoryService) UpdateUserHistory(id string, historyDTO *dto.UserHistoryDTO) error {
	// Copy data from historyDTO to the UserHistory model
	var userHistory models.UserHistory
	copier.Copy(&userHistory, historyDTO)

	err := us.userHistoryRepo.UpdateUserHistory(id, userHistory)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserHistoryService) DeleteUserHistory(id string) error {
	err := us.userHistoryRepo.DeleteUserHistory(id)
	if err != nil {
		return err
	}

	return nil
}
