package controllers

import (
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type DrinkPurchaseRequestController struct {
	DrinkPurchaseRequestService *service.DrinkPurchaseRequestService
}

// Constructor function for DrinkPurchaseRequestController
func NewDrinkPurchaseRequestController(s *service.DrinkPurchaseRequestService) *DrinkPurchaseRequestController {
	return &DrinkPurchaseRequestController{DrinkPurchaseRequestService: s}
}

// Add your controller functions here
