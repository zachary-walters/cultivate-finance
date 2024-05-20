package calculator

type TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation Calculation

type TotalTaxesOwedAfterStandardDeductionAndContributionsSingle struct {
	TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation
	TotalTaxesOwedAfterStandardDeductionSingleCalculation
}

func NewTotalTaxesOwedAfterStandardDeductionAndContributionsSingle() TotalTaxesOwedAfterStandardDeductionAndContributionsSingle {
	return TotalTaxesOwedAfterStandardDeductionAndContributionsSingle{
		TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation: NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle(),
		TotalTaxesOwedAfterStandardDeductionSingleCalculation:                      NewTotalTaxesOwedAfterStandardDeductionSingle(),
	}
}

func (c TotalTaxesOwedAfterStandardDeductionAndContributionsSingle) CalculateTraditional(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := c.TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation.CalculateTraditional(model)

	totalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := 0.0
	for _, value := range taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle {
		totalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle += value
	}

	return totalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle
}

func (c TotalTaxesOwedAfterStandardDeductionAndContributionsSingle) CalculateTraditionalRetirement(model *Model) float64 {
	return c.CalculateTraditional(model)
}

func (c TotalTaxesOwedAfterStandardDeductionAndContributionsSingle) CalculateRoth(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := c.TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation.CalculateRoth(model)

	totalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := 0.0
	for _, value := range taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle {
		totalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle += value
	}

	return totalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle
}

func (c TotalTaxesOwedAfterStandardDeductionAndContributionsSingle) CalculateRothRetirement(model *Model) float64 {
	return c.CalculateRoth(model)
}
