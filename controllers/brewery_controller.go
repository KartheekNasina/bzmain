// controllers/brewery_details_controller.go
package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type BreweryController struct {
	BreweryService *service.BreweryService
}

// Constructor function for BreweryController
func NewBreweryController(s *service.BreweryService) *BreweryController {
	return &BreweryController{BreweryService: s}
}

func (bc *BreweryController) GetBrewery(c *gin.Context) {
	breweryID := c.Param("id")
	Brewery, err := bc.BreweryService.GetBrewery(breweryID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Brewery details not found"})
		return
	}
	c.JSON(200, Brewery)
}

func (bc *BreweryController) GetBreweries(c *gin.Context) {
	limit, offset := parseLimitOffset(c)
	breweries, err := bc.BreweryService.GetBreweries(limit, offset)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch breweries"})
		return
	}
	c.JSON(200, breweries)
}

func parseLimitOffset(c *gin.Context) (int, int) {
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)
	return limitInt, offsetInt
}
