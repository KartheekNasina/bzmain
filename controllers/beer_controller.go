package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type BeerController struct {
	beerService *service.BeerService
}

// Constructor function for BeerController
func NewBeerController(s *service.BeerService) *BeerController {
	return &BeerController{beerService: s}
}

func (uc *BeerController) GetBeer(c *gin.Context) {
	beerID := c.Param("id")
	beer, err := uc.beerService.GetBeer(beerID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, beer)
}

func (uc *BeerController) GetBeersBasedOnType(c *gin.Context) {
	beerType := c.Param("type")

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid offset parameter"})
		return
	}

	beers, err := uc.beerService.GetBeersBasedOnType(beerType, limit, offset)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, beers)
}

func (uc *BeerController) GetBeersBasedOnBreweryID(c *gin.Context) {
	breweryID := c.Param("id")

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid limit parameter"})
		return
	}

	fmt.Println(breweryID)

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid offset parameter"})
		return
	}

	beers, err := uc.beerService.GetBeersBasedOnBreweryID(breweryID, limit, offset)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, beers)
}
