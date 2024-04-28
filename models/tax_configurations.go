package models

import (
	"time"
)

type TaxConfiguration struct {
	Id            int       `json:"id"`
	ConditionName string    `json:"condition_name"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
