package calculator

type TotalInterestAfterTaxCalculation Calculation

type TotalInterestAfterTax struct {
	TotalDisbursementsAfterTaxCalculation
	TotalContributionsCalculation
}

func NewTotalInterestAfterTax() TotalInterestAfterTax {
	return TotalInterestAfterTax{
		TotalDisbursementsAfterTaxCalculation: NewTotalDisbursementsAfterTax(),
		TotalContributionsCalculation:         NewTotalContributions(),
	}
}

func (c TotalInterestAfterTax) Calculate(model Model) float64 {
	totalDisbursementsAfterTax := c.TotalDisbursementsAfterTaxCalculation.Calculate(model)
	totalContributions := c.TotalContributionsCalculation.Calculate(model)

	return totalDisbursementsAfterTax - totalContributions
}

func (c TotalInterestAfterTax) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
