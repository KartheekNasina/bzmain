package repository

import (
	"context"

	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type PaymentsRepository struct {
	db *driver.DB
}

func NewPaymentsRepository(database *driver.DB) *PaymentsRepository {
	return &PaymentsRepository{
		db: database,
	}
}

// CreatePaymentAssociation creates a new payment association record.
func (repo *PaymentsRepository) CreatePaymentAssociation(paymentAssoc models.PaymentAssociation) error {
	query := `
		INSERT INTO payment_association (payment_transaction_id, entity_type, entity_id)
		VALUES ($1, $2, $3)
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, paymentAssoc.PaymentTransactionID, paymentAssoc.EntityType, paymentAssoc.EntityID)
	if err != nil {
		return err
	}
	return nil
}

// GetPaymentAssociationByID fetches a payment association record by ID.
func (repo *PaymentsRepository) GetPaymentAssociationByID(id string) (models.PaymentAssociation, error) {
	var paymentAssoc models.PaymentAssociation
	query := `
		SELECT payment_transaction_id, entity_type, entity_id
		FROM payment_association
		WHERE id = $1
	`
	row := repo.db.Pool.QueryRow(context.Background(), query, id)
	err := row.Scan(&paymentAssoc.PaymentTransactionID, &paymentAssoc.EntityType, &paymentAssoc.EntityID)
	if err != nil {
		return models.PaymentAssociation{}, err
	}
	return paymentAssoc, nil
}

// UpdatePaymentAssociation updates an existing payment association record.
func (repo *PaymentsRepository) UpdatePaymentAssociation(id string, paymentAssoc models.PaymentAssociation) error {
	query := `
		UPDATE payment_association
		SET payment_transaction_id = $1, entity_type = $2, entity_id = $3
		WHERE id = $4
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, paymentAssoc.PaymentTransactionID, paymentAssoc.EntityType, paymentAssoc.EntityID, id)
	if err != nil {
		return err
	}
	return nil
}

// DeletePaymentAssociation deletes a payment association record by ID.
func (repo *PaymentsRepository) DeletePaymentAssociation(id string) error {
	query := `
		DELETE FROM payment_association
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}

// CreatePaymentLog creates a new payment log record.
func (repo *PaymentsRepository) CreatePaymentLog(paymentLog models.PaymentLog) error {
	query := `
		INSERT INTO payment_log (payment_transaction_id, interaction, payload, timestamp, endpoint)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, paymentLog.PaymentTransactionID, paymentLog.Interaction, paymentLog.Payload, paymentLog.Timestamp, paymentLog.Endpoint)
	if err != nil {
		return err
	}
	return nil
}

// GetPaymentLogByID fetches a payment log record by ID.
func (repo *PaymentsRepository) GetPaymentLogByID(id string) (models.PaymentLog, error) {
	var paymentLog models.PaymentLog
	query := `
		SELECT payment_transaction_id, interaction, payload, timestamp, endpoint
		FROM payment_log
		WHERE id = $1
	`
	row := repo.db.Pool.QueryRow(context.Background(), query, id)
	err := row.Scan(&paymentLog.PaymentTransactionID, &paymentLog.Interaction, &paymentLog.Payload, &paymentLog.Timestamp, &paymentLog.Endpoint)
	if err != nil {
		return models.PaymentLog{}, err
	}
	return paymentLog, nil
}

// UpdatePaymentLog updates an existing payment log record.
func (repo *PaymentsRepository) UpdatePaymentLog(id string, paymentLog models.PaymentLog) error {
	query := `
		UPDATE payment_log
		SET payment_transaction_id = $1, interaction = $2, payload = $3, timestamp = $4, endpoint = $5
		WHERE id = $6
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, paymentLog.PaymentTransactionID, paymentLog.Interaction, paymentLog.Payload, paymentLog.Timestamp, paymentLog.Endpoint, id)
	if err != nil {
		return err
	}
	return nil
}

// DeletePaymentLog deletes a payment log record by ID.
func (repo *PaymentsRepository) DeletePaymentLog(id string) error {
	query := `
		DELETE FROM payment_log
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}

