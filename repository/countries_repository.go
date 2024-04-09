package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	common "github.com/vivekbnwork/bz-backend/bz-main/common"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

// CountryRepository handles database operations for Country.
type CountryRepository struct {
	db *driver.DB
}

// NewCountryRepository creates a new CountryRepository.
func NewCountryRepository(database *driver.DB) *CountryRepository {
	return &CountryRepository{
		db: database,
	}
}

// ListCountries fetches a list of countries.
func (repo *CountryRepository) ListCountries() ([]models.Country, error) {
	var countries []models.Country

	// Replace "countries" with your actual table name.
	query := "SELECT id, name, code, created_at, updated_at FROM countries"
	rows, err := repo.db.Pool.Query(context.Background(), query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "country_repository",
			"function": "ListCountries",
		}).Errorf("Unable to execute the query. %v", err)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var country models.Country
		err := common.ScanRowsToModel(rows, &country)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":  "bz-main",
				"file":     "country_repository",
				"function": "ListCountries",
			}).Errorf("Unable to scan the row. %v", err)

			return nil, err
		}

		countries = append(countries, country)
	}

	return countries, nil
}
