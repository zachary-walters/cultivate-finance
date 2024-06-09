package calculator

type TopTierTaxRateCalculation Calculation

type TopTierTaxRate struct {
	TaxOnTraditionalIRAWithdrawalCalculation
}

func NewTopTierTaxRate() TopTierTaxRate {
	return TopTierTaxRate{
		TaxOnTraditionalIRAWithdrawalCalculation: NewTaxOnTraditionalIRAWithdrawal(),
	}
}

func (c TopTierTaxRate) CalculateTraditional(model *Model) float64 {
	return 0
}

func (c TopTierTaxRate) CalculateTraditionalRetirement(model *Model) float64 {
	taxOnTraditionalIRAWithdrawal := c.TaxOnTraditionalIRAWithdrawalCalculation.CalculateTraditional(model)

	if model.Input.YearlyWithdrawal == 0 {
		return 0.0
	}

	return taxOnTraditionalIRAWithdrawal / model.Input.YearlyWithdrawal
}

func (c TopTierTaxRate) CalculateRoth(model *Model) float64 {
	return 0
}

func (c TopTierTaxRate) CalculateRothRetirement(model *Model) float64 {
	return 0
}
