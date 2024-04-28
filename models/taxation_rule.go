package models

import (
	"time"
)

type TaxationRule struct {
	ID        int       `json:"id"`
	RuleName  string    `json:"rule_name"`
	MaxIncome float64   `json:"max_income"`
	MinIncome float64   `json:"min_income"`
	TaxRate   float64   `json:"tax_rate"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
