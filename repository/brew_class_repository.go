// repository/brew_classes_repository.go
package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	common "github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type BrewClassRepository struct {
	db *driver.DB
}

func NewBrewClassRepository(database *driver.DB) *BrewClassRepository {
	return &BrewClassRepository{
		db: database,
	}
}

// GetBrewClassWithSchedule fetches a BrewClass based on its ID along with its schedule.
func (repo *BrewClassRepository) GetBrewClassWithSchedule(id string) (*models.BrewClassBasicInfo, []*models.BrewClassSchedule, error) {
	var brewClassInfo models.BrewClassBasicInfo
	var brewClassSchedules []*models.BrewClassSchedule

	// Replace "brew_classes" with your actual table name.
	query := `
		SELECT *
		FROM brew_classes
		WHERE id = $1
	`
	row := repo.db.Pool.QueryRow(context.Background(), query, id)
	err := common.ScanRowToModel(row, &brewClassInfo)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_classes_repository",
			"function": "GetBrewClassWithSchedule",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, nil, err
	}

	// Fetch the schedule for the BrewClass
	// Replace "brew_class_schedules" with your actual table name.
	scheduleQuery := `
		SELECT *
		FROM brew_class_schedules
		WHERE class_id = $1
	`

	rows, err := repo.db.Pool.Query(context.Background(), scheduleQuery, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_classes_repository",
			"function": "GetBrewClassWithSchedule",
		}).Errorf("Unable to execute the schedule query. %v", err)

		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schedule models.BrewClassSchedule

		err := common.ScanRowsToModel(rows, &schedule)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "brew_classes_repository",
				"function": "GetBrewClassWithSchedule",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, nil, err
		}

		brewClassSchedules = append(brewClassSchedules, &schedule)
	}

	return &brewClassInfo, brewClassSchedules, nil
}

// UpdateUserBrewClassRegistration updates the user's brew class registration information in the user_brew_class_registrations table.
func (repo *BrewClassRepository) UpdateUserBrewClassRegistration(id string, registration models.UserBrewClassRegistration) error {
	// Replace "user_brew_class_registrations" with your actual table name.
	tableName := "user_brew_class_registrations"
	idColumn := "id"

	registration.RegistrationDate = time.Now()

	query, values, err := common.GenerateUpdateQuery(tableName, idColumn, id, registration)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_classes_repository",
			"function": "UpdateUserBrewClassRegistration",
		}).Errorf("Error generating update query: %v", err)
		return err
	}

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Values:", values)

	resp, err := repo.db.Pool.Exec(context.Background(), query, values...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_classes_repository",
			"function": "UpdateUserBrewClassRegistration",
		}).Errorf("Unable to update user's brew class registration. %v", err)
		return err
	}

	// Check the affected rows to determine the outcome
	affectedRows := resp.RowsAffected()
	if affectedRows == 0 {
		message := fmt.Sprintf("No brew class registration found with the ID: %s", id)
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_classes_repository",
			"function": "UpdateUserBrewClassRegistration",
		}).Debug(message)
		return fmt.Errorf(message)
	} else if affectedRows > 0 {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_classes_repository",
			"function": "UpdateUserBrewClassRegistration",
		}).Debug("User's brew class registration update successful")
	}

	return nil
}

// DeleteUserBrewClassRegistration deletes a user's brew class registration by ID from the user_brew_class_registrations table.
func (repo *BrewClassRepository) DeleteUserBrewClassRegistration(registrationID string) error {
	// Replace "user_brew_class_registrations" with your actual table name.
	query := "DELETE FROM user_brew_class_registrations WHERE id = $1"

	_, err := repo.db.Pool.Exec(context.Background(), query, registrationID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_classes_repository",
			"function": "DeleteUserBrewClassRegistration",
		}).Errorf("Unable to delete user's brew class registration. %v", err)

		return err
	}

	return nil
}

// CreateUserBrewClassRegistration inserts a new user's brew class registration into the user_brew_class_registrations table.
func (repo *BrewClassRepository) CreateUserBrewClassRegistration(registration models.UserBrewClassRegistration) error {
	// Replace "user_brew_class_registrations" with your actual table name.
	query := `
		INSERT INTO user_brew_class_registrations
		(user_id, class_id, status, registration_date, number_of_attendees, notes)
		VALUES
		($1, $2, $3, $4, $5, $6)
	`

	registration.RegistrationDate = time.Now()

	_, err := repo.db.Pool.Exec(context.Background(), query,
		registration.UserID, registration.ClassID, registration.Status, registration.RegistrationDate,
		registration.NumberOfAttendees, registration.Notes)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_classes_repository",
			"function": "CreateUserBrewClassRegistration",
		}).Errorf("Unable to create user's brew class registration. %v", err)

		return err
	}

	return nil
}

// GetUserBrewClassRegistration retrieves a user's brew class registration by user ID and class ID.
func (repo *BrewClassRepository) GetUserBrewClassRegistration(userID, classID string) (*models.UserBrewClassRegistration, error) {
	// Replace "user_brew_class_registrations" with your actual table name.
	query := `
		SELECT id, user_id, class_id, status, registration_date, number_of_attendees, notes
		FROM user_brew_class_registrations
		WHERE user_id = $1 AND class_id = $2
		LIMIT 1
	`

	var registration models.UserBrewClassRegistration

	err := repo.db.Pool.QueryRow(context.Background(), query, userID, classID).
		Scan(&registration.ID, &registration.UserID, &registration.ClassID, &registration.Status,
			&registration.RegistrationDate, &registration.NumberOfAttendees, &registration.Notes)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if no registration is found
		}

		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_classes_repository",
			"function": "GetUserBrewClassRegistration",
		}).Errorf("Unable to fetch user's brew class registration. %v", err)

		return nil, err
	}

	return &registration, nil
}
