package calculator

type TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation SequenceCalculation

type TaxesOwedPerBracketAfterStandardDeductionMarriedSeperate struct {
	IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation
}

func NewTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate() TaxesOwedPerBracketAfterStandardDeductionMarriedSeperate {
	return TaxesOwedPerBracketAfterStandardDeductionMarriedSeperate{
		IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation: NewIncomePerBracketAfterStandardDeductionMarriedSeperate(),
	}
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedSeperate) Calculate(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation.Calculate(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := make([]float64, len(model.MarriedSeperateTaxRates))

	for idx, taxRate := range model.MarriedSeperateTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedSeperate[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedSeperate
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedSeperate) CalculateRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation.CalculateRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := make([]float64, len(model.MarriedSeperateTaxRates))

	for idx, taxRate := range model.MarriedSeperateTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedSeperate[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedSeperate
}
