// brew_tours_repository.go
package repository

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type BrewTourRepository struct {
	db *driver.DB
}

func NewBrewTourRepository(database *driver.DB) *BrewTourRepository {
	return &BrewTourRepository{
		db: database,
	}
}

func (repo *BrewTourRepository) GetTours(limit, offset int) ([]models.BrewTour, error) {
	var items []models.BrewTour

	query := `
		SELECT
		    b.id AS tour_id,
		    b.title,
		    b.short_description,
		    b.brewery_id,
		    b.thumbnail,
		    b.is_published,
		    b.published_at,
		    b.images,
		    b.created_at,
		    b.updated_at,
		    b.created_by_id,
		    b.updated_by_id,
		    s.start_date,
		    s.end_date,
		    s.start_time,
		    s.end_time,
		    s.repeats,
		    s.max_people_count
		FROM
		    brew_tours_basic_info b
		JOIN
		    brew_tours_schedule s
		ON
		    b.id = s.tour_id
		LIMIT $1 OFFSET $2
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "tour_repository",
			"function": "GetTours",
		}).Errorf("Unable to execute the query. %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var i models.BrewTour
		if err := rows.Scan(
			&i.TourID,
			&i.Title,
			&i.ShortDescription,
			&i.BreweryID,
			&i.Thumbnail,
			&i.IsPublished,
			&i.PublishedAt,
			&i.Images,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedByID,
			&i.UpdatedByID,
			&i.StartDate,
			&i.EndDate,
			&i.StartTime,
			&i.EndTime,
			&i.Repeats,
			&i.MaxPeopleCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// UpdateUserBrewTourRegistration updates the user's brew tour registration information in the user_brew_tour_registrations table.
func (repo *BrewTourRepository) UpdateUserBrewTourRegistration(id string, registration models.UserBrewTourRegistration) error {
	// Replace "user_brew_tour_registrations" with your actual table name.
	tableName := "user_brew_tour_registrations"
	idColumn := "id"

	query, values, err := common.GenerateUpdateQuery(tableName, idColumn, id, registration)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_tours_repository",
			"function": "UpdateUserBrewTourRegistration",
		}).Errorf("Error generating update query: %v", err)
		return err
	}

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Values:", values)

	resp, err := repo.db.Pool.Exec(context.Background(), query, values...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_tours_repository",
			"function": "UpdateUserBrewTourRegistration",
		}).Errorf("Unable to update user's brew tour registration. %v", err)
		return err
	}

	// Check the affected rows to determine the outcome
	affectedRows := resp.RowsAffected()
	if affectedRows == 0 {
		message := fmt.Sprintf("No brew tour registration found with the ID: %s", id)
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_tours_repository",
			"function": "UpdateUserBrewTourRegistration",
		}).Debug(message)
		return fmt.Errorf(message)
	} else if affectedRows > 0 {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_tours_repository",
			"function": "UpdateUserBrewTourRegistration",
		}).Debug("User's brew tour registration update successful")
	}

	return nil
}

// DeleteUserBrewTourRegistration deletes a user's brew tour registration by ID from the user_brew_tour_registrations table.
func (repo *BrewTourRepository) DeleteUserBrewTourRegistration(registrationID string) error {
	// Replace "user_brew_tour_registrations" with your actual table name.
	query := "DELETE FROM user_brew_tour_registrations WHERE id = $1"

	_, err := repo.db.Pool.Exec(context.Background(), query, registrationID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_tours_repository",
			"function": "DeleteUserBrewTourRegistration",
		}).Errorf("Unable to delete user's brew tour registration. %v", err)

		return err
	}

	return nil
}

// CreateUserBrewTourRegistration inserts a new user's brew tour registration into the user_brew_tour_registrations table.
func (repo *BrewTourRepository) CreateUserBrewTourRegistration(registration models.UserBrewTourRegistration) error {
	// Replace "user_brew_tour_registrations" with your actual table name.
	query := `
		INSERT INTO user_brew_tour_registrations
		(user_id, tour_id, status, registration_date, number_of_people, notes)
		VALUES
		($1, $2, $3, $4, $5, $6)
	`

	_, err := repo.db.Pool.Exec(context.Background(), query,
		registration.UserID, registration.TourID, registration.Status, registration.RegistrationDate,
		registration.NumberOfPeople, registration.Notes)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brew_tours_repository",
			"function": "CreateUserBrewTourRegistration",
		}).Errorf("Unable to create user's brew tour registration. %v", err)

		return err
	}

	return nil
}
