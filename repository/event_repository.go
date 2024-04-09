// event_repository.go
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

type EventRepository struct {
	db *driver.DB
}

func NewEventRepository(database *driver.DB) *EventRepository {
	return &EventRepository{
		db: database,
	}
}

// GetEvents fetches a list of events with IsPublished true, limit, and offset.
func (repo *EventRepository) GetEvents(limit, offset int) ([]models.EventBasicInfo, error) {
	var events []models.EventBasicInfo

	// Replace "events" with your actual table name.
	query := `
		SELECT *
		FROM events
		WHERE is_published = true
		LIMIT $1 OFFSET $2
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "event_repository",
			"function": "GetEvents",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event models.EventBasicInfo

		err = rows.Scan(
			&event.ID, &event.Title, &event.ShortDescription, &event.BreweryID, &event.Thumbnail,
			&event.CreatedAt, &event.UpdatedAt, &event.IsPublished, &event.PublishedAt, &event.CreatedByID,
			&event.UpdatedByID, pq.Array(event.Images),
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "event_repository",
				"function": "GetEvents",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

// GetEventAndSchedule fetches an event and its schedule using a join query based on EventID.
func (repo *EventRepository) GetEventAndSchedule(eventID string) (*models.EventBasicInfo, *models.EventSchedule, error) {
	var event models.EventBasicInfo
	var schedule models.EventSchedule

	// Replace "events" and "event_schedules" with your actual table names.
	query := `
		SELECT e.*, s.*
		FROM events AS e
		INNER JOIN event_schedules AS s ON e.id = s.event_id
		WHERE e.id = $1 AND e.is_published = true
		LIMIT 1
	`

	row := repo.db.Pool.QueryRow(context.Background(), query, eventID)

	err := row.Scan(
		&event.ID, &event.Title, &event.ShortDescription, &event.BreweryID, &event.Thumbnail,
		&event.CreatedAt, &event.UpdatedAt, &event.IsPublished, &event.PublishedAt, &event.CreatedByID,
		&event.UpdatedByID, pq.Array(event.Images),
		&schedule.EventID, &schedule.StartDate, &schedule.EndDate, &schedule.StartTime, &schedule.EndTime,
		&schedule.Repeats, &schedule.MaxPeopleCount,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "event_repository",
			"function": "GetEventAndSchedule",
		}).Errorf("Unable to scan the row. %v", err)

		return nil, nil, err
	}

	return &event, &schedule, nil
}

// CreateEventJoinRequest inserts a new event join request into the event_join_requests table.
func (repo *EventRepository) CreateEventJoinRequest(request models.EventJoinRequest) error {
	// Replace "event_join_requests" with your actual table name.
	query := `
		INSERT INTO event_join_requests
		(event_id, user_id, status, requested_at, notes)
		VALUES
		($1, $2, $3, $4, $5)
	`

	request.RequestedAt = time.Now()

	_, err := repo.db.Pool.Exec(context.Background(), query,
		request.EventID, request.UserID, request.Status, request.RequestedAt, request.Notes)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "event_repository",
			"function": "CreateEventJoinRequest",
		}).Errorf("Unable to create event join request. %v", err)

		return err
	}

	return nil
}

// UpdateEventJoinRequest updates event join request information in the event_join_requests table.
func (repo *EventRepository) UpdateEventJoinRequest(id string, request models.EventJoinRequest) error {
	// Replace "event_join_requests" with your actual table name.
	tableName := "event_join_requests"
	idColumn := "id"

	request.UpdatedAt = time.Now()

	query, values, err := common.GenerateUpdateQuery(tableName, idColumn, id, request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "event_repository",
			"function": "UpdateEventJoinRequest",
		}).Errorf("Error generating update query: %v", err)
		return err
	}

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Values:", values)

	resp, err := repo.db.Pool.Exec(context.Background(), query, values...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "event_repository",
			"function": "UpdateEventJoinRequest",
		}).Errorf("Unable to update event join request. %v", err)
		return err
	}

	// Check the affected rows to determine the outcome
	affectedRows := resp.RowsAffected()
	if affectedRows == 0 {
		message := fmt.Sprintf("No event join request found with the ID: %s", id)
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "event_repository",
			"function": "UpdateEventJoinRequest",
		}).Debug(message)
		return fmt.Errorf(message)
	} else if affectedRows > 0 {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "event_repository",
			"function": "UpdateEventJoinRequest",
		}).Debug("Event join request update successful")
	}

	return nil
}

// DeleteEventJoinRequest deletes event join request by ID from the event_join_requests table.
func (repo *EventRepository) DeleteEventJoinRequest(requestID string) error {
	// Replace "event_join_requests" with your actual table name.
	query := "DELETE FROM event_join_requests WHERE id = $1"

	_, err := repo.db.Pool.Exec(context.Background(), query, requestID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "event_repository",
			"function": "DeleteEventJoinRequest",
		}).Errorf("Unable to delete event join request. %v", err)

		return err
	}

	return nil
}
