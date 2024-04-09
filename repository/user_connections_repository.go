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

// UserConnectionRepository handles database operations for UserConnection.
type UserConnectionRepository struct {
	db *driver.DB
}

// NewUserConnectionRepository creates a new UserConnectionRepository.
func NewUserConnectionRepository(database *driver.DB) *UserConnectionRepository {
	return &UserConnectionRepository{
		db: database,
	}
}

// CreateUserConnection inserts a new user connection into the user_connections table.
func (repo *UserConnectionRepository) CreateUserConnection(userConnection models.UserConnection) error {
	// Replace "user_connections" with your actual table name.
	query, args := common.GetInsertSQLAndArgs(&userConnection, "user_connections", "ID")
	fmt.Println(query)

	_, err := repo.db.Pool.Exec(context.Background(), query, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_connection_repository",
			"function": "CreateUserConnection",
		}).Errorf("Unable to create user connection. %v", err)

		return err
	}

	return nil
}

// UpdateUserConnection updates the user connection in the user_connections table.
func (repo *UserConnectionRepository) UpdateUserConnection(id string, userConnection models.UserConnection) error {
	// Replace "user_connections" with your actual table name.
	tableName := "user_connections"
	idColumn := "id"

	userConnection.ConnectionDate = time.Now()

	query, values, err := common.GenerateUpdateQuery(tableName, idColumn, id, userConnection)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_connection_repository",
			"function": "UpdateUserConnection",
		}).Errorf("Error generating update query: %v", err)
		return err
	}

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Values:", values)

	resp, err := repo.db.Pool.Exec(context.Background(), query, values...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_connection_repository",
			"function": "UpdateUserConnection",
		}).Errorf("Unable to update user connection. %v", err)
		return err
	}

	// Check the affected rows to determine the outcome
	affectedRows := resp.RowsAffected()
	if affectedRows == 0 {
		message := fmt.Sprintf("No user connection found with the ID: %s", id)
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_connection_repository",
			"function": "UpdateUserConnection",
		}).Debug(message)
		return fmt.Errorf(message)
	} else if affectedRows > 0 {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_connection_repository",
			"function": "UpdateUserConnection",
		}).Debug("User connection update successful")
	}

	return nil
}

// DeleteUserConnection deletes a user connection by ID from the user_connections table.
func (repo *UserConnectionRepository) DeleteUserConnection(id string) error {
	// Replace "user_connections" with your actual table name.
	query := "DELETE FROM user_connections WHERE id = $1"

	result, err := repo.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_connection_repository",
			"function": "DeleteUserConnection",
		}).Errorf("Unable to delete user connection. %v", err)

		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("no entry found with the given ID: %s", id)
	}
	return nil
}

// ListUserConnections fetches a list of user connections.
func (repo *UserConnectionRepository) ListUserConnections() ([]models.UserConnection, error) {
	var userConnections []models.UserConnection

	// Replace "user_connections" with your actual table name.
	query := "SELECT id, user1_id, user2_id, connection_date FROM user_connections"
	rows, err := repo.db.Pool.Query(context.Background(), query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_connection_repository",
			"function": "ListUserConnections",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userConnection models.UserConnection
		err := common.ScanRowsToModel(rows, &userConnection)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "user_connection_repository",
				"function": "ListUserConnections",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		userConnections = append(userConnections, userConnection)
	}

	return userConnections, nil
}
