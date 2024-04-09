package models

type PaymentAssociation struct {
	PaymentTransactionID string                `db:"payment_transaction_id"`
	EntityType           EntityAssociationType `db:"entity_type"`
	EntityID             string                `db:"entity_id"`
}
