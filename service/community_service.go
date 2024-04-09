package service

import (
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type CommunityService struct {
	communityRepo *repository.CommunityRepository
}

func NewCommunityService(r *repository.CommunityRepository) *CommunityService {
	return &CommunityService{communityRepo: r}
}

// Add your service functions here.
