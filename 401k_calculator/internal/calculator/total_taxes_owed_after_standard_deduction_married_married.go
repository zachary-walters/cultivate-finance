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

func (c TotalTaxesOwedAfterStandardDeductionMarriedSeperate) Calculate(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := c.TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation.Calculate(model)

	totalTaxesOwedAfterStandardDeductionMarriedSeperate := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedSeperate {
		totalTaxesOwedAfterStandardDeductionMarriedSeperate += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedSeperate
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedSeperate) CalculateRetirement(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := c.TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation.CalculateRetirement(model)

	totalTaxesOwedAfterStandardDeductionMarriedSeperate := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedSeperate {
		totalTaxesOwedAfterStandardDeductionMarriedSeperate += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedSeperate
}
