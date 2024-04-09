package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type BrewTourController struct {
	brewTourService *service.BrewTourService
}

// Constructor function for BrewToursController
func NewBrewToursController(s *service.BrewTourService) *BrewTourController {
	return &BrewTourController{brewTourService: s}
}

// BrewTourController.go
func (btc *BrewTourController) GetTours(c *gin.Context) {
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

	tours, err := btc.brewTourService.GetTours(limit, offset)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tours)
}

// ... (The previous controller functions you provided for Update, Delete, and Create are also included here.)

// BrewTourController.go
func (btc *BrewTourController) UpdateUserBrewTourRegistration(c *gin.Context) {
	var registrationDTO dto.UserBrewTourRegistrationDTO
	if err := c.BindJSON(&registrationDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	err := btc.brewTourService.UpdateUserBrewTourRegistration(id, &registrationDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User Brew Tour Registration updated successfully"})
}

func (btc *BrewTourController) DeleteUserBrewTourRegistration(c *gin.Context) {
	registrationID := c.Param("registrationID")

	err := btc.brewTourService.DeleteUserBrewTourRegistration(registrationID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User Brew Tour Registration deleted successfully"})
}

func (btc *BrewTourController) CreateUserBrewTourRegistration(c *gin.Context) {
	var registrationDTO dto.UserBrewTourRegistrationDTO
	if err := c.BindJSON(&registrationDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := btc.brewTourService.CreateUserBrewTourRegistration(&registrationDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User Brew Tour Registration created successfully"})
}
