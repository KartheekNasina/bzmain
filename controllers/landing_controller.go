package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type LandingController struct {
	LandingService *service.LandingService
}

// Constructor function for UserController
func NewLandingController(s *service.LandingService) *LandingController {
	return &LandingController{LandingService: s}
}

func (bc *LandingController) FetchLandingData(c *gin.Context) {
	Landings := bc.LandingService.FetchLandingData()
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "Failed to fetch Landings"})
	// 	return
	// }
	c.JSON(200, Landings)
}
