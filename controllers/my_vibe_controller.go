package controllers

import (
	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type MyVibeController struct {
	myVibeService *service.MyVibeService
}

// Constructor function for MyVibeController
func NewMyVibeController(s *service.MyVibeService) *MyVibeController {
	return &MyVibeController{myVibeService: s}
}

func (vc *MyVibeController) CreateFoodDrinkRating(c *gin.Context) {
	var ratingDTO dto.FoodDrinkRatingDTO
	if err := c.BindJSON(&ratingDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := vc.myVibeService.CreateFoodDrinkRating(&ratingDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Rating created successfully"})
}

func (vc *MyVibeController) GetFoodDrinkRatingByID(c *gin.Context) {
	id := c.Param("id")

	rating, err := vc.myVibeService.GetFoodDrinkRatingByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, rating)
}

// ... (repeat similar structure for the other functions)
