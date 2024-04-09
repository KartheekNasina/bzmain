package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	common "github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

// CityRepository handles database operations for City.
type CityRepository struct {
	db *driver.DB
}

// NewCityRepository creates a new CityRepository.
func NewCityRepository(database *driver.DB) *CityRepository {
	return &CityRepository{
		db: database,
	}
}

// ListCities fetches a list of cities.
func (repo *CityRepository) ListCities() ([]models.City, error) {
	var cities []models.City

	// Replace "cities" with your actual table name.
	query := "SELECT id, name, country_id, created_at, updated_at FROM cities"
	rows, err := repo.db.Pool.Query(context.Background(), query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "city_repository",
			"function": "ListCities",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var city models.City
		err := common.ScanRowsToModel(rows, &city)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "city_repository",
				"function": "ListCities",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		cities = append(cities, city)
	}

	return cities, nil
}
