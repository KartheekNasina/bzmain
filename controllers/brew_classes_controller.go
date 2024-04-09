package controllers

import (
	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type BrewClassController struct {
	brewClassService *service.BrewClassService
}

// Constructor function for BrewClassesController
func NewBrewClassController(s *service.BrewClassService) *BrewClassController {
	return &BrewClassController{brewClassService: s}
}

func (uc *BrewClassController) GetUserBrewClassRegistration(c *gin.Context) {
	userID := c.Param("userID")
	classID := c.Param("classID")

	registration, err := uc.brewClassService.GetUserBrewClassRegistration(userID, classID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, registration)
}

func (uc *BrewClassController) UpdateUserBrewClassRegistration(c *gin.Context) {
	userID := c.Param("userID")
	//classID := c.Param("classID")

	var registrationDTO dto.UserBrewClassRegistrationDTO
	if err := c.BindJSON(&registrationDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.brewClassService.UpdateUserBrewClassRegistration(userID, &registrationDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User Brew Class Registration updated successfully"})
}

func (uc *BrewClassController) DeleteUserBrewClassRegistration(c *gin.Context) {
	registrationID := c.Param("registrationID")

	err := uc.brewClassService.DeleteUserBrewClassRegistration(registrationID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User Brew Class Registration deleted successfully"})
}

func (uc *BrewClassController) CreateUserBrewClassRegistration(c *gin.Context) {
	var registrationDTO dto.UserBrewClassRegistrationDTO
	if err := c.BindJSON(&registrationDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.brewClassService.CreateUserBrewClassRegistration(&registrationDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User Brew Class Registration created successfully"})
}
