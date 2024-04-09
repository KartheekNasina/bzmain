package models

type AuthProviderDTO struct {
	Provider string `json:"provider" db:"provider"`
}
