package service

import (
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type LeaderboardService struct {
	leaderboardRepo *repository.LeaderboardRepository
}

func NewLeaderboardService(r *repository.LeaderboardRepository) *LeaderboardService {
	return &LeaderboardService{leaderboardRepo: r}
}

// Add your service functions here.
