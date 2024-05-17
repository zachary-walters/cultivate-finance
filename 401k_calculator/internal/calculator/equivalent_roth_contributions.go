package calculator

type EquivalentRothContributionsCalculation Calculation

type EquivalentRothContributions struct {
	AnnualTaxSavingsWithContributionCalculation
}

func NewEquivalentRothContributions() EquivalentRothContributions {
	return EquivalentRothContributions{
		AnnualTaxSavingsWithContributionCalculation: NewAnnualTaxSavingsWithContribution(),
	}
}

func (c EquivalentRothContributions) CalculateTraditional(model Model) float64 {
	annualTaxSavingsWithContribution := c.AnnualTaxSavingsWithContributionCalculation.CalculateTraditional(model)
	annualContributionsPreTax := model.Input.AnnualContributionsPreTax

	return annualContributionsPreTax - annualTaxSavingsWithContribution
}

func (c EquivalentRothContributions) CalculateTraditionalRetirement(model Model) float64 {
	return c.CalculateTraditional(model)
}

func (c EquivalentRothContributions) CalculateRoth(model Model) float64 {
	return c.CalculateTraditional(model)
}

func (c EquivalentRothContributions) CalculateRothRetirement(model Model) float64 {
	return c.CalculateTraditional(model)
}
