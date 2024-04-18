package calculator

type NetDistributionAfterTaxesCalculation Calculation

type NetDistributionAfterTaxes struct {
	TotalTaxesOwedAfterStandardDeductionCalculation
}

func NewNetDistributionAfterTaxes() NetDistributionAfterTaxes {
	return NetDistributionAfterTaxes{
		TotalTaxesOwedAfterStandardDeductionCalculation: NewTotalTaxesOwedAfterStandardDeduction(),
	}
}

func (c NetDistributionAfterTaxes) Calculate(model Model) float64 {
	return c.CalculateRetirement(model)
}

func (c NetDistributionAfterTaxes) CalculateRetirement(model Model) float64 {
	totalTaxesOwedAfterStandardDeduction := c.TotalTaxesOwedAfterStandardDeductionCalculation.CalculateRetirement(model)
	yearlyWithdrawal := model.Input.YearlyWithdrawal

	return yearlyWithdrawal - totalTaxesOwedAfterStandardDeduction
}
