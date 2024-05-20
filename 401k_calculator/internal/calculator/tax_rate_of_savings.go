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

func (c TaxRateOfSavings) CalculateTraditional(model *Model) float64 {
	annualTaxSavingsWithContribution := c.AnnualTaxSavingsWithContributionCalculation.CalculateTraditional(model)
	annualContributions := model.Input.AnnualContributionsPreTax

	return annualTaxSavingsWithContribution / annualContributions
}

func (c TaxRateOfSavings) CalculateTraditionalRetirement(model *Model) float64 {
	return c.CalculateTraditional(model)
}

func (c TaxRateOfSavings) CalculateRoth(model *Model) float64 {
	return c.CalculateTraditional(model)
}

func (c TaxRateOfSavings) CalculateRothRetirement(model *Model) float64 {
	return c.CalculateTraditional(model)
}
