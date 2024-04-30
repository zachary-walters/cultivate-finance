package calculator

type TotalInterestCalculation Calculation

type TotalInterest struct {
	TotalDisbursementsAfterTaxCalculation
	TotalContributionsCalculation
}

func NewTotalInterest() TotalInterest {
	return TotalInterest{
		TotalDisbursementsAfterTaxCalculation: NewTotalDisbursementsAfterTax(),
		TotalContributionsCalculation:         NewTotalContributions(),
	}
}

func (c TotalInterest) Calculate(model Model) float64 {
	totalDisbursements := c.TotalDisbursementsAfterTaxCalculation.Calculate(model)
	totalContributions := c.TotalContributionsCalculation.Calculate(model)

	return totalDisbursements - totalContributions
}

func (c TotalInterest) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
