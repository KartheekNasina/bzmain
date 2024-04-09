// brewery_offers_repository.go
package repository

import (
	"context"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type BreweryOffersRepository struct {
	db *driver.DB
}

func NewBreweryOffersRepository(database *driver.DB) *BreweryOffersRepository {
	return &BreweryOffersRepository{
		db: database,
	}
}

// GetBreweryOffersByBreweryID fetches active brewery offers by brewery ID.
func (repo *BreweryOffersRepository) GetBreweryOffersByBreweryID(breweryID string) ([]models.BreweryOffer, error) {
	var breweryOffers []models.BreweryOffer

	// Replace "brewery_offers" with your actual table name.
	query := `
		SELECT *
		FROM brewery_offers
		WHERE brewery_id = $1 AND status = 'active'
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, breweryID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_offers_repository",
			"function": "GetBreweryOffersByBreweryID",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var breweryOffer models.BreweryOffer

		err = rows.Scan(
			&breweryOffer.ID, &breweryOffer.BreweryID, &breweryOffer.OfferType, &breweryOffer.Title,
			&breweryOffer.Description, &breweryOffer.Thumbnail, pq.Array(breweryOffer.Images),
			&breweryOffer.StartDate, &breweryOffer.EndDate, &breweryOffer.StartTime, &breweryOffer.EndTime,
			&breweryOffer.Repeats, &breweryOffer.Status, &breweryOffer.CreatedAt, &breweryOffer.UpdatedAt,
			&breweryOffer.CreatedByID, &breweryOffer.UpdatedByID,
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "brewery_offers_repository",
				"function": "GetBreweryOffersByBreweryID",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		breweryOffers = append(breweryOffers, breweryOffer)
	}

	return breweryOffers, nil
}

// GetBreweryOffers fetches a list of brewery offers with limit and offset.
func (repo *BreweryOffersRepository) GetBreweryOffers(limit, offset int) ([]models.BreweryOffer, error) {
	var breweryOffers []models.BreweryOffer

	// Replace "brewery_offers" with your actual table name.
	query := `
		SELECT *
		FROM brewery_offers
		LIMIT $1 OFFSET $2
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_offers_repository",
			"function": "GetBreweryOffers",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var breweryOffer models.BreweryOffer

		err = rows.Scan(
			&breweryOffer.ID, &breweryOffer.BreweryID, &breweryOffer.OfferType, &breweryOffer.Title,
			&breweryOffer.Description, &breweryOffer.Thumbnail, pq.Array(breweryOffer.Images),
			&breweryOffer.StartDate, &breweryOffer.EndDate, &breweryOffer.StartTime, &breweryOffer.EndTime,
			&breweryOffer.Repeats, &breweryOffer.Status, &breweryOffer.CreatedAt, &breweryOffer.UpdatedAt,
			&breweryOffer.CreatedByID, &breweryOffer.UpdatedByID,
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "brewery_offers_repository",
				"function": "GetBreweryOffers",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		breweryOffers = append(breweryOffers, breweryOffer)
	}

	return breweryOffers, nil
}
