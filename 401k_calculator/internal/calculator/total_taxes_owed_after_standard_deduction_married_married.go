package calculator

type TotalTaxesOwedAfterStandardDeductionMarriedSeperateCalculation Calculation

type TotalTaxesOwedAfterStandardDeductionMarriedSeperate struct {
	TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation
}

func NewTotalTaxesOwedAfterStandardDeductionMarriedSeperate() TotalTaxesOwedAfterStandardDeductionMarriedSeperate {
	return TotalTaxesOwedAfterStandardDeductionMarriedSeperate{
		TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation: NewTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate(),
	}
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedSeperate) CalculateTraditional(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := c.TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation.CalculateTraditional(model)

	totalTaxesOwedAfterStandardDeductionMarriedSeperate := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedSeperate {
		totalTaxesOwedAfterStandardDeductionMarriedSeperate += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedSeperate
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedSeperate) CalculateTraditionalRetirement(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := c.TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation.CalculateTraditionalRetirement(model)

	totalTaxesOwedAfterStandardDeductionMarriedSeperate := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedSeperate {
		totalTaxesOwedAfterStandardDeductionMarriedSeperate += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedSeperate
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedSeperate) CalculateRoth(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := c.TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation.CalculateRoth(model)

	totalTaxesOwedAfterStandardDeductionMarriedSeperate := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedSeperate {
		totalTaxesOwedAfterStandardDeductionMarriedSeperate += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedSeperate
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedSeperate) CalculateRothRetirement(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := c.TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation.CalculateRothRetirement(model)

	totalTaxesOwedAfterStandardDeductionMarriedSeperate := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedSeperate {
		totalTaxesOwedAfterStandardDeductionMarriedSeperate += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedSeperate
}
