package models

import (
	"time"
)

type TaxConfiguration struct {
	ID            int       `json:"id" postgres:"id"`
	ConditionName string    `json:"condition_name" postgres:"condition_name"`
	Amount        float64   `json:"amount" postgres:"amount"`
	CreatedAt     time.Time `json:"created_at" postgres:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" postgres:"updated_at"`
}
