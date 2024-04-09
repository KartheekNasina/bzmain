package controllers

import (
	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type UserNotificationController struct {
	userNotificationService *service.UserNotificationService
}

// Constructor function for UserNotificationController
func NewUserNotificationController(s *service.UserNotificationService) *UserNotificationController {
	return &UserNotificationController{userNotificationService: s}
}

func (uc *UserNotificationController) CreateUserNotification(c *gin.Context) {
	var notificationDTO dto.UserNotificationDTO
	if err := c.BindJSON(&notificationDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userNotificationService.CreateUserNotification(&notificationDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User Notification created successfully"})
}

func (uc *UserNotificationController) GetUserNotificationByID(c *gin.Context) {
	notificationID := c.Param("notificationID")

	userNotification, err := uc.userNotificationService.GetUserNotificationByID(notificationID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, userNotification)
}

func (uc *UserNotificationController) UpdateUserNotification(c *gin.Context) {
	var notificationDTO dto.UserNotificationDTO
	if err := c.BindJSON(&notificationDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userNotificationService.UpdateUserNotification(&notificationDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User Notification updated successfully"})
}

func (uc *UserNotificationController) DeleteUserNotification(c *gin.Context) {
	notificationID := c.Param("notificationID")

	err := uc.userNotificationService.DeleteUserNotification(notificationID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User Notification deleted successfully"})
}
