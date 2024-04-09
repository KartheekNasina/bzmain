package service

import (
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type DrinkPurchaseRequestService struct {
	drinkPurchaseRequestRepo *repository.DrinkPurchaseRequestRepository
}

func NewDrinkPurchaseRequestService(r *repository.DrinkPurchaseRequestRepository) *DrinkPurchaseRequestService {
	return &DrinkPurchaseRequestService{drinkPurchaseRequestRepo: r}
}

// Add your service functions here.
