package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

// UserFavoriteBreweryRepository handles database operations for UserFavoriteBrewery.
type UserFavoriteBreweryRepository struct {
	db *driver.DB
}

// NewUserFavoriteBreweryRepository creates a new UserFavoriteBreweryRepository.
func NewUserFavoriteBreweryRepository(database *driver.DB) *UserFavoriteBreweryRepository {
	return &UserFavoriteBreweryRepository{
		db: database,
	}
}

// CreateUserFavoriteBrewery inserts a new user favorite brewery into the database.
func (repo *UserFavoriteBreweryRepository) CreateUserFavoriteBrewery(favoriteBrewery models.UserFavoriteBrewery) error {
	query := `
		INSERT INTO user_favorite_breweries (user_id, brewery_id, added_date)
		VALUES ($1, $2, $3)
	`

	_, err := repo.db.Pool.Exec(context.Background(), query,
		favoriteBrewery.UserID, favoriteBrewery.BreweryID, favoriteBrewery.AddedDate)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_favorite_brewery_repository",
			"function": "CreateUserFavoriteBrewery",
		}).Errorf("Unable to create user favorite brewery. %v", err)

		return err
	}

	return nil
}

// GetUserFavoriteBrewery retrieves a user favorite brewery by user and brewery IDs.
func (repo *UserFavoriteBreweryRepository) GetUserFavoriteBrewery(userID string) (*models.UserFavoriteBrewery, error) {
	var favoriteBrewery models.UserFavoriteBrewery

	query := "SELECT user_id, brewery_id, added_date FROM user_favorite_breweries WHERE user_id = $1 "
	err := repo.db.Pool.QueryRow(context.Background(), query, userID).Scan(
		&favoriteBrewery.UserID, &favoriteBrewery.BreweryID, &favoriteBrewery.AddedDate)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_favorite_brewery_repository",
			"function": "GetUserFavoriteBrewery",
		}).Errorf("Unable to retrieve user favorite brewery. %v", err)

		return nil, err
	}

	return &favoriteBrewery, nil
}

// DeleteUserFavoriteBrewery deletes a user favorite brewery from the database by user and brewery IDs.
func (repo *UserFavoriteBreweryRepository) DeleteUserFavoriteBrewery(userID, breweryID string) error {
	query := "DELETE FROM user_favorite_breweries WHERE user_id = $1 AND brewery_id = $2"
	_, err := repo.db.Pool.Exec(context.Background(), query, userID, breweryID)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_favorite_brewery_repository",
			"function": "DeleteUserFavoriteBrewery",
		}).Errorf("Unable to delete user favorite brewery. %v", err)

		return err
	}

	return nil
}
