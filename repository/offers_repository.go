package repository

import (
	"context"

	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type OffersRepository struct {
	db *driver.DB
}

func NewOffersRepository(database *driver.DB) *OffersRepository {
	return &OffersRepository{
		db: database,
	}
}

// CreateBreweryOffer creates a new brewery offer.
func (repo *OffersRepository) CreateBreweryOffer(breweryOffer models.BreweryOffer) error {
	query := `
		INSERT INTO brewery_offers (brewery_id, offer_type, title, description, thumbnail, images, start_date, end_date, start_time, end_time, repeats, status, created_by_id, updated_by_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`
	_, err := repo.db.Pool.Exec(context.Background(),
		query,
		breweryOffer.BreweryID,
		breweryOffer.OfferType,
		breweryOffer.Title,
		breweryOffer.Description,
		breweryOffer.Thumbnail,
		breweryOffer.Images,
		breweryOffer.StartDate,
		breweryOffer.EndDate,
		breweryOffer.StartTime,
		breweryOffer.EndTime,
		breweryOffer.Repeats,
		breweryOffer.Status,
		breweryOffer.CreatedByID,
		breweryOffer.UpdatedByID,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetBreweryOffers fetches brewery offers with pagination.
func (repo *OffersRepository) GetBreweryOffers(limit, offset int) ([]models.BreweryOffer, error) {
	var breweryOffers []models.BreweryOffer
	query := `
		SELECT *
		FROM brewery_offers
		LIMIT $1 OFFSET $2
	`
	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var breweryOffer models.BreweryOffer
		var imagesArray []string // Use pq.StringArray for the "images" column

		err = rows.Scan(
			&breweryOffer.ID, &breweryOffer.BreweryID, &breweryOffer.OfferType,
			&breweryOffer.Title, &breweryOffer.Description, &breweryOffer.Thumbnail, &imagesArray,
			&breweryOffer.StartDate, &breweryOffer.EndDate, &breweryOffer.StartTime, &breweryOffer.EndTime,
			&breweryOffer.Repeats, &breweryOffer.Status, &breweryOffer.CreatedAt, &breweryOffer.UpdatedAt,
			&breweryOffer.CreatedByID, &breweryOffer.UpdatedByID,
		)
		if err != nil {
			return nil, err
		}

		// Convert pq.StringArray to a slice of strings
		breweryOffer.Images = imagesArray

		breweryOffers = append(breweryOffers, breweryOffer)
	}
	return breweryOffers, nil
}

// GetBreweryOfferByID fetches a brewery offer based on its ID.
func (repo *OffersRepository) GetBreweryOfferByID(offerID string) (*models.BreweryOffer, error) {
	var breweryOffer models.BreweryOffer
	query := `
		SELECT *
		FROM brewery_offers
		WHERE id = $1
	`
	err := repo.db.Pool.QueryRow(context.Background(), query, offerID).Scan(
		&breweryOffer.ID, &breweryOffer.BreweryID, &breweryOffer.OfferType,
		&breweryOffer.Title, &breweryOffer.Description, &breweryOffer.Thumbnail, &breweryOffer.Images,
		&breweryOffer.StartDate, &breweryOffer.EndDate, &breweryOffer.StartTime, &breweryOffer.EndTime,
		&breweryOffer.Repeats, &breweryOffer.Status, &breweryOffer.CreatedAt, &breweryOffer.UpdatedAt,
		&breweryOffer.CreatedByID, &breweryOffer.UpdatedByID,
	)
	if err != nil {
		return nil, err
	}
	return &breweryOffer, nil
}

// UpdateBreweryOffer updates an existing brewery offer.
func (repo *OffersRepository) UpdateBreweryOffer(breweryOffer models.BreweryOffer) error {
	query := `
		UPDATE brewery_offers
		SET brewery_id = $2, offer_type = $3, title = $4, description = $5, thumbnail = $6, images = $7,
			start_date = $8, end_date = $9, start_time = $10, end_time = $11, repeats = $12, status = $13, updated_by_id = $14
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(),
		query,
		breweryOffer.ID, breweryOffer.BreweryID, breweryOffer.OfferType,
		breweryOffer.Title, breweryOffer.Description, breweryOffer.Thumbnail, breweryOffer.Images,
		breweryOffer.StartDate, breweryOffer.EndDate, breweryOffer.StartTime, breweryOffer.EndTime,
		breweryOffer.Repeats, breweryOffer.Status, breweryOffer.UpdatedByID,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBreweryOffer deletes a brewery offer based on its ID.
func (repo *OffersRepository) DeleteBreweryOffer(offerID string) error {
	query := `
		DELETE FROM brewery_offers
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, offerID)
	if err != nil {
		return err
	}
	return nil
}

// CreateUserExclusiveOffer creates a new user exclusive offer.
func (repo *OffersRepository) CreateUserExclusiveOffer(userExclusiveOffer models.UserExclusiveOffer) error {
	query := `
		INSERT INTO user_exclusive_offers (user_id, brewery_id, title, description, thumbnail, start_date, end_date, start_time, end_time, repeats, status, created_by_id, updated_by_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	_, err := repo.db.Pool.Exec(context.Background(),
		query,
		userExclusiveOffer.UserID, userExclusiveOffer.BreweryID, userExclusiveOffer.Title,
		userExclusiveOffer.Description, userExclusiveOffer.Thumbnail, userExclusiveOffer.StartDate,
		userExclusiveOffer.EndDate, userExclusiveOffer.StartTime, userExclusiveOffer.EndTime,
		userExclusiveOffer.Repeats, userExclusiveOffer.Status, userExclusiveOffer.CreatedByID,
		userExclusiveOffer.UpdatedByID,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetUserExclusiveOffers fetches user exclusive offers with pagination.
func (repo *OffersRepository) GetUserExclusiveOffers(limit, offset int) ([]models.UserExclusiveOffer, error) {
	var userExclusiveOffers []models.UserExclusiveOffer
	query := `
		SELECT *
		FROM user_exclusive_offers
		LIMIT $1 OFFSET $2
	`
	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userExclusiveOffer models.UserExclusiveOffer

		err = rows.Scan(
			&userExclusiveOffer.ID, &userExclusiveOffer.UserID, &userExclusiveOffer.BreweryID, &userExclusiveOffer.Title,
			&userExclusiveOffer.Description, &userExclusiveOffer.Thumbnail, &userExclusiveOffer.StartDate,
			&userExclusiveOffer.EndDate, &userExclusiveOffer.StartTime, &userExclusiveOffer.EndTime,
			&userExclusiveOffer.Repeats, &userExclusiveOffer.Status, &userExclusiveOffer.CreatedAt,
			&userExclusiveOffer.UpdatedAt, &userExclusiveOffer.CreatedByID, &userExclusiveOffer.UpdatedByID,
		)
		if err != nil {
			return nil, err
		}

		userExclusiveOffers = append(userExclusiveOffers, userExclusiveOffer)
	}
	return userExclusiveOffers, nil
}

// GetUserExclusiveOfferByID fetches a user exclusive offer based on its ID.
func (repo *OffersRepository) GetUserExclusiveOfferByID(offerID string) (*models.UserExclusiveOffer, error) {
	var userExclusiveOffer models.UserExclusiveOffer
	query := `
		SELECT *
		FROM user_exclusive_offers
		WHERE id = $1
	`
	err := repo.db.Pool.QueryRow(context.Background(), query, offerID).Scan(
		&userExclusiveOffer.ID, &userExclusiveOffer.UserID, &userExclusiveOffer.BreweryID, &userExclusiveOffer.Title,
		&userExclusiveOffer.Description, &userExclusiveOffer.Thumbnail, &userExclusiveOffer.StartDate,
		&userExclusiveOffer.EndDate, &userExclusiveOffer.StartTime, &userExclusiveOffer.EndTime,
		&userExclusiveOffer.Repeats, &userExclusiveOffer.Status, &userExclusiveOffer.CreatedAt,
		&userExclusiveOffer.UpdatedAt, &userExclusiveOffer.CreatedByID, &userExclusiveOffer.UpdatedByID,
	)
	if err != nil {
		return nil, err
	}
	return &userExclusiveOffer, nil
}

// UpdateUserExclusiveOffer updates an existing user exclusive offer.
func (repo *OffersRepository) UpdateUserExclusiveOffer(userExclusiveOffer models.UserExclusiveOffer) error {
	query := `
		UPDATE user_exclusive_offers
		SET user_id = $2, brewery_id = $3, title = $4, description = $5, thumbnail = $6,
			start_date = $7, end_date = $8, start_time = $9, end_time = $10, repeats = $11,
			status = $12, updated_by_id = $13
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(),
		query,
		userExclusiveOffer.ID, userExclusiveOffer.UserID, userExclusiveOffer.BreweryID, userExclusiveOffer.Title,
		userExclusiveOffer.Description, userExclusiveOffer.Thumbnail, userExclusiveOffer.StartDate,
		userExclusiveOffer.EndDate, userExclusiveOffer.StartTime, userExclusiveOffer.EndTime,
		userExclusiveOffer.Repeats, userExclusiveOffer.Status, userExclusiveOffer.UpdatedByID,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserExclusiveOffer deletes a user exclusive offer based on its ID.
func (repo *OffersRepository) DeleteUserExclusiveOffer(offerID string) error {
	query := `
		DELETE FROM user_exclusive_offers
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, offerID)
	if err != nil {
		return err
	}
	return nil
}

// CreateUserOfferRedemption creates a new user offer redemption.
func (repo *OffersRepository) CreateUserOfferRedemption(userOfferRedemption models.UserOfferRedemption) error {
	query := `
		INSERT INTO user_offer_redemptions (offer_id, user_id, redemption_date, notes)
		VALUES ($1, $2, $3, $4)
	`
	_, err := repo.db.Pool.Exec(context.Background(),
		query,
		userOfferRedemption.OfferID, userOfferRedemption.UserID, userOfferRedemption.RedemptionDate, userOfferRedemption.Notes,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetUserOfferRedemptionByID fetches a user offer redemption based on its ID.
func (repo *OffersRepository) GetUserOfferRedemptionByID(redemptionID string) (*models.UserOfferRedemption, error) {
	var userOfferRedemption models.UserOfferRedemption
	query := `
		SELECT *
		FROM user_offer_redemptions
		WHERE id = $1
	`
	err := repo.db.Pool.QueryRow(context.Background(), query, redemptionID).Scan(
		&userOfferRedemption.ID, &userOfferRedemption.OfferID, &userOfferRedemption.UserID,
		&userOfferRedemption.RedemptionDate, &userOfferRedemption.Notes, &userOfferRedemption.CreatedAt,
		&userOfferRedemption.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &userOfferRedemption, nil
}

// DeleteUserOfferRedemption deletes a user offer redemption based on its ID.
func (repo *OffersRepository) DeleteUserOfferRedemption(redemptionID string) error {
	query := `
		DELETE FROM user_offer_redemptions
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, redemptionID)
	if err != nil {
		return err
	}
	return nil
}
