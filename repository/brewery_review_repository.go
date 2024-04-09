// brewery_review_repository.go
package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type BreweryReviewRepository struct {
	db *driver.DB
}

func NewBreweryReviewRepository(database *driver.DB) *BreweryReviewRepository {
	return &BreweryReviewRepository{
		db: database,
	}
}

// CreateBreweryReview inserts a new brewery review into the brewery_reviews table.
func (repo *BreweryReviewRepository) CreateBreweryReview(review models.BreweryReview) error {
	// Replace "brewery_reviews" with your actual table name.
	query := `
		INSERT INTO brewery_reviews
		(brewery_id, rating, user_id, title, description, created_at, updated_at)
		VALUES
		($1, $2, $3, $4, $5, $6, $7)
	`

	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()

	_, err := repo.db.Pool.Exec(context.Background(), query,
		review.BreweryID, review.Rating, review.UserID, review.Title, review.Description, review.CreatedAt, review.UpdatedAt)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_review_repository",
			"function": "CreateBreweryReview",
		}).Errorf("Unable to create brewery review. %v", err)

		return err
	}

	return nil
}

// GetBreweryReviewByID fetches a brewery review by its ID from the brewery_reviews table.
func (repo *BreweryReviewRepository) GetBreweryReviewByID(id string) (*models.BreweryReview, error) {
	// Replace "brewery_reviews" with your actual table name.
	query := `
		SELECT *
		FROM brewery_reviews
		WHERE id = $1
	`

	var review models.BreweryReview

	err := repo.db.Pool.QueryRow(context.Background(), query, id).Scan(
		&review.ID, &review.BreweryID, &review.Rating, &review.UserID, &review.Title, &review.Description, &review.CreatedAt, &review.UpdatedAt,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_review_repository",
			"function": "GetBreweryReviewByID",
		}).Errorf("Unable to execute the query. %v", err)
		return nil, err
	}

	return &review, nil
}

// UpdateBreweryReview updates the brewery review information in the brewery_reviews table.
func (repo *BreweryReviewRepository) UpdateBreweryReview(id string, review models.BreweryReview) error {
	// Replace "brewery_reviews" with your actual table name.
	tableName := "brewery_reviews"
	idColumn := "id"

	review.UpdatedAt = time.Now()

	query, values, err := common.GenerateUpdateQuery(tableName, idColumn, id, review)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_review_repository",
			"function": "UpdateBreweryReview",
		}).Errorf("Error generating update query: %v", err)
		return err
	}

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Values:", values)

	resp, err := repo.db.Pool.Exec(context.Background(), query, values...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_review_repository",
			"function": "UpdateBreweryReview",
		}).Errorf("Unable to update brewery review. %v", err)
		return err
	}

	// Check the affected rows to determine the outcome
	affectedRows := resp.RowsAffected()
	if affectedRows == 0 {
		message := fmt.Sprintf("No brewery review found with the ID: %s", id)
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_review_repository",
			"function": "UpdateBreweryReview",
		}).Debug(message)
		return fmt.Errorf(message)
	} else if affectedRows > 0 {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_review_repository",
			"function": "UpdateBreweryReview",
		}).Debug("Brewery review update successful")
	}

	return nil
}

// DeleteBreweryReview deletes a brewery review by ID and BreweryID from the brewery_reviews table.
func (repo *BreweryReviewRepository) DeleteBreweryReview(reviewID, breweryID string) error {
	// Replace "brewery_reviews" with your actual table name.
	query := "DELETE FROM brewery_reviews WHERE id = $1 AND brewery_id = $2"

	_, err := repo.db.Pool.Exec(context.Background(), query, reviewID, breweryID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_review_repository",
			"function": "DeleteBreweryReview",
		}).Errorf("Unable to delete brewery review. %v", err)

		return err
	}

	return nil
}

// GetBreweryReviews fetches a list of brewery reviews with pagination and optional filtering by BreweryID.
func (repo *BreweryReviewRepository) GetBreweryReviews(limit, offset int, breweryID string) ([]models.BreweryReview, error) {
	var reviews []models.BreweryReview

	// Replace "brewery_reviews" with your actual table name.
	query := `
		SELECT *
		FROM brewery_reviews
		WHERE brewery_id = $1
		LIMIT $2 OFFSET $3
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, breweryID, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_review_repository",
			"function": "GetBreweryReviews",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var review models.BreweryReview

		err = rows.Scan(
			&review.ID, &review.BreweryID, &review.Rating, &review.UserID, &review.Title,
			&review.Description, &review.CreatedAt, &review.UpdatedAt,
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "brewery_review_repository",
				"function": "GetBreweryReviews",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}
