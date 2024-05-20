package calculator

type EffectiveTaxRateOnGrossCalculation Calculation

type EffectiveTaxRateOnGross struct {
	TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation
	TotalTaxesOwedAfterStandardDeductionCalculation
	TotalAnnualRetirementIncomeBeforeTaxCalculation
	TotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalCalculation
}

func NewEffectiveTaxRateOnGross() EffectiveTaxRateOnGross {
	return EffectiveTaxRateOnGross{
		TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation:                  NewTotalTaxesOwedAfterStandardDeductionAndContributions(),
		TotalTaxesOwedAfterStandardDeductionCalculation:                                  NewTotalTaxesOwedAfterStandardDeduction(),
		TotalAnnualRetirementIncomeBeforeTaxCalculation:                                  NewTotalAnnualRetirementIncomeBeforeTax(),
		TotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalCalculation: NewTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal(),
	}
}

func (c EffectiveTaxRateOnGross) CalculateTraditional(model *Model) float64 {
	totalTaxesOwedAfterStandardDeductionAndContributions := c.TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation.CalculateTraditional(model)
	currentAnnualIncome := model.Input.CurrentAnnualIncome

	if currentAnnualIncome == 0.0 {
		return 0.0
	}

	return totalTaxesOwedAfterStandardDeductionAndContributions / currentAnnualIncome
}

func (c EffectiveTaxRateOnGross) CalculateTraditionalRetirement(model *Model) float64 {
	totalTaxesOwedAfterStandardDeduction := c.TotalTaxesOwedAfterStandardDeductionCalculation.CalculateTraditionalRetirement(model)
	totalAnnualRetirementIncomeBeforeTax := c.TotalAnnualRetirementIncomeBeforeTaxCalculation.CalculateTraditionalRetirement(model)

	if totalAnnualRetirementIncomeBeforeTax == 0.0 {
		return 0.0
	}

	return totalTaxesOwedAfterStandardDeduction / totalAnnualRetirementIncomeBeforeTax
}

func (c EffectiveTaxRateOnGross) CalculateRoth(model *Model) float64 {
	return c.CalculateTraditional(model)
}

func (c EffectiveTaxRateOnGross) CalculateRothRetirement(model *Model) float64 {
	totalTaxesOwedAfterStandardDeduction := c.TotalTaxesOwedAfterStandardDeductionCalculation.CalculateRothRetirement(model)
	totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal := c.TotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalCalculation.CalculateRothRetirement(model)

	if totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal == 0.0 {
		return 0.0
	}

	return totalTaxesOwedAfterStandardDeduction / totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal
}
