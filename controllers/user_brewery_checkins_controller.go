package controllers

import (
	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type UserBreweryCheckinsController struct {
	userBreweryCheckinsService *service.UserBreweryCheckinsService
}

// Constructor function for UserBreweryCheckinsController
func NewUserBreweryCheckinsController(s *service.UserBreweryCheckinsService) *UserBreweryCheckinsController {
	return &UserBreweryCheckinsController{userBreweryCheckinsService: s}
}

func (ubc *UserBreweryCheckinsController) CreateUserBreweryCheckin(c *gin.Context) {
	var checkinDTO dto.UserBreweryCheckinDTO
	if err := c.BindJSON(&checkinDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := ubc.userBreweryCheckinsService.CreateUserBreweryCheckin(&checkinDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "UserBreweryCheckin created successfully"})
}

func (ubc *UserBreweryCheckinsController) UpdateUserBreweryCheckin(c *gin.Context) {
	id := c.Param("id")
	var checkinDTO dto.UserBreweryCheckinDTO
	if err := c.BindJSON(&checkinDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := ubc.userBreweryCheckinsService.UpdateUserBreweryCheckin(id, &checkinDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "UserBreweryCheckin updated successfully"})
}

func (ubc *UserBreweryCheckinsController) DeleteUserBreweryCheckin(c *gin.Context) {
	id := c.Param("id")
	err := ubc.userBreweryCheckinsService.DeleteUserBreweryCheckin(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "UserBreweryCheckin deleted successfully"})
}

func (ubc *UserBreweryCheckinsController) ListUserBreweryCheckins(c *gin.Context) {
	checkins, err := ubc.userBreweryCheckinsService.ListUserBreweryCheckins()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, checkins)
}
