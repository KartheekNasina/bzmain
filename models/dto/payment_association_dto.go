package models

type PaymentAssociationDTO struct {
	PaymentTransactionId string `json:"payment_transaction_id" db:"payment_transaction_id"`
	EntityType string `json:"entity_type" db:"entity_type"`
	EntityId string `json:"entity_id" db:"entity_id"`
}
