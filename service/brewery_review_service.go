package service

import (
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type BreweryReviewService struct {
	breweryReviewRepo *repository.BreweryReviewRepository
}

func NewBreweryReviewService(r *repository.BreweryReviewRepository) *BreweryReviewService {
	return &BreweryReviewService{breweryReviewRepo: r}
}

// Add your service functions here.
