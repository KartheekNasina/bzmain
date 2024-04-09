package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type UserSubscriptionController struct {
	userSubscriptionService *service.UserSubscriptionService
}

// Constructor function for UserSubscriptionController
func NewUserSubscriptionController(s *service.UserSubscriptionService) *UserSubscriptionController {
	return &UserSubscriptionController{userSubscriptionService: s}
}
func (uc *UserSubscriptionController) CreateUserSubscription(c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserSubscription",
	}).Debug("Create User Subscription - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateUserSubscription",
	}).Debug("Create User Subscription - End")

	var subscriptionDTO dto.UserSubscriptionDTO
	if err := c.BindJSON(&subscriptionDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userSubscriptionService.CreateUserSubscription(&subscriptionDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User Subscription created successfully"})
}

func (uc *UserSubscriptionController) GetUserSubscriptionByUserID(c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetUserSubscriptionByUserID",
	}).Debug("Get User Subscription By UserID - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetUserSubscriptionByUserID",
	}).Debug("Get User Subscription By UserID - End")

	userID := c.Param("userID")

	subscription, err := uc.userSubscriptionService.GetUserSubscriptionByUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, subscription)
}

func (uc *UserSubscriptionController) UpdateUserSubscription(c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserSubscription",
	}).Debug("Update User Subscription - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateUserSubscription",
	}).Debug("Update User Subscription - End")

	userID := c.Param("userID")

	var subscriptionDTO dto.UserSubscriptionDTO
	if err := c.BindJSON(&subscriptionDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userSubscriptionService.UpdateUserSubscription(userID, &subscriptionDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User Subscription updated successfully"})
}

func (uc *UserSubscriptionController) DeleteUserSubscription(c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserSubscription",
	}).Debug("Delete User Subscription - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteUserSubscription",
	}).Debug("Delete User Subscription - End")

	userID := c.Param("userID")

	err := uc.userSubscriptionService.DeleteUserSubscription(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User Subscription deleted successfully"})
}
