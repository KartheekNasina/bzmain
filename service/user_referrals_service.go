package service

import (
	"github.com/jinzhu/copier"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type UserReferralService struct {
	userReferralRepo *repository.UserReferralRepository
}

func NewUserReferralService(r *repository.UserReferralRepository) *UserReferralService {
	return &UserReferralService{userReferralRepo: r}
}

func (us *UserReferralService) GetUserReferral(limit, offset int) ([]*dto.UserReferralDTO, error) {
	userReferrals, err := us.userReferralRepo.GetUserReferral(limit, offset)
	if err != nil {
		return nil, err
	}

	var userReferralDTOs []*dto.UserReferralDTO
	for _, userReferral := range userReferrals {
		var userReferralDTO dto.UserReferralDTO
		copier.Copy(&userReferralDTO, userReferral)
		// You can also copy other fields as needed.
		userReferralDTOs = append(userReferralDTOs, &userReferralDTO)
	}

	return userReferralDTOs, nil
}

func (us *UserReferralService) GetUserReferralByID(id string) (*dto.UserReferralDTO, error) {
	userReferral, err := us.userReferralRepo.GetUserReferralByID(id)
	if err != nil {
		return nil, err
	}

	var userReferralDTO dto.UserReferralDTO
	copier.Copy(&userReferralDTO, userReferral)
	// You can also copy other fields as needed.

	return &userReferralDTO, nil
}

func (us *UserReferralService) CreateUserReferral(referralDTO *dto.UserReferralDTO) error {
	// Copy data from referralDTO to the UserReferral model
	var userReferral models.UserReferral
	copier.Copy(&userReferral, referralDTO)

	err := us.userReferralRepo.CreateUserReferral(userReferral)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserReferralService) UpdateUserReferral(id string, referralDTO *dto.UserReferralDTO) error {
	// Copy data from referralDTO to the UserReferral model
	var userReferral models.UserReferral
	copier.Copy(&userReferral, referralDTO)

	err := us.userReferralRepo.UpdateUserReferral(id, userReferral)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserReferralService) DeleteUserReferral(id string) error {
	err := us.userReferralRepo.DeleteUserReferral(id)
	if err != nil {
		return err
	}

	return nil
}