// CreatePaymentRefund creates a new payment refund record.
func (repo *PaymentsRepository) CreatePaymentRefund(paymentRefund models.PaymentRefund) error {
	query := `
		INSERT INTO payment_refund (payment_transaction_id, amount, refund_date, razorpay_refund_id, reason)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, paymentRefund.ID, paymentRefund.Amount, paymentRefund.RefundDate, paymentRefund.RazorpayRefundID, paymentRefund.Reason)
	if err != nil {
		return err
	}
	return nil
}

// GetPaymentRefundByID fetches a payment refund record by ID.
func (repo *PaymentsRepository) GetPaymentRefundByID(id string) (models.PaymentRefund, error) {
	var paymentRefund models.PaymentRefund
	query := `
		SELECT payment_transaction_id, amount, refund_date, razorpay_refund_id, reason
		FROM payment_refund
		WHERE id = $1
	`
	row := repo.db.Pool.QueryRow(context.Background(), query, id)
	err := row.Scan(&paymentRefund.ID, &paymentRefund.Amount, &paymentRefund.RefundDate, &paymentRefund.RazorpayRefundID, &paymentRefund.Reason)
	if err != nil {
		return models.PaymentRefund{}, err
	}
	return paymentRefund, nil
}

// UpdatePaymentRefund updates an existing payment refund record.
func (repo *PaymentsRepository) UpdatePaymentRefund(id string, paymentRefund models.PaymentRefund) error {
	query := `
		UPDATE payment_refund
		SET payment_transaction_id = $1, amount = $2, refund_date = $3, razorpay_refund_id = $4, reason = $5
		WHERE id = $6
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, paymentRefund.ID, paymentRefund.Amount, paymentRefund.RefundDate, paymentRefund.RazorpayRefundID, paymentRefund.Reason, id)
	if err != nil {
		return err
	}
	return nil
}

// DeletePaymentRefund deletes a payment refund record by ID.
func (repo *PaymentsRepository) DeletePaymentRefund(id string) error {
	query := `
		DELETE FROM payment_refund
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}

// CreatePaymentTransaction creates a new payment transaction record.
func (repo *PaymentsRepository) CreatePaymentTransaction(paymentTransaction models.PaymentTransaction) error {
	query := `
		INSERT INTO payment_transaction (user_id, amount, status, created_at, updated_at, currency_code, razorpay_payment_id, order_description)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, paymentTransaction.UserID, paymentTransaction.Amount, paymentTransaction.Status, paymentTransaction.CreatedAt, paymentTransaction.UpdatedAt, paymentTransaction.CurrencyCode, paymentTransaction.RazorpayPaymentID, paymentTransaction.OrderDescription)
	if err != nil {
		return err
	}
	return nil
}

// GetPaymentTransactionByID fetches a payment transaction record by ID.
func (repo *PaymentsRepository) GetPaymentTransactionByID(id string) (models.PaymentTransaction, error) {
	var paymentTransaction models.PaymentTransaction
	query := `
		SELECT user_id, amount, status, created_at, updated_at, currency_code, razorpay_payment_id, order_description
		FROM payment_transaction
		WHERE id = $1
	`
	row := repo.db.Pool.QueryRow(context.Background(), query, id)
	err := row.Scan(&paymentTransaction.UserID, &paymentTransaction.Amount, &paymentTransaction.Status, &paymentTransaction.CreatedAt, &paymentTransaction.UpdatedAt, &paymentTransaction.CurrencyCode, &paymentTransaction.RazorpayPaymentID, &paymentTransaction.OrderDescription)
	if err != nil {
		return models.PaymentTransaction{}, err
	}
	return paymentTransaction, nil
}

// UpdatePaymentTransaction updates an existing payment transaction record.
func (repo *PaymentsRepository) UpdatePaymentTransaction(id string, paymentTransaction models.PaymentTransaction) error {
	query := `
		UPDATE payment_transaction
		SET user_id = $1, amount = $2, status = $3, created_at = $4, updated_at = $5, currency_code = $6, razorpay_payment_id = $7, order_description = $8
		WHERE id = $9
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, paymentTransaction.UserID, paymentTransaction.Amount, paymentTransaction.Status, paymentTransaction.CreatedAt, paymentTransaction.UpdatedAt, paymentTransaction.CurrencyCode, paymentTransaction.RazorpayPaymentID, paymentTransaction.OrderDescription, id)
	if err != nil {
		return err
	}
	return nil
}

// DeletePaymentTransaction deletes a payment transaction record by ID.
func (repo *PaymentsRepository) DeletePaymentTransaction(id string) error {
	query := `
		DELETE FROM payment_transaction
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
