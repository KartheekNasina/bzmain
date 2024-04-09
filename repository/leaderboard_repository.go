// Leaderboard_repository.go
package repository

import (
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
)

type LeaderboardRepository struct {
	db *driver.DB
}

func NewLeaderboardRepository(database *driver.DB) *LeaderboardRepository {
	return &LeaderboardRepository{
		db: database,
	}
}
