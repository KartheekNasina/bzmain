// user_history_repository.go
package repository

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type UserHistoryRepository struct {
	db *driver.DB
}

func NewUserHistoryRepository(database *driver.DB) *UserHistoryRepository {
	return &UserHistoryRepository{
		db: database,
	}
}

// GetUserHistory fetches a list of user history with limit and offset.
func (repo *UserHistoryRepository) GetUserHistory(limit, offset int) ([]models.UserHistory, error) {
	var userHistory []models.UserHistory

	// Replace "user_history" with your actual table name.
	query := `
		SELECT *
		FROM user_history
		LIMIT $1 OFFSET $2
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_history_repository",
			"function": "GetUserHistory",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var history models.UserHistory

		err = rows.Scan(
			&history.ID, &history.UserID, &history.InteractionID, &history.InteractionType,
			&history.InteractionDate, &history.Notes,
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "user_history_repository",
				"function": "GetUserHistory",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		userHistory = append(userHistory, history)
	}

	return userHistory, nil
}

// GetUserHistoryByID fetches a user history based on its ID.
func (repo *UserHistoryRepository) GetUserHistoryByID(id string) (*models.UserHistory, error) {
	var history models.UserHistory

	// Replace "user_history" with your actual table name.
	query := `
		SELECT *
		FROM user_history
		WHERE id = $1
		LIMIT 1
	`

	row := repo.db.Pool.QueryRow(context.Background(), query, id)

	err := row.Scan(
		&history.ID, &history.UserID, &history.InteractionID, &history.InteractionType,
		&history.InteractionDate, &history.Notes,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_history_repository",
			"function": "GetUserHistoryByID",
		}).Errorf("Unable to scan the row. %v", err)

		return nil, err
	}

	return &history, nil
}

// CreateUserHistory inserts a new user history into the user_history table.
func (repo *UserHistoryRepository) CreateUserHistory(history models.UserHistory) error {
	// Replace "user_history" with your actual table name.
	query := `
		INSERT INTO user_history
		(user_id, interaction_id, interaction_type, interaction_date, notes)
		VALUES
		($1, $2, $3, $4, $5)
	`

	_, err := repo.db.Pool.Exec(context.Background(), query,
		history.UserID, history.InteractionID, history.InteractionType, history.InteractionDate, history.Notes)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_history_repository",
			"function": "CreateUserHistory",
		}).Errorf("Unable to create user history. %v", err)

		return err
	}

	return nil
}

// UpdateUserHistory updates user history information in the user_history table.
func (repo *UserHistoryRepository) UpdateUserHistory(id string, history models.UserHistory) error {
	// Replace "user_history" with your actual table name.
	tableName := "user_history"
	idColumn := "id"

	query, values, err := common.GenerateUpdateQuery(tableName, idColumn, id, history)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_history_repository",
			"function": "UpdateUserHistory",
		}).Errorf("Error generating update query: %v", err)
		return err
	}

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Values:", values)

	resp, err := repo.db.Pool.Exec(context.Background(), query, values...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_history_repository",
			"function": "UpdateUserHistory",
		}).Errorf("Unable to update user history. %v", err)
		return err
	}

	// Check the affected rows to determine the outcome
	affectedRows := resp.RowsAffected()
	if affectedRows == 0 {
		message := fmt.Sprintf("No user history found with the ID: %s", id)
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_history_repository",
			"function": "UpdateUserHistory",
		}).Debug(message)
		return fmt.Errorf(message)
	} else if affectedRows > 0 {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_history_repository",
			"function": "UpdateUserHistory",
			"id":       id,
		}).Debug("User history update successful")
	}

	return nil
}

// DeleteUserHistory deletes user history by ID from the user_history table.
func (repo *UserHistoryRepository) DeleteUserHistory(id string) error {
	// Replace "user_history" with your actual table name.
	query := "DELETE FROM user_history WHERE id = $1"

	_, err := repo.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_history_repository",
			"function": "DeleteUserHistory",
		}).Errorf("Unable to delete user history. %v", err)

		return err
	}

	return nil
}
