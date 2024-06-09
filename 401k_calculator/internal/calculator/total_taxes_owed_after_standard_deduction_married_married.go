package calculator

type TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation Calculation

type TotalTaxesOwedAfterStandardDeductionMarriedSeparate struct {
	TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation
}

func NewTotalTaxesOwedAfterStandardDeductionMarriedSeparate() TotalTaxesOwedAfterStandardDeductionMarriedSeparate {
	return TotalTaxesOwedAfterStandardDeductionMarriedSeparate{
		TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: NewTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate(),
	}
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateTraditional(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedSeparate := c.TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateTraditional(model)

	totalTaxesOwedAfterStandardDeductionMarriedSeparate := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedSeparate {
		totalTaxesOwedAfterStandardDeductionMarriedSeparate += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedSeparate
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateTraditionalRetirement(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedSeparate := c.TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateTraditionalRetirement(model)

	totalTaxesOwedAfterStandardDeductionMarriedSeparate := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedSeparate {
		totalTaxesOwedAfterStandardDeductionMarriedSeparate += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedSeparate
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateRoth(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedSeparate := c.TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateRoth(model)

	totalTaxesOwedAfterStandardDeductionMarriedSeparate := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedSeparate {
		totalTaxesOwedAfterStandardDeductionMarriedSeparate += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedSeparate
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateRothRetirement(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedSeparate := c.TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateRothRetirement(model)

	totalTaxesOwedAfterStandardDeductionMarriedSeparate := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedSeparate {
		totalTaxesOwedAfterStandardDeductionMarriedSeparate += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedSeparate
}
