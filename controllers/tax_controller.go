package controllers

import (
	"fmt"

	"github.com/anon-kae/assessment-tax/errortype"
	"github.com/anon-kae/assessment-tax/helper"
	"github.com/anon-kae/assessment-tax/models"
	"github.com/labstack/echo/v4"
)

type Store interface {
	FindAllTaxRules() ([]models.TaxationRule, error)
}

type TaxController struct {
	store Store
}

func New(db Store) *TaxController {
	return &TaxController{store: db}
}

const (
	AllowanceTypeDonation string = "donation"
	AllowanceTypeKReceipt string = "k-receipt"
)

type TaxResponse struct {
	Tax float64 `json:"tax"`
}

type Allowance struct {
	AllowanceType string  `json:"allowanceType" example:"donation"`
	Amount        float64 `json:"amount" example:"0.0"`
}

type TaxPayload struct {
	TotalIncome float64     `json:"totalIncome" example:"500000.0" validate:"required"`
	Wht         float64     `json:"wht"`
	Allowances  []Allowance `json:"allowances" validate:"required"`
}

func (tc *TaxController) TaxCalculate(c echo.Context) error {
	var income TaxPayload
	var err error
	if err = c.Bind(&income); err != nil {
		fmt.Println(income)
		return err
	}

	if err = c.Validate(&income); err != nil {
		return err
	}

	taxRules, err := tc.store.FindAllTaxRules()
	if err != nil {
		return err
	}

	taxableIncome, err := tc.calculateTaxableIncome(income)
	if err != nil {
		return err
	}

	tax := calculateTax(taxableIncome, taxRules) - income.Wht
	

	return helper.SuccessHandler(c, TaxResponse{Tax: tax})
}

func (c *TaxController) calculateTaxableIncome(payload TaxPayload) (float64, error) {
	taxableIncome := payload.TotalIncome - 60000.0

	for _, allowance := range payload.Allowances {
		switch allowance.AllowanceType {
		case AllowanceTypeDonation:
			taxableIncome -= allowance.Amount
		case AllowanceTypeKReceipt:
			// TODO: Implement K-Receipt logic
		default:
			return 0, errortype.ValidationError{Message: "Invalid type"}
		}
	}

	return taxableIncome, nil
}

func calculateTax(income float64, taxRules []models.TaxationRule) float64 {
	var tax float64

	for _, taxRule := range taxRules {
		if income <= 0 {
			break
		}

		taxableAmount := income
		if taxableAmount > taxRule.MaxIncome {
			taxableAmount = taxRule.MaxIncome
		}

		tax += taxableAmount * (taxRule.TaxRate / 100.0)
		income -= taxableAmount
	}

	return tax
}

func (tc *TaxController) RegisterRoutes(e *echo.Echo) {
	e.POST("/tax/calculations", tc.TaxCalculate)
}
