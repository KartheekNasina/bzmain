package service

import (
	"github.com/jinzhu/copier"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type UserFeedbackService struct {
	userFeedbackRepo *repository.UserFeedbackRepository
}

func NewUserFeedbackService(r *repository.UserFeedbackRepository) *UserFeedbackService {
	return &UserFeedbackService{userFeedbackRepo: r}
}

// CreateUserFeedback creates a new user feedback entry.
func (us *UserFeedbackService) CreateUserFeedback(feedbackDTO *dto.UserFeedbackDTO) error {
	// Copy data from feedbackDTO to the UserFeedback model
	var userFeedback models.UserFeedback
	copier.Copy(&userFeedback, feedbackDTO)

	err := us.userFeedbackRepo.CreateUserFeedback(userFeedback)
	if err != nil {
		return err
	}

	return nil
}

// GetUserFeedbackByID retrieves user feedback by ID.
func (us *UserFeedbackService) GetUserFeedbackByID(feedbackID string) (*dto.UserFeedbackDTO, error) {
	userFeedback, err := us.userFeedbackRepo.GetUserFeedbackByID(feedbackID)
	if err != nil {
		return nil, err
	}

	var userFeedbackDTO dto.UserFeedbackDTO
	copier.Copy(&userFeedbackDTO, userFeedback)
	// You can also copy other fields as needed.

	return &userFeedbackDTO, nil
}

// UpdateUserFeedback updates an existing user feedback entry.
func (us *UserFeedbackService) UpdateUserFeedback(feedbackID string, feedbackDTO *dto.UserFeedbackDTO) error {
	// Copy data from feedbackDTO to the UserFeedback model
	var userFeedback models.UserFeedback
	copier.Copy(&userFeedback, feedbackDTO)

	err := us.userFeedbackRepo.UpdateUserFeedback(feedbackID, userFeedback)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserFeedback deletes a user feedback entry by ID.
func (us *UserFeedbackService) DeleteUserFeedback(feedbackID string) error {
	err := us.userFeedbackRepo.DeleteUserFeedback(feedbackID)
	if err != nil {
		return err
	}

	return nil
}
