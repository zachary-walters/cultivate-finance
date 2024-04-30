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

func (c EquivalentRothContributions) Calculate(model Model) float64 {
	annualTaxSavingsWithContribution := c.AnnualTaxSavingsWithContributionCalculation.Calculate(model)
	annualContributionsPreTax := model.Input.AnnualContributionsPreTax

	return annualContributionsPreTax - annualTaxSavingsWithContribution
}

func (c EquivalentRothContributions) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
