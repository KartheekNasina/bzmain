// repository/beer_repository.go
package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	common "github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type BeerRepository struct {
	db *driver.DB
}

func NewBeerRepository(database *driver.DB) *BeerRepository {
	return &BeerRepository{
		db: database,
	}
}

// GetBeer retrieves a beer by its ID where IsPublished is true.
func (repo *BeerRepository) GetBeer(id string) (*models.Beer, error) {
	var beer models.Beer

	// Replace "beers" with your actual table name.
	query := `
		SELECT id, name, type, image_url, created_at, updated_at, is_published, published_at, brewery_id, abv, ibu , description
		FROM beers
		WHERE id = $1 AND is_published = true
	`

	row := repo.db.Pool.QueryRow(context.Background(), query, id)

	err := common.ScanRowToModel(row, &beer)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "beer_repository",
			"function": "GetBeer",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}

	return &beer, nil
}

// GetBeersBasedOnType fetches beers of a specific type where IsPublished is true.
func (repo *BeerRepository) GetBeersBasedOnType(beerType string, limit, offset int) ([]models.Beer, error) {
	var beers []models.Beer

	// Replace "beers" with your actual table name.
	query := `
		SELECT id, name, type, image_url, created_at, updated_at, is_published, published_at, brewery_id, abv, ibu , description
		FROM beers
		WHERE type = $1 AND is_published = true
		LIMIT $2 OFFSET $3
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, beerType, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "beer_repository",
			"function": "GetBeersBasedOnType",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var beer models.Beer

		err := common.ScanRowsToModel(rows, &beer)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "beer_repository",
				"function": "GetBeersBasedOnType",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		beers = append(beers, beer)
	}
	return beers, nil
}

// GetBeersBasedOnBreweryID fetches beers for a specific brewery where IsPublished is true.
func (repo *BeerRepository) GetBeersBasedOnBreweryID(breweryID string, limit, offset int) ([]models.Beer, error) {
	var beers []models.Beer

	// Replace "beers" with your actual table name.
	query := `
		SELECT id, name, type, image_url, created_at, updated_at, is_published, published_at, brewery_id, abv, ibu , description
		FROM beers
		WHERE brewery_id = $1 AND is_published = true
		LIMIT $2 OFFSET $3
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, breweryID, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "beer_repository",
			"function": "GetBeersBasedOnBreweryID",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var beer models.Beer

		err := common.ScanRowsToModel(rows, &beer)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "beer_repository",
				"function": "GetBeersBasedOnBreweryID",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		beers = append(beers, beer)
	}

	return beers, nil
}
