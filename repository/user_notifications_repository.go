package repository

import (
	"context"

	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

// UserNotificationRepository handles database operations for user notifications.
type UserNotificationRepository struct {
	db *driver.DB
}

// NewUserNotificationRepository creates a new UserNotificationRepository.
func NewUserNotificationRepository(database *driver.DB) *UserNotificationRepository {
	return &UserNotificationRepository{db: database}
}

// CreateUserNotification creates a new user notification.
func (repo *UserNotificationRepository) CreateUserNotification(notification models.UserNotification) error {
	query := `
		INSERT INTO user_notifications (user_id, content, created_at, is_read, read_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := repo.db.Pool.Exec(context.Background(),
		query,
		notification.UserID, notification.Content, notification.CreatedAt,
		notification.IsRead, notification.ReadAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetUserNotificationByID fetches a user notification based on its ID.
func (repo *UserNotificationRepository) GetUserNotificationByID(notificationID string) (*models.UserNotification, error) {
	var notification models.UserNotification
	query := `
		SELECT *
		FROM user_notifications
		WHERE id = $1
	`
	err := repo.db.Pool.QueryRow(context.Background(), query, notificationID).Scan(
		&notification.ID, &notification.UserID, &notification.Content, &notification.CreatedAt,
		&notification.IsRead, &notification.ReadAt,
	)
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

// UpdateUserNotification updates an existing user notification.
func (repo *UserNotificationRepository) UpdateUserNotification(notification models.UserNotification) error {
	query := `
		UPDATE user_notifications
		SET content = $2, created_at = $3, is_read = $4, read_at = $5
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(),
		query,
		notification.ID, notification.Content, notification.CreatedAt,
		notification.IsRead, notification.ReadAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserNotification deletes a user notification based on its ID.
func (repo *UserNotificationRepository) DeleteUserNotification(notificationID string) error {
	query := `
		DELETE FROM user_notifications
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, notificationID)
	if err != nil {
		return err
	}
	return nil
}
