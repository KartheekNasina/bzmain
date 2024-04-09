package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

// UserBreweryCheckinRepository handles database operations for UserBreweryCheckin.
type UserBreweryCheckinRepository struct {
	db *driver.DB
}

// NewUserBreweryCheckinRepository creates a new UserBreweryCheckinRepository.
func NewUserBreweryCheckinRepository(database *driver.DB) *UserBreweryCheckinRepository {
	return &UserBreweryCheckinRepository{
		db: database,
	}
}

// CreateUserBreweryCheckin inserts a new user brewery check-in into the user_brewery_checkins table.
func (repo *UserBreweryCheckinRepository) CreateUserBreweryCheckin(userBreweryCheckin models.UserBreweryCheckin) error {
	// Replace "user_brewery_checkins" with your actual table name.
	query, args := common.GetInsertSQLAndArgs(&userBreweryCheckin, "user_brewery_checkins", "ID")
	fmt.Println(query)

	_, err := repo.db.Pool.Exec(context.Background(), query, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_brewery_checkin_repository",
			"function": "CreateUserBreweryCheckin",
		}).Errorf("Unable to create user brewery check-in. %v", err)

		return err
	}

	return nil
}

// UpdateUserBreweryCheckin updates the user brewery check-in in the user_brewery_checkins table.
func (repo *UserBreweryCheckinRepository) UpdateUserBreweryCheckin(id string, userBreweryCheckin models.UserBreweryCheckin) error {
	// Replace "user_brewery_checkins" with your actual table name.
	tableName := "user_brewery_checkins"
	idColumn := "id"

	userBreweryCheckin.CheckinDate = time.Now()

	query, values, err := common.GenerateUpdateQuery(tableName, idColumn, id, userBreweryCheckin)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_brewery_checkin_repository",
			"function": "UpdateUserBreweryCheckin",
		}).Errorf("Error generating update query: %v", err)
		return err
	}

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Values:", values)

	resp, err := repo.db.Pool.Exec(context.Background(), query, values...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_brewery_checkin_repository",
			"function": "UpdateUserBreweryCheckin",
		}).Errorf("Unable to update user brewery check-in. %v", err)
		return err
	}

	// Check the affected rows to determine the outcome
	affectedRows := resp.RowsAffected()
	if affectedRows == 0 {
		message := fmt.Sprintf("No user brewery check-in found with the ID: %s", id)
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_brewery_checkin_repository",
			"function": "UpdateUserBreweryCheckin",
		}).Debug(message)
		return fmt.Errorf(message)
	} else if affectedRows > 0 {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_brewery_checkin_repository",
			"function": "UpdateUserBreweryCheckin",
		}).Debug("User brewery check-in update successful")
	}

	return nil
}

// DeleteUserBreweryCheckin deletes a user brewery check-in by ID from the user_brewery_checkins table.
func (repo *UserBreweryCheckinRepository) DeleteUserBreweryCheckin(id string) error {
	// Replace "user_brewery_checkins" with your actual table name.
	query := "DELETE FROM user_brewery_checkins WHERE id = $1"

	result, err := repo.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_brewery_checkin_repository",
			"function": "DeleteUserBreweryCheckin",
		}).Errorf("Unable to delete user brewery check-in. %v", err)

		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("no entry found with the given ID: %s", id)
	}
	return nil
}

// ListUserBreweryCheckins fetches a list of user brewery check-ins.
func (repo *UserBreweryCheckinRepository) ListUserBreweryCheckins() ([]models.UserBreweryCheckin, error) {
	var userBreweryCheckins []models.UserBreweryCheckin

	// Replace "user_brewery_checkins" with your actual table name.
	query := "SELECT id, user_id, brewery_id, checkin_date, checkin_time, notes FROM user_brewery_checkins"
	rows, err := repo.db.Pool.Query(context.Background(), query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_brewery_checkin_repository",
			"function": "ListUserBreweryCheckins",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userBreweryCheckin models.UserBreweryCheckin
		err := common.ScanRowsToModel(rows, &userBreweryCheckin)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "user_brewery_checkin_repository",
				"function": "ListUserBreweryCheckins",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		userBreweryCheckins = append(userBreweryCheckins, userBreweryCheckin)
	}

	return userBreweryCheckins, nil
}
