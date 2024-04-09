package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

// DrinkRequestStatus represents the status of a drink purchase request.
type DrinkRequestStatus string

const (
	// DrinkRequestPending indicates that the request is pending.
	DrinkRequestPending DrinkRequestStatus = "pending"
	// DrinkRequestAccepted indicates that the request has been accepted.
	DrinkRequestAccepted DrinkRequestStatus = "accepted"
	// DrinkRequestDeclined indicates that the request has been declined.
	DrinkRequestDeclined DrinkRequestStatus = "declined"
)

// DrinkPurchaseRequestRepository handles database operations for DrinkPurchaseRequest.
type DrinkPurchaseRequestRepository struct {
	db *driver.DB
}

// NewDrinkPurchaseRequestRepository creates a new DrinkPurchaseRequestRepository.
func NewDrinkPurchaseRequestRepository(database *driver.DB) *DrinkPurchaseRequestRepository {
	return &DrinkPurchaseRequestRepository{
		db: database,
	}
}

// CreateDrinkPurchaseRequest inserts a new drink purchase request into the database.
func (repo *DrinkPurchaseRequestRepository) CreateDrinkPurchaseRequest(request models.DrinkPurchaseRequest) error {
	query := `
		INSERT INTO drink_purchase_requests (buyer_id, recipient_id, message_id, request_date, status)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := repo.db.Pool.Exec(context.Background(), query,
		request.BuyerID, request.RecipientID, request.MessageID, request.RequestDate, request.Status)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_request_repository",
			"function": "CreateDrinkPurchaseRequest",
		}).Errorf("Unable to create drink purchase request. %v", err)

		return err
	}

	return nil
}

// GetDrinkPurchaseRequestByID retrieves a drink purchase request by its ID.
func (repo *DrinkPurchaseRequestRepository) GetDrinkPurchaseRequestByID(requestID string) (*models.DrinkPurchaseRequest, error) {
	var request models.DrinkPurchaseRequest

	query := "SELECT id, buyer_id, recipient_id, message_id, request_date, status FROM drink_purchase_requests WHERE id = $1"
	err := repo.db.Pool.QueryRow(context.Background(), query, requestID).Scan(
		&request.ID, &request.BuyerID, &request.RecipientID, &request.MessageID,
		&request.RequestDate, &request.Status)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_request_repository",
			"function": "GetDrinkPurchaseRequestByID",
		}).Errorf("Unable to retrieve drink purchase request. %v", err)

		return nil, err
	}

	return &request, nil
}

// UpdateDrinkPurchaseRequest updates a drink purchase request in the database.
func (repo *DrinkPurchaseRequestRepository) UpdateDrinkPurchaseRequest(request models.DrinkPurchaseRequest) error {
	query := `
		UPDATE drink_purchase_requests
		SET buyer_id = $2, recipient_id = $3, message_id = $4, request_date = $5, status = $6
		WHERE id = $1
	`

	_, err := repo.db.Pool.Exec(context.Background(), query,
		request.ID, request.BuyerID, request.RecipientID, request.MessageID, request.RequestDate, request.Status)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_request_repository",
			"function": "UpdateDrinkPurchaseRequest",
		}).Errorf("Unable to update drink purchase request. %v", err)

		return err
	}

	return nil
}

// DeleteDrinkPurchaseRequest deletes a drink purchase request from the database by its ID.
func (repo *DrinkPurchaseRequestRepository) DeleteDrinkPurchaseRequest(requestID string) error {
	query := "DELETE FROM drink_purchase_requests WHERE id = $1"
	_, err := repo.db.Pool.Exec(context.Background(), query, requestID)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "bz-main",
			"file":     "user_request_repository",
			"function": "DeleteDrinkPurchaseRequest",
		}).Errorf("Unable to delete drink purchase request. %v", err)

		return err
	}

	return nil
}
