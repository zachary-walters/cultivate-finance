package calculator

type TotalTaxesOwedAfterStandardDeductionCalculation Calculation

type TotalTaxesOwedAfterStandardDeduction struct {
	TotalTaxesOwedAfterStandardDeductionSingleCalculation
}

func NewTotalTaxesOwedAfterStandardDeduction() TotalTaxesOwedAfterStandardDeduction {
	return TotalTaxesOwedAfterStandardDeduction{
		TotalTaxesOwedAfterStandardDeductionSingleCalculation: NewTotalTaxesOwedAfterStandardDeductionSingle(),
	}
}

func (c TotalTaxesOwedAfterStandardDeduction) Calculate(model Model) float64 {
	totalTaxesOwedPerBracketAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.Calculate(model)

	switch model.Input.CurrentFilingStatus {
	case "single":
		return totalTaxesOwedPerBracketAfterStandardDeductionSingle
	default:
		return 0
	}
}

func (c TotalTaxesOwedAfterStandardDeduction) CalculateRetirement(model Model) float64 {
	totalTaxesOwedPerBracketAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.CalculateRetirement(model)

	switch model.Input.RetirementFilingStatus {
	case "single":
		return totalTaxesOwedPerBracketAfterStandardDeductionSingle
	default:
		return 0
	}
}
