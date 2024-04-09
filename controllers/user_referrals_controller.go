package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type UserReferralController struct {
	userReferralService *service.UserReferralService
}

// Constructor function for UserReferralController
func NewUserReferralController(s *service.UserReferralService) *UserReferralController {
	return &UserReferralController{userReferralService: s}
}

func (uc *UserReferralController) GetUserReferral(c *gin.Context) {
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

	userReferrals, err := uc.userReferralService.GetUserReferral(limit, offset)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, userReferrals)
}

func (uc *UserReferralController) GetUserReferralByID(c *gin.Context) {
	id := c.Param("id")

	userReferral, err := uc.userReferralService.GetUserReferralByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, userReferral)
}

func (uc *UserReferralController) CreateUserReferral(c *gin.Context) {
	var referralDTO dto.UserReferralDTO
	if err := c.BindJSON(&referralDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userReferralService.CreateUserReferral(&referralDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User Referral created successfully"})
}

func (uc *UserReferralController) UpdateUserReferral(c *gin.Context) {
	id := c.Param("id")

	var referralDTO dto.UserReferralDTO
	if err := c.BindJSON(&referralDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userReferralService.UpdateUserReferral(id, &referralDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User Referral updated successfully"})
}

func (uc *UserReferralController) DeleteUserReferral(c *gin.Context) {
	id := c.Param("id")

	err := uc.userReferralService.DeleteUserReferral(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User Referral deleted successfully"})
}
