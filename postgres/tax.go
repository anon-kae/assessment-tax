package postgres

import "github.com/anon-kae/assessment-tax/models"

func (p *Postgres) FindAllTaxRules() ([]models.TaxationRule, error) {
	rows, err := p.Db.Query("SELECT * FROM taxation_rules")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var taxation_rules []models.TaxationRule
	for rows.Next() {
		var tr models.TaxationRule
		err := rows.Scan(&tr.ID, &tr.RuleName, &tr.MaxIncome, &tr.MinIncome, &tr.TaxRate, &tr.CreatedAt, &tr.UpdatedAt)
		if err != nil {
			return nil, err
		}
		taxation_rules = append(taxation_rules, models.TaxationRule{
			ID:        tr.ID,
			RuleName:  tr.RuleName,
			MaxIncome: tr.MaxIncome,
			MinIncome: tr.MinIncome,
			TaxRate:   tr.TaxRate,
			CreatedAt: tr.CreatedAt,
			UpdatedAt: tr.UpdatedAt,
		})
	}
	return taxation_rules, nil
}
