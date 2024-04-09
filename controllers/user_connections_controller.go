package controllers

import (
	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type UserConnectionController struct {
	userConnectionService *service.UserConnectionService
}

// Constructor function for UserConnectionsController
func NewUserConnectionsController(s *service.UserConnectionService) *UserConnectionController {
	return &UserConnectionController{userConnectionService: s}
}

func (ucc *UserConnectionController) CreateUserConnection(c *gin.Context) {
	var userConnectionDTO dto.UserConnectionDTO
	if err := c.BindJSON(&userConnectionDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := ucc.userConnectionService.CreateUserConnection(&userConnectionDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "UserConnection created successfully"})
}

func (ucc *UserConnectionController) UpdateUserConnection(c *gin.Context) {
	id := c.Param("id")
	var userConnectionDTO dto.UserConnectionDTO
	if err := c.BindJSON(&userConnectionDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := ucc.userConnectionService.UpdateUserConnection(id, &userConnectionDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "UserConnection updated successfully"})
}

func (ucc *UserConnectionController) DeleteUserConnection(c *gin.Context) {
	id := c.Param("id")
	err := ucc.userConnectionService.DeleteUserConnection(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "UserConnection deleted successfully"})
}
