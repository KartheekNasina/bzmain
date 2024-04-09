package service

import (
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type PaymentService struct {
	paymentsRepo *repository.PaymentsRepository
}

func NewPaymentService(r *repository.PaymentsRepository) *PaymentService {
	return &PaymentService{paymentsRepo: r}
}

// Add your service functions here.
