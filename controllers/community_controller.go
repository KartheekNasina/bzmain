package controllers

import (
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type CommunityController struct {
	CommunityService *service.CommunityService
}

// Constructor function for CommunityController
func NewCommunityController(s *service.CommunityService) *CommunityController {
	return &CommunityController{CommunityService: s}
}

// Add your controller functions here
