package models

import (
	"time"
)

type AdminDTO struct {
	Id string `json:"id" db:"id"`
	Role string `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Name string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
