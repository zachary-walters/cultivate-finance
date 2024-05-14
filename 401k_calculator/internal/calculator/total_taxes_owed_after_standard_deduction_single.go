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

func (c TotalTaxesOwedAfterStandardDeductionSingle) CalculateTraditional(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionSingle := c.TaxesOwedPerBracketAfterStandardDeductionSingleCalculation.CalculateTraditional(model)

	totalTaxesOwedAfterStandardDeductionSingle := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionSingle {
		totalTaxesOwedAfterStandardDeductionSingle += value
	}

	return totalTaxesOwedAfterStandardDeductionSingle
}

func (c TotalTaxesOwedAfterStandardDeductionSingle) CalculateTraditionalRetirement(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionSingle := c.TaxesOwedPerBracketAfterStandardDeductionSingleCalculation.CalculateTraditionalRetirement(model)

	totalTaxesOwedAfterStandardDeductionSingle := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionSingle {
		totalTaxesOwedAfterStandardDeductionSingle += value
	}

	return totalTaxesOwedAfterStandardDeductionSingle
}

func (c TotalTaxesOwedAfterStandardDeductionSingle) CalculateRoth(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionSingle := c.TaxesOwedPerBracketAfterStandardDeductionSingleCalculation.CalculateRoth(model)

	totalTaxesOwedAfterStandardDeductionSingle := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionSingle {
		totalTaxesOwedAfterStandardDeductionSingle += value
	}

	return totalTaxesOwedAfterStandardDeductionSingle
}

func (c TotalTaxesOwedAfterStandardDeductionSingle) CalculateRothRetirement(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionSingle := c.TaxesOwedPerBracketAfterStandardDeductionSingleCalculation.CalculateRothRetirement(model)

	totalTaxesOwedAfterStandardDeductionSingle := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionSingle {
		totalTaxesOwedAfterStandardDeductionSingle += value
	}

	return totalTaxesOwedAfterStandardDeductionSingle
}
