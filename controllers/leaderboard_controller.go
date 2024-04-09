package controllers

import (
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type LeaderboardController struct {
	LeaderboardService *service.LeaderboardService
}

// Constructor function for LeaderboardController
func NewLeaderboardController(s *service.LeaderboardService) *LeaderboardController {
	return &LeaderboardController{LeaderboardService: s}
}

// Add your controller functions here
