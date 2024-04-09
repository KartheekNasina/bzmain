// community_repository.go
package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type CommunityRepository struct {
	db *driver.DB
}

func NewCommunityRepository(database *driver.DB) *CommunityRepository {
	return &CommunityRepository{
		db: database,
	}
}

// GetCommunityCategories fetches a list of community categories.
func (repo *CommunityRepository) GetCommunityCategories() ([]models.CommunityCategory, error) {
	var categories []models.CommunityCategory

	// Replace "community_categories" with your actual table name.
	query := `
		SELECT *
		FROM community_categories
	`

	rows, err := repo.db.Pool.Query(context.Background(), query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "community_repository",
			"function": "GetCommunityCategories",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category models.CommunityCategory

		err = rows.Scan(
			&category.ID, &category.Type, &category.Thumbnail, &category.Title,
			&category.CreatedAt, &category.UpdatedAt, &category.CreatedByID, &category.UpdatedByID,
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "community_repository",
				"function": "GetCommunityCategories",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

// GetCommunityCenters fetches a list of community centers with IsPublished true, limit, and offset.
func (repo *CommunityRepository) GetCommunityCenters(limit, offset int) ([]models.CommunityCenter, error) {
	var centers []models.CommunityCenter

	// Replace "community_centers" with your actual table name.
	query := `
		SELECT *
		FROM community_centers
		WHERE is_published = true
		LIMIT $1 OFFSET $2
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "community_repository",
			"function": "GetCommunityCenters",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var center models.CommunityCenter

		err := common.ScanRowsToModel(rows, &center)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "community_repository",
				"function": "GetCommunityCenters",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		centers = append(centers, center)
	}

	return centers, nil
}

// GetCommunityDrives fetches a list of community drives with IsPublished true, limit, and offset.
func (repo *CommunityRepository) GetCommunityDrives(limit, offset int) ([]models.CommunityDrive, error) {
	var drives []models.CommunityDrive

	// Replace "community_drives" with your actual table name.
	query := `
		SELECT *
		FROM community_drives
		WHERE is_published = true
		LIMIT $1 OFFSET $2
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "community_repository",
			"function": "GetCommunityDrives",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var drive models.CommunityDrive

		err = rows.Scan(
			&drive.ID, &drive.Title, &drive.CommunityType, &drive.Description, &drive.Address,
			&drive.Lat, &drive.Lng, &drive.ContactName, &drive.ContactPhoneNumber, &drive.StartTime,
			&drive.EndTime, &drive.IsPublished, &drive.CreatedAt, &drive.UpdatedAt, &drive.PublishedAt,
			&drive.Thumbnail, pq.Array(drive.Images),
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "community_repository",
				"function": "GetCommunityDrives",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		drives = append(drives, drive)
	}

	return drives, nil
}

// UpdateUserCommunitiesRegistration updates user community drive registration information in the user_community_drive_registrations table.
func (repo *CommunityRepository) UpdateUserCommunitiesRegistration(id string, registration models.UserCommunityDriveRegistration) error {
	// Replace "user_community_drive_registrations" with your actual table name.
	tableName := "user_community_drive_registrations"
	idColumn := "id"

	registration.UpdatedAt = time.Now()

	query, values, err := common.GenerateUpdateQuery(tableName, idColumn, id, registration)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "community_repository",
			"function": "UpdateUserCommunitiesRegistration",
		}).Errorf("Error generating update query: %v", err)
		return err
	}

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Values:", values)

	resp, err := repo.db.Pool.Exec(context.Background(), query, values...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "community_repository",
			"function": "UpdateUserCommunitiesRegistration",
		}).Errorf("Unable to update user community drive registration. %v", err)
		return err
	}

	// Check the affected rows to determine the outcome
	affectedRows := resp.RowsAffected()
	if affectedRows == 0 {
		message := fmt.Sprintf("No user community drive registration found with the ID: %s", id)
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "community_repository",
			"function": "UpdateUserCommunitiesRegistration",
		}).Debug(message)
		return fmt.Errorf(message)
	} else if affectedRows > 0 {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "community_repository",
			"function": "UpdateUserCommunitiesRegistration",
		}).Debug("User community drive registration update successful")
	}

	return nil
}

// DeleteUserCommunitiesRegistration deletes user community drive registration by ID from the user_community_drive_registrations table.
func (repo *CommunityRepository) DeleteUserCommunitiesRegistration(registrationID string) error {
	// Replace "user_community_drive_registrations" with your actual table name.
	query := "DELETE FROM user_community_drive_registrations WHERE id = $1"

	_, err := repo.db.Pool.Exec(context.Background(), query, registrationID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "community_repository",
			"function": "DeleteUserCommunitiesRegistration",
		}).Errorf("Unable to delete user community drive registration. %v", err)

		return err
	}

	return nil
}

// CreateUserCommunitiesRegistration inserts a new user community drive registration into the user_community_drive_registrations table.
func (repo *CommunityRepository) CreateUserCommunitiesRegistration(registration models.UserCommunityDriveRegistration) error {
	// Replace "user_community_drive_registrations" with your actual table name.
	query := `
		INSERT INTO user_community_drive_registrations
		(user_id, drive_id, status, registration_date, notes)
		VALUES
		($1, $2, $3, $4, $5)
	`

	registration.RegistrationDate = time.Now()

	_, err := repo.db.Pool.Exec(context.Background(), query,
		registration.UserID, registration.DriveID, registration.Status, registration.RegistrationDate, registration.Notes)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "community_repository",
			"function": "CreateUserCommunitiesRegistration",
		}).Errorf("Unable to create user community drive registration. %v", err)

		return err
	}

	return nil
}
