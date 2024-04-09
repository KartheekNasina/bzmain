package controllers

import (
	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type UserFavoriteBreweryController struct {
	userFavoriteBreweryService *service.UserFavoriteBreweryService
}

// Constructor function for UserFavoriteBreweryController
func NewUserFavoriteBreweryController(s *service.UserFavoriteBreweryService) *UserFavoriteBreweryController {
	return &UserFavoriteBreweryController{userFavoriteBreweryService: s}
}

// CreateUserFavoriteBrewery handles the route for creating a new user's favorite brewery entry.
func (uc *UserFavoriteBreweryController) CreateUserFavoriteBrewery(c *gin.Context) {
	var favoriteBrewery dto.UserFavoriteBreweryDTO
	if err := c.BindJSON(&favoriteBrewery); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := uc.userFavoriteBreweryService.CreateUserFavoriteBrewery(&favoriteBrewery)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User's favorite brewery created successfully"})
}

// GetUserFavoriteBrewery handles the route for retrieving a user's favorite breweries by userID.
func (uc *UserFavoriteBreweryController) GetUserFavoriteBrewery(c *gin.Context) {
	userID := c.Param("id")

	favoriteBreweries, err := uc.userFavoriteBreweryService.GetUserFavoriteBreweryByID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, favoriteBreweries)
}

// DeleteUserFavoriteBrewery handles the route for deleting a user's favorite brewery entry by userID and breweryID.
func (uc *UserFavoriteBreweryController) DeleteUserFavoriteBrewery(c *gin.Context) {
	userID := c.Param("userID")
	breweryID := c.Param("breweryID")

	err := uc.userFavoriteBreweryService.DeleteUserFavoriteBrewery(userID, breweryID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User's favorite brewery deleted successfully"})
}
