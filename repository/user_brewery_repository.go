package repository

import (
	"context"

	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type UserBreweryVisitRepository struct {
	db *driver.DB
}

func NewUserBreweryVisitRepository(database *driver.DB) *UserBreweryVisitRepository {
	return &UserBreweryVisitRepository{
		db: database,
	}
}

// CreateUserBreweryVisit creates a new brewery visit entry for a user.
func (repo *UserBreweryVisitRepository) CreateUserBreweryVisit(userBreweryVisit models.UserBreweryVisit) error {
	query := `
		INSERT INTO user_brewery_visits (user_id, brewery_id, visit_date, visit_time)
		VALUES ($1, $2, $3, $4)
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, userBreweryVisit.UserID, userBreweryVisit.BreweryID, userBreweryVisit.VisitDate, userBreweryVisit.VisitTime)
	if err != nil {
		return err
	}
	return nil
}

// GetUserBreweryVisitsByUserID fetches a list of brewery visits by a specific user.
func (repo *UserBreweryVisitRepository) GetUserBreweryVisitsByUserID(userId string) ([]models.UserBreweryVisit, error) {
	var userBreweryVisits []models.UserBreweryVisit
	query := `
		SELECT id, user_id, brewery_id, visit_date, visit_time
		FROM user_brewery_visits
		WHERE user_id = $1
	`
	rows, err := repo.db.Pool.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userBreweryVisit models.UserBreweryVisit
		err := rows.Scan(&userBreweryVisit.ID, &userBreweryVisit.UserID, &userBreweryVisit.BreweryID, &userBreweryVisit.VisitDate, &userBreweryVisit.VisitTime)
		if err != nil {
			return nil, err
		}
		userBreweryVisits = append(userBreweryVisits, userBreweryVisit)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return userBreweryVisits, nil
}

// DeleteUserBreweryVisit deletes a user's brewery visit entry.
func (repo *UserBreweryVisitRepository) DeleteUserBreweryVisit(userId, visitId string) error {
	query := `
		DELETE FROM user_brewery_visits
		WHERE user_id = $1 AND id = $2
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, userId, visitId)
	if err != nil {
		return err
	}
	return nil
}
