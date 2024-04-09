// repository/user_repository.go
package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	common "github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type UserRepository struct {
	db *driver.DB
}

func NewUserRepository(database *driver.DB) *UserRepository {
	return &UserRepository{
		db: database,
	}
}

// GetUsers fetches a list of users with pagination.
func (repo *UserRepository) GetUsers(limit, offset int) ([]models.User, error) {
	var users []models.User

	// Replace "users" with your actual table name.
	query := `
		SELECT id, user_id, name, about_me, gender, email, profile_url, images, dob, is_online, allow_notifications, allow_location, phone_number, is_legal_age, provider, brew_interests, personal_interests, phone_number_verified, phone_number_verification_date, onboarding_status, otp_verified_at, profile_completed_at, created_at, updated_at

		FROM users
		LIMIT $1 OFFSET $2
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_repository",
			"function": "GetUsers",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := common.ScanRowsToModel(rows, &user)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "user_repository",
				"function": "ListUsers",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// UpdateUser updates the user's information in the users table.
func (repo *UserRepository) UpdateUser(id string, user models.User) error {
	// Replace "users" with your actual table name.
	tableName := "users"
	idColumn := "id"

	user.UpdatedAt = time.Now()

	query, values, err := common.GenerateUpdateQuery(tableName, idColumn, id, user)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_repository",
			"function": "UpdateUser",
		}).Errorf("Error generating update query: %v", err)
		return err
	}

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Values:", values)

	resp, err := repo.db.Pool.Exec(context.Background(), query, values...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_repository",
			"function": "UpdateUser",
		}).Errorf("Unable to update user. %v", err)
		return err
	}

	// Check the affected rows to determine the outcome
	affectedRows := resp.RowsAffected()
	if affectedRows == 0 {
		message := fmt.Sprintf("No user found with the ID: %s", id)
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_repository",
			"function": "UpdateUser",
		}).Debug(message)
		return fmt.Errorf(message)
	} else if affectedRows > 0 {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_repository",
			"function": "UpdateUser",
		}).Debug("User update successful")
	}

	return nil
}

// DeleteUser deletes a user by ID from the users table.
func (repo *UserRepository) DeleteUser(userID string) error {
	// Replace "users" with your actual table name.
	query := "DELETE FROM users WHERE id = $1"

	result, err := repo.db.Pool.Exec(context.Background(), query, userID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_repository",
			"function": "DeleteUser",
		}).Errorf("Unable to delete user. %v", err)

		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("no entry found with the given userID")
	}
	return nil
}

func (repo *UserRepository) emailExists(email string) (bool, error) {
	var existingUserID uuid.UUID
	err := repo.db.Pool.QueryRow(context.Background(), "SELECT ID FROM users WHERE email=$1", email).Scan(&existingUserID)

	// If no rows are returned, it's not an error scenario for our use case.
	if err != nil && err.Error() == "no rows in result set" {
		return false, nil
	} else if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_repository",
			"function": "emailExists",
		}).Errorf("Error checking for existing user. %v", err)
		return false, err
	}

	return existingUserID != uuid.Nil, nil
}

func (repo *UserRepository) CreateUser(user models.User) error {
	exists, err := repo.emailExists(user.Email)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("User with email %s already exists", user.Email)
	}

	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()

	query, args := common.GetInsertSQLAndArgs(&user, "users", "ID")
	fmt.Println(query)

	_, err = repo.db.Pool.Exec(context.Background(), query, args...)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_repository",
			"function": "CreateUser",
		}).Errorf("Unable to create user. %v", err)

		return err
	}

	return nil
}

// ListUsers fetches a list of users.
func (repo *UserRepository) ListUsers() ([]models.User, error) {
	var users []models.User

	// Replace "users" with your actual table name.
	query := "SELECT id, user_id, name, about_me, gender, email, profile_url, images, dob, is_online, allow_notifications, allow_location, phone_number, is_legal_age, provider, brew_interests, personal_interests, phone_number_verified, phone_number_verification_date, onboarding_status, otp_verified_at, profile_completed_at, created_at, updated_at FROM users"
	rows, err := repo.db.Pool.Query(context.Background(), query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_repository",
			"function": "ListUsers",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := common.ScanRowsToModel(rows, &user)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "user_repository",
				"function": "ListUsers",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
