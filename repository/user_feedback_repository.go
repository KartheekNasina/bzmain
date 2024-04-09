package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

// FeedbackType represents the type of user feedback.
type FeedbackType string

// UserFeedbackRepository handles database operations for UserFeedback.
type UserFeedbackRepository struct {
	db *driver.DB
}

// NewUserFeedbackRepository creates a new UserFeedbackRepository.
func NewUserFeedbackRepository(database *driver.DB) *UserFeedbackRepository {
	return &UserFeedbackRepository{
		db: database,
	}
}

// CreateUserFeedback inserts a new user feedback into the database.
func (repo *UserFeedbackRepository) CreateUserFeedback(feedback models.UserFeedback) error {
	query := `
		INSERT INTO user_feedback (user_id, feedback_date, type, subject, description, is_resolved, resolved_at, resolution_notes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := repo.db.Pool.Exec(context.Background(), query,
		feedback.UserID, feedback.FeedbackDate, feedback.Type, feedback.Subject,
		feedback.Description, feedback.IsResolved, feedback.ResolvedAt, feedback.ResolutionNotes)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_feedback_repository",
			"function": "CreateUserFeedback",
		}).Errorf("Unable to create user feedback. %v", err)

		return err
	}

	return nil
}

// GetUserFeedbackByID retrieves a user feedback by its ID.
func (repo *UserFeedbackRepository) GetUserFeedbackByID(feedbackID string) (*models.UserFeedback, error) {
	var feedback models.UserFeedback

	query := "SELECT id, user_id, feedback_date, type, subject, description, is_resolved, resolved_at, resolution_notes FROM user_feedback WHERE id = $1"
	err := repo.db.Pool.QueryRow(context.Background(), query, feedbackID).Scan(
		&feedback.ID, &feedback.UserID, &feedback.FeedbackDate, &feedback.Type, &feedback.Subject,
		&feedback.Description, &feedback.IsResolved, &feedback.ResolvedAt, &feedback.ResolutionNotes)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_feedback_repository",
			"function": "GetUserFeedbackByID",
		}).Errorf("Unable to retrieve user feedback. %v", err)

		return nil, err
	}

	return &feedback, nil
}

// UpdateUserFeedback updates a user feedback in the database.
func (repo *UserFeedbackRepository) UpdateUserFeedback(feedbackId string, feedback models.UserFeedback) error {
	query := `
		UPDATE user_feedback
		SET user_id = $2, feedback_date = $3, type = $4, subject = $5, description = $6,
			is_resolved = $7, resolved_at = $8, resolution_notes = $9
		WHERE id = $1
	`

	_, err := repo.db.Pool.Exec(context.Background(), query,
		feedbackId, feedback.UserID, feedback.FeedbackDate, feedback.Type, feedback.Subject,
		feedback.Description, feedback.IsResolved, feedback.ResolvedAt, feedback.ResolutionNotes)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_feedback_repository",
			"function": "UpdateUserFeedback",
		}).Errorf("Unable to update user feedback. %v", err)

		return err
	}

	return nil
}

// DeleteUserFeedback deletes a user feedback from the database by its ID.
func (repo *UserFeedbackRepository) DeleteUserFeedback(feedbackID string) error {
	query := "DELETE FROM user_feedback WHERE id = $1"
	_, err := repo.db.Pool.Exec(context.Background(), query, feedbackID)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_feedback_repository",
			"function": "DeleteUserFeedback",
		}).Errorf("Unable to delete user feedback. %v", err)

		return err
	}

	return nil
}
