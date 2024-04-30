package calculator

type TaxRateOfSavingsCalculation Calculation

type TaxRateOfSavings struct {
	AnnualTaxSavingsWithContributionCalculation
}

func NewTaxRateOfSavings() TaxRateOfSavings {
	return TaxRateOfSavings{
		AnnualTaxSavingsWithContributionCalculation: NewAnnualTaxSavingsWithContribution(),
	}
}

func (c TaxRateOfSavings) Calculate(model Model) float64 {
	annualTaxSavingsWithContribution := c.AnnualTaxSavingsWithContributionCalculation.Calculate(model)
	annualContributions := model.Input.AnnualContributionsPreTax

	return annualTaxSavingsWithContribution / annualContributions
}

func (c TaxRateOfSavings) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
