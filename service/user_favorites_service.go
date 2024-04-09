package service

import (
	"github.com/jinzhu/copier"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type UserFavoriteBreweryService struct {
	userFavoriteBreweryRepo *repository.UserFavoriteBreweryRepository
}

func NewUserFavoriteBreweryService(r *repository.UserFavoriteBreweryRepository) *UserFavoriteBreweryService {
	return &UserFavoriteBreweryService{userFavoriteBreweryRepo: r}
}

// GetUserFavoriteBreweryByID retrieves a user's favorite brewery by ID.
func (us *UserFavoriteBreweryService) GetUserFavoriteBreweryByID(id string) (*dto.UserFavoriteBreweryDTO, error) {
	favoriteBrewery, err := us.userFavoriteBreweryRepo.GetUserFavoriteBrewery(id)
	if err != nil {
		return nil, err
	}

	var favoriteBreweryDTO dto.UserFavoriteBreweryDTO
	copier.Copy(&favoriteBreweryDTO, favoriteBrewery)
	// You can also copy other fields as needed.

	return &favoriteBreweryDTO, nil
}

// CreateUserFavoriteBrewery creates a new user's favorite brewery entry.
func (us *UserFavoriteBreweryService) CreateUserFavoriteBrewery(favoriteBreweryDTO *dto.UserFavoriteBreweryDTO) error {
	// Copy data from favoriteBreweryDTO to the UserFavoriteBrewery model
	var favoriteBrewery models.UserFavoriteBrewery
	copier.Copy(&favoriteBrewery, favoriteBreweryDTO)

	err := us.userFavoriteBreweryRepo.CreateUserFavoriteBrewery(favoriteBrewery)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserFavoriteBrewery deletes a user's favorite brewery entry by ID.
func (us *UserFavoriteBreweryService) DeleteUserFavoriteBrewery(id string, breweryID string) error {
	err := us.userFavoriteBreweryRepo.DeleteUserFavoriteBrewery(id, breweryID)
	if err != nil {
		return err
	}

	return nil
}
