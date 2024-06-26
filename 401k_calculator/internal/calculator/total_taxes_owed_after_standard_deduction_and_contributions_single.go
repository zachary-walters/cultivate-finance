package calculator

type TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation Calculation

type TotalTaxesOwedAfterStandardDeductionAndContributionsSingle struct {
	TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation
}

func NewTotalTaxesOwedAfterStandardDeductionAndContributionsSingle() TotalTaxesOwedAfterStandardDeductionAndContributionsSingle {
	return TotalTaxesOwedAfterStandardDeductionAndContributionsSingle{
		TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation: NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle(),
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
	return 0.0
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
	return 0.0
}
