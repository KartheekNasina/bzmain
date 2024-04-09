package controllers

import (
	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type UserBlockedListController struct {
	userBlockedListService *service.UserBlockedListService
}

// Constructor function for UserBlockedListController
func NewUserBlockedListController(s *service.UserBlockedListService) *UserBlockedListController {
	return &UserBlockedListController{userBlockedListService: s}
}

func (uc *UserBlockedListController) CreateUserBlocked(c *gin.Context) {
	var userBlockedDTO dto.UserBlockedListDTO
	if err := c.BindJSON(&userBlockedDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userBlockedListService.CreateUserBlocked(&userBlockedDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "UserBlocked created successfully"})
}

func (uc *UserBlockedListController) GetUserBlockedByUserID(c *gin.Context) {
	userID := c.Param("userId")

	userBlockeds, err := uc.userBlockedListService.GetUserBlockedByUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, userBlockeds)
}

func (uc *UserBlockedListController) DeleteUserBlocked(c *gin.Context) {
	userID := c.Param("userId")
	blockedUserID := c.Param("blockedUserId")

	err := uc.userBlockedListService.DeleteUserBlocked(userID, blockedUserID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "UserBlocked deleted successfully"})
}
