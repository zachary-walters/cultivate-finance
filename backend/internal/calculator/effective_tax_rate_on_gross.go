package calculator

type EffectiveTaxRateOnGrossCalculation Calculation

type EffectiveTaxRateOnGross struct {
	TotalTaxesOwedAfterStandardDeductionCalculation
}

func NewEffectiveTaxRateOnGross() EffectiveTaxRateOnGross {
	return EffectiveTaxRateOnGross{
		TotalTaxesOwedAfterStandardDeductionCalculation: NewTotalTaxesOwedAfterStandardDeduction(),
	}
}

func (c EffectiveTaxRateOnGross) Calculate(model Model) float64 {
	return c.CalculateRetirement(model)
}

func (c EffectiveTaxRateOnGross) CalculateRetirement(model Model) float64 {
	totalTaxesOwedAfterStandardDeduction := c.TotalTaxesOwedAfterStandardDeductionCalculation.CalculateRetirement(model)
	yearlyWithdrawal := model.Input.YearlyWithdrawal

	if yearlyWithdrawal == 0.0 {
		return 0.0
	}

	return totalTaxesOwedAfterStandardDeduction / yearlyWithdrawal
}
