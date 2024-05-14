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

func (c AnnualTaxSavingsWithContribution) CalculateTraditional(model Model) float64 {
	totalTaxesOwedAfterStandardDeduction := c.TotalTaxesOwedAfterStandardDeductionCalculation.CalculateTraditional(model)
	totalTaxesOwedAfterStandardDeductionAndContributions := c.TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation.CalculateTraditional(model)

	return totalTaxesOwedAfterStandardDeduction - totalTaxesOwedAfterStandardDeductionAndContributions
}

func (c AnnualTaxSavingsWithContribution) CalculateTraditionalRetirement(model Model) float64 {
	return c.CalculateTraditional(model)
}

func (c AnnualTaxSavingsWithContribution) CalculateRoth(model Model) float64 {
	return 0
}

func (c AnnualTaxSavingsWithContribution) CalculateRothRetirement(model Model) float64 {
	return c.CalculateRoth(model)
}
