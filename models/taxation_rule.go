package models

import (
	"time"
)

type TaxationRule struct {
	ID        int       `json:"id" postgres:"id"`
	RuleName  string    `json:"rule_name" postgres:"rule_name"`
	MaxIncome float64   `json:"max_income" postgres:"max_income"`
	MinIncome float64   `json:"min_income" postgres:"min_income"`
	TaxRate   float64   `json:"tax_rate" postgres:"tax_rate"`
	CreatedAt time.Time `json:"created_at" postgres:"created_at"`
	UpdatedAt time.Time `json:"updated_at" postgres:"updated_at"`
}
