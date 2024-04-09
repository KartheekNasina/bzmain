package controllers

import (
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type PaymentController struct {
	PaymentService *service.PaymentService
}

// Constructor function for PaymentController
func NewPaymentController(s *service.PaymentService) *PaymentController {
	return &PaymentController{PaymentService: s}
}

// Add your controller functions here
