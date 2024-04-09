package controllers

import (
	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type UserFeedbackController struct {
	userFeedbackService *service.UserFeedbackService
}

// Constructor function for UserFeedbackController
func NewUserFeedbackController(s *service.UserFeedbackService) *UserFeedbackController {
	return &UserFeedbackController{userFeedbackService: s}
}

func (uc *UserFeedbackController) CreateUserFeedback(c *gin.Context) {
	var feedback dto.UserFeedbackDTO
	if err := c.BindJSON(&feedback); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userFeedbackService.CreateUserFeedback(&feedback)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User feedback created successfully"})
}

// GetUserFeedbackByID handles the route for getting user feedback by ID.
func (uc *UserFeedbackController) GetUserFeedbackByID(c *gin.Context) {
	feedbackID := c.Param("id")

	feedback, err := uc.userFeedbackService.GetUserFeedbackByID(feedbackID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if feedback == nil {
		c.JSON(404, gin.H{"error": "User feedback not found"})
		return
	}

	c.JSON(200, feedback)
}

// UpdateUserFeedback handles the route for updating an existing user feedback entry.
func (uc *UserFeedbackController) UpdateUserFeedback(c *gin.Context) {
	feedbackID := c.Param("id")

	var feedback dto.UserFeedbackDTO
	if err := c.BindJSON(&feedback); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userFeedbackService.UpdateUserFeedback(feedbackID, &feedback)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User feedback updated successfully"})
}

// DeleteUserFeedback handles the route for deleting a user feedback entry by ID.
func (uc *UserFeedbackController) DeleteUserFeedback(c *gin.Context) {
	feedbackID := c.Param("id")

	err := uc.userFeedbackService.DeleteUserFeedback(feedbackID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User feedback deleted successfully"})
}
