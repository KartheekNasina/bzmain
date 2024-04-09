package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type UserHistoryController struct {
	userHistoryService *service.UserHistoryService
}

// Constructor function for UserHistoryController
func NewUserHistoryController(s *service.UserHistoryService) *UserHistoryController {
	return &UserHistoryController{userHistoryService: s}
}

func (uc *UserHistoryController) GetUserHistory(c *gin.Context) {
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

	userHistories, err := uc.userHistoryService.GetUserHistory(limit, offset)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, userHistories)
}

func (uc *UserHistoryController) GetUserHistoryByID(c *gin.Context) {
	id := c.Param("id")

	userHistory, err := uc.userHistoryService.GetUserHistoryByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, userHistory)
}

func (uc *UserHistoryController) CreateUserHistory(c *gin.Context) {
	var historyDTO dto.UserHistoryDTO
	if err := c.BindJSON(&historyDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userHistoryService.CreateUserHistory(&historyDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User History created successfully"})
}

func (uc *UserHistoryController) UpdateUserHistory(c *gin.Context) {
	id := c.Param("id")

	var historyDTO dto.UserHistoryDTO
	if err := c.BindJSON(&historyDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userHistoryService.UpdateUserHistory(id, &historyDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User History updated successfully"})
}

func (uc *UserHistoryController) DeleteUserHistory(c *gin.Context) {
	id := c.Param("id")

	err := uc.userHistoryService.DeleteUserHistory(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User History deleted successfully"})
}
