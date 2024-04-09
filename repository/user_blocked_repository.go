package repository

import (
	"context"

	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type UserBlockedListRepository struct {
	db *driver.DB
}

func NewUserBlockedListRepository(database *driver.DB) *UserBlockedListRepository {
	return &UserBlockedListRepository{
		db: database,
	}
}

// CreateUserBlocked creates a new entry in the user blocked list.
func (repo *UserBlockedListRepository) CreateUserBlocked(userBlocked models.UserBlockedList) error {
	query := `
		INSERT INTO user_blocked_list (user_id, blocked_user_id, blocked_at)
		VALUES ($1, $2, $3)
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, userBlocked.UserId, userBlocked.BlockedUserId, userBlocked.BlockedAt)
	if err != nil {
		return err
	}
	return nil
}

// GetUserBlockedByUserID fetches a list of users blocked by a specific user.
func (repo *UserBlockedListRepository) GetUserBlockedByUserID(userId string) ([]models.UserBlockedList, error) {
	var userBlockedEntries []models.UserBlockedList
	query := `
		SELECT user_id, blocked_user_id, blocked_at
		FROM user_blocked_list
		WHERE user_id = $1
	`
	rows, err := repo.db.Pool.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userBlocked models.UserBlockedList
		err := rows.Scan(&userBlocked.UserId, &userBlocked.BlockedUserId, &userBlocked.BlockedAt)
		if err != nil {
			return nil, err
		}
		userBlockedEntries = append(userBlockedEntries, userBlocked)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return userBlockedEntries, nil
}

// DeleteUserBlocked deletes a user from the blocked list.
func (repo *UserBlockedListRepository) DeleteUserBlocked(userId, blockedUserId string) error {
	query := `
		DELETE FROM user_blocked_list
		WHERE user_id = $1 AND blocked_user_id = $2
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, userId, blockedUserId)
	if err != nil {
		return err
	}
	return nil
}
