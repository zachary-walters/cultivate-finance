package calculator

type AnnualTaxSavingsWithContributionCalculation Calculation

type AnnualTaxSavingsWithContribution struct {
	TotalTaxesOwedAfterStandardDeductionCalculation
	TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation
}

func NewAnnualTaxSavingsWithContribution() AnnualTaxSavingsWithContribution {
	return AnnualTaxSavingsWithContribution{
		TotalTaxesOwedAfterStandardDeductionCalculation:                 NewTotalTaxesOwedAfterStandardDeduction(),
		TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation: NewTotalTaxesOwedAfterStandardDeductionAndContributions(),
	}
}

func (c AnnualTaxSavingsWithContribution) Calculate(model Model) float64 {
	totalTaxesOwedAfterStandardDeduction := c.TotalTaxesOwedAfterStandardDeductionCalculation.Calculate(model)
	totalTaxesOwedAfterStandardDeductionAndContributions := c.TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation.Calculate(model)

	return totalTaxesOwedAfterStandardDeduction - totalTaxesOwedAfterStandardDeductionAndContributions
}

func (c AnnualTaxSavingsWithContribution) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
