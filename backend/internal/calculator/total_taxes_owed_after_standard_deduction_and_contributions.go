package calculator

type TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation Calculation

type TotalTaxesOwedAfterStandardDeductionAndContributions struct {
	TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation
}

func NewTotalTaxesOwedAfterStandardDeductionAndContributions() TotalTaxesOwedAfterStandardDeductionAndContributions {
	return TotalTaxesOwedAfterStandardDeductionAndContributions{
		TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation: NewTotalTaxesOwedAfterStandardDeductionAndContributionsSingle(),
	}
}

func (c TotalTaxesOwedAfterStandardDeductionAndContributions) Calculate(model Model) float64 {
	totalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := c.TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation.Calculate(model)

	switch model.Input.CurrentFilingStatus {
	case "single":
		return totalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle
	default:
		return 0.0
	}
}

func (c TotalTaxesOwedAfterStandardDeductionAndContributions) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
