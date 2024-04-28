package controllers

import (
	"fmt"

	_ "github.com/anon-kae/assessment-tax/errortype"
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
	taxableIncome := income.TotalIncome - 60000.0
	tax := calculateTax(taxableIncome, taxRules)

	return helper.SuccessHandler(c, TaxResponse{Tax: tax})
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
	e.GET("/tax/calculations", tc.TaxCalculate)
}
