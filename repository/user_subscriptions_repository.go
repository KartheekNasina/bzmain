package repository

import (
	"context"

	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type UserSubscriptionRepository struct {
	db *driver.DB
}

func NewUserSubscriptionRepository(database *driver.DB) *UserSubscriptionRepository {
	return &UserSubscriptionRepository{
		db: database,
	}
}

// CreateUserSubscription creates a new user subscription.
func (repo *UserSubscriptionRepository) CreateUserSubscription(userSubscription models.UserSubscription) error {
	query := `
		INSERT INTO user_subscriptions (user_id, subscription_tier_id, start_date, end_date, is_active)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := repo.db.Pool.Exec(context.Background(),
		query,
		userSubscription.UserID,
		userSubscription.SubscriptionTierID,
		userSubscription.StartDate,
		userSubscription.EndDate,
		userSubscription.IsActive,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetUserSubscriptionByUserID fetches a user's subscription based on their user ID.
func (repo *UserSubscriptionRepository) GetUserSubscriptionByUserID(userID string) (*models.UserSubscription, error) {
	var userSubscription models.UserSubscription
	query := `
		SELECT *
		FROM user_subscriptions
		WHERE user_id = $1
	`
	err := repo.db.Pool.QueryRow(context.Background(), query, userID).Scan(
		&userSubscription.UserID, &userSubscription.SubscriptionTierID, &userSubscription.StartDate,
		&userSubscription.EndDate, &userSubscription.IsActive,
	)
	if err != nil {
		return nil, err
	}
	return &userSubscription, nil
}

// UpdateUserSubscription updates an existing user subscription.
func (repo *UserSubscriptionRepository) UpdateUserSubscription(userID string, userSubscription models.UserSubscription) error {
	query := `
		UPDATE user_subscriptions
		SET subscription_tier_id = $2, start_date = $3, end_date = $4, is_active = $5
		WHERE user_id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(),
		query,
		userID,
		userSubscription.SubscriptionTierID,
		userSubscription.StartDate,
		userSubscription.EndDate,
		userSubscription.IsActive,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserSubscription deletes a user subscription based on their user ID.
func (repo *UserSubscriptionRepository) DeleteUserSubscription(userID string) error {
	query := `
		DELETE FROM user_subscriptions
		WHERE user_id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, userID)
	if err != nil {
		return err
	}
	return nil
}
