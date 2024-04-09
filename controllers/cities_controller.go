// controllers/city_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type CityController struct {
	cityService *service.CityService
}

// Constructor function for CityController
func NewCityController(s *service.CityService) *CityController {
	return &CityController{cityService: s}
}

func (cc *CityController) ListCities(c *gin.Context) {
	cities := cc.cityService.ListCities()
	c.JSON(200, cities)
}
