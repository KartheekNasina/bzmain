// user_referral_repository.go
package repository

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type UserReferralRepository struct {
	db *driver.DB
}

func NewUserReferralRepository(database *driver.DB) *UserReferralRepository {
	return &UserReferralRepository{
		db: database,
	}
}

// GetUserReferral fetches a list of user referrals with limit and offset.
func (repo *UserReferralRepository) GetUserReferral(limit, offset int) ([]models.UserReferral, error) {
	var userReferrals []models.UserReferral

	// Replace "user_referrals" with your actual table name.
	query := `
		SELECT *
		FROM user_referrals
		LIMIT $1 OFFSET $2
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_referral_repository",
			"function": "GetUserReferral",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var referral models.UserReferral

		err = rows.Scan(
			&referral.ID, &referral.ReferrerID, &referral.RefereeID, &referral.ReferralCode,
			&referral.Status, &referral.ReferredAt, &referral.CompletedAt, &referral.RewardClaimed,
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "user_referral_repository",
				"function": "GetUserReferral",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		userReferrals = append(userReferrals, referral)
	}

	return userReferrals, nil
}

// GetUserReferralByID fetches a user referral based on its ID.
func (repo *UserReferralRepository) GetUserReferralByID(id string) (*models.UserReferral, error) {
	var referral models.UserReferral

	// Replace "user_referrals" with your actual table name.
	query := `
		SELECT *
		FROM user_referrals
		WHERE id = $1
		LIMIT 1
	`

	row := repo.db.Pool.QueryRow(context.Background(), query, id)

	err := row.Scan(
		&referral.ID, &referral.ReferrerID, &referral.RefereeID, &referral.ReferralCode,
		&referral.Status, &referral.ReferredAt, &referral.CompletedAt, &referral.RewardClaimed,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_referral_repository",
			"function": "GetUserReferralByID",
		}).Errorf("Unable to scan the row. %v", err)

		return nil, err
	}

	return &referral, nil
}

// CreateUserReferral inserts a new user referral into the user_referrals table.
func (repo *UserReferralRepository) CreateUserReferral(referral models.UserReferral) error {
	// Replace "user_referrals" with your actual table name.
	query := `
		INSERT INTO user_referrals
		(referrer_id, referee_id, referral_code, status, referred_at, completed_at, reward_claimed)
		VALUES
		($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := repo.db.Pool.Exec(context.Background(), query,
		referral.ReferrerID, referral.RefereeID, referral.ReferralCode, referral.Status,
		referral.ReferredAt, referral.CompletedAt, referral.RewardClaimed)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_referral_repository",
			"function": "CreateUserReferral",
		}).Errorf("Unable to create user referral. %v", err)

		return err
	}

	return nil
}

// UpdateUserReferral updates user referral information in the user_referrals table.
func (repo *UserReferralRepository) UpdateUserReferral(id string, referral models.UserReferral) error {
	// Replace "user_referrals" with your actual table name.
	tableName := "user_referrals"
	idColumn := "id"

	query, values, err := common.GenerateUpdateQuery(tableName, idColumn, id, referral)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_referral_repository",
			"function": "UpdateUserReferral",
		}).Errorf("Error generating update query: %v", err)
		return err
	}

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Values:", values)

	resp, err := repo.db.Pool.Exec(context.Background(), query, values...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_referral_repository",
			"function": "UpdateUserReferral",
		}).Errorf("Unable to update user referral. %v", err)
		return err
	}

	// Check the affected rows to determine the outcome
	affectedRows := resp.RowsAffected()
	if affectedRows == 0 {
		message := fmt.Sprintf("No user referral found with the ID: %s", id)
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_referral_repository",
			"function": "UpdateUserReferral",
		}).Debug(message)
		return fmt.Errorf(message)
	} else if affectedRows > 0 {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_referral_repository",
			"function": "UpdateUserReferral",
			"id":       id,
		}).Debug("User referral update successful")
	}

	return nil
}

// DeleteUserReferral deletes user referral by ID from the user_referrals table.
func (repo *UserReferralRepository) DeleteUserReferral(id string) error {
	// Replace "user_referrals" with your actual table name.
	query := "DELETE FROM user_referrals WHERE id = $1"

	_, err := repo.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_referral_repository",
			"function": "DeleteUserReferral",
		}).Errorf("Unable to delete user referral. %v", err)

		return err
	}

	return nil
}
