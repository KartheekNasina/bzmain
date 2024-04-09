// brewery_details_repository.go
package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	common "github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type BreweryRepository struct {
	db *driver.DB
}

func NewBreweryRepository(database *driver.DB) *BreweryRepository {
	return &BreweryRepository{
		db: database,
	}
}

// GetBrewery fetches brewery details by brewery ID.
func (repo *BreweryRepository) GetBrewery(breweryID string) (*models.Brewery, error) {
	var Brewery models.Brewery

	// Replace "brewery_details" with your actual table name.
	query := `
		SELECT *
		FROM brewery_details
		WHERE id = $1
		LIMIT 1
	`

	row := repo.db.Pool.QueryRow(context.Background(), query, breweryID)

	err := common.ScanRowToModel(row, &Brewery)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_details_repository",
			"function": "GetBrewery",
		}).Errorf("Unable to scan the row. %v", err)

		return nil, err
	}

	return &Brewery, nil
}

// GetBreweries fetches a list of breweries with limit and offset.
func (repo *BreweryRepository) GetBreweries(limit, offset int) ([]models.Brewery, error) {
	var breweries []models.Brewery

	// Replace "brewery_details" with your actual table name.
	query := `
		SELECT *
		FROM brewery_details
		LIMIT $1 OFFSET $2
	`

	rows, err := repo.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "brewery_details_repository",
			"function": "GetBreweries",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Brewery models.Brewery

		err := common.ScanRowsToModel(rows, &Brewery)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "brewery_details_repository",
				"function": "GetBreweries",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		breweries = append(breweries, Brewery)
	}

	return breweries, nil
}
