package calculator

type TotalTaxesOwedAfterStandardDeductionSingleCalculation Calculation

type TotalTaxesOwedAfterStandardDeductionSingle struct {
	TaxesOwedPerBracketAfterStandardDeductionSingleCalculation
}

func NewTotalTaxesOwedAfterStandardDeductionSingle() TotalTaxesOwedAfterStandardDeductionSingle {
	return TotalTaxesOwedAfterStandardDeductionSingle{
		TaxesOwedPerBracketAfterStandardDeductionSingleCalculation: NewTaxesOwedPerBracketAfterStandardDeductionSingle(),
	}
}

func (c TotalTaxesOwedAfterStandardDeductionSingle) Calculate(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionSingle := c.TaxesOwedPerBracketAfterStandardDeductionSingleCalculation.Calculate(model)

	totalTaxesOwedAfterStandardDeductionSingle := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionSingle {
		totalTaxesOwedAfterStandardDeductionSingle += value
	}

	return totalTaxesOwedAfterStandardDeductionSingle
}

func (c TotalTaxesOwedAfterStandardDeductionSingle) CalculateRetirement(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionSingle := c.TaxesOwedPerBracketAfterStandardDeductionSingleCalculation.CalculateRetirement(model)

	totalTaxesOwedAfterStandardDeductionSingle := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionSingle {
		totalTaxesOwedAfterStandardDeductionSingle += value
	}

	return totalTaxesOwedAfterStandardDeductionSingle
}
