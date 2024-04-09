// meetup_repository.go
package repository

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type MeetupRepository struct {
	db *driver.DB
}

func NewMeetupRepository(database *driver.DB) *MeetupRepository {
	return &MeetupRepository{
		db: database,
	}
}

// GetMeetups fetches a list of meetups with limit and offset.
func (repo *MeetupRepository) GetMeetups(limit, offset int) ([]models.Meetup, error) {
	var meetups []models.Meetup

	// Replace "meetups" with your actual table name.
	query := `
		SELECT *
		FROM meetups
		LIMIT $1 OFFSET $2
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "meetup_repository",
			"function": "GetMeetups",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var meetup models.Meetup

		err = rows.Scan(
			&meetup.ID, &meetup.BreweryID, &meetup.OrganizerID, &meetup.Title, &meetup.Description,
			&meetup.MeetupDate, &meetup.StartTime, &meetup.EndTime, &meetup.MaxAttendees, &meetup.CreatedAt,
			&meetup.UpdatedAt, &meetup.Status, &meetup.ApprovedAt, &meetup.RejectedAt, &meetup.CancellationReason,
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "meetup_repository",
				"function": "GetMeetups",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		meetups = append(meetups, meetup)
	}

	return meetups, nil
}

// GetMeetup fetches a meetup based on its ID.
func (repo *MeetupRepository) GetMeetup(id string) (*models.Meetup, error) {
	var meetup models.Meetup

	// Replace "meetups" with your actual table name.
	query := `
		SELECT *
		FROM meetups
		WHERE id = $1
		LIMIT 1
	`

	row := repo.db.Pool.QueryRow(context.Background(), query, id)

	err := row.Scan(
		&meetup.ID, &meetup.BreweryID, &meetup.OrganizerID, &meetup.Title, &meetup.Description,
		&meetup.MeetupDate, &meetup.StartTime, &meetup.EndTime, &meetup.MaxAttendees, &meetup.CreatedAt,
		&meetup.UpdatedAt, &meetup.Status, &meetup.ApprovedAt, &meetup.RejectedAt, &meetup.CancellationReason,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "meetup_repository",
			"function": "GetMeetup",
		}).Errorf("Unable to scan the row. %v", err)

		return nil, err
	}

	return &meetup, nil
}

// CreateMeetupAttendee inserts a new meetup attendee into the meetup_attendees table.
func (repo *MeetupRepository) CreateMeetupAttendee(attendee models.MeetupAttendee) error {
	// Replace "meetup_attendees" with your actual table name.
	query := `
		INSERT INTO meetup_attendees
		(meetup_id, user_id, rsvp_date, attended)
		VALUES
		($1, $2, $3, $4)
	`

	_, err := repo.db.Pool.Exec(context.Background(), query,
		attendee.MeetupID, attendee.UserID, attendee.RSVPDate, attendee.Attended)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "meetup_repository",
			"function": "CreateMeetupAttendee",
		}).Errorf("Unable to create meetup attendee. %v", err)

		return err
	}

	return nil
}

// UpdateMeetupAttendee updates meetup attendee information in the meetup_attendees table.
func (repo *MeetupRepository) UpdateMeetupAttendee(id string, attendee models.MeetupAttendee) error {
	// Replace "meetup_attendees" with your actual table name.
	tableName := "meetup_attendees"
	idColumn := "meetup_id"

	query, values, err := common.GenerateUpdateQuery(tableName, idColumn, id, attendee)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "meetup_repository",
			"function": "UpdateMeetupAttendee",
		}).Errorf("Error generating update query: %v", err)
		return err
	}

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Values:", values)

	resp, err := repo.db.Pool.Exec(context.Background(), query, values...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "meetup_repository",
			"function": "UpdateMeetupAttendee",
		}).Errorf("Unable to update meetup attendee. %v", err)
		return err
	}

	// Check the affected rows to determine the outcome
	affectedRows := resp.RowsAffected()
	if affectedRows == 0 {
		message := fmt.Sprintf("No meetup attendee found with the ID: %s", id)
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "meetup_repository",
			"function": "UpdateMeetupAttendee",
		}).Debug(message)
		return fmt.Errorf(message)
	} else if affectedRows > 0 {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "meetup_repository",
			"function": "UpdateMeetupAttendee",
		}).Debug("Meetup attendee update successful")
	}

	return nil
}

// DeleteMeetupAttendee deletes meetup attendee by meetup ID from the meetup_attendees table.
func (repo *MeetupRepository) DeleteMeetupAttendee(meetupID, userID string) error {
	// Replace "meetup_attendees" with your actual table name.
	query := "DELETE FROM meetup_attendees WHERE meetup_id = $1 AND user_id = $2"

	_, err := repo.db.Pool.Exec(context.Background(), query, meetupID, userID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "meetup_repository",
			"function": "DeleteMeetupAttendee",
		}).Errorf("Unable to delete meetup attendee. %v", err)

		return err
	}

	return nil
}
