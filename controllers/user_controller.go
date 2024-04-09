package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type UserController struct {
	userService *service.UserService
}

// Constructor function for UserController
func NewUserController(s *service.UserService) *UserController {
	return &UserController{userService: s}
}

func (uc *UserController) ListUsers(c *gin.Context) {
	users, err := uc.userService.ListUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	// Extract the user ID from the URL parameter
	userID := c.Param("id")

	var userDTO dto.UserDTO
	if err := c.BindJSON(&userDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Pass the user ID and userDTO to the service for updating
	err := uc.userService.UpdateUser(userID, &userDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	err := uc.userService.DeleteUser(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userDTO dto.UserDTO
	if err := c.BindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponseDTO{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Result:  "FAILURE",
		})
		return
	}

	err := uc.userService.CreateUser(&userDTO)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(http.StatusConflict, dto.APIResponseDTO{
				Status:  http.StatusConflict,
				Message: err.Error(),
				Result:  "SUCCESS",
			})
		} else {
			c.JSON(http.StatusInternalServerError, dto.APIResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Result:  "FAILURE",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, dto.APIResponseDTO{
		Status:  http.StatusCreated,
		Message: "User created successfully",
		Result:  "SUCCESS",
	})
}

func (uc *UserController) GetUsers(c *gin.Context) {
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

	users, err := uc.userService.GetUsers(limit, offset)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}
