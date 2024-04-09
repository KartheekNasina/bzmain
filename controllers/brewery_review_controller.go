package controllers

import (
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type BreweryReviewController struct {
	BreweryReviewService *service.BreweryReviewService
}

// Constructor function for BreweryReviewController
func NewBreweryReviewController(s *service.BreweryReviewService) *BreweryReviewController {
	return &BreweryReviewController{BreweryReviewService: s}
}

// Add your controller functions here
