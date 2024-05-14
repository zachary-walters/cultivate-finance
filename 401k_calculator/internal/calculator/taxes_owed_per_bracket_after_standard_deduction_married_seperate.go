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

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedSeperate) CalculateTraditional(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation.CalculateTraditional(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := make([]float64, len(model.MarriedSeperateTaxRates))

	for idx, taxRate := range model.MarriedSeperateTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedSeperate[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedSeperate
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedSeperate) CalculateTraditionalRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation.CalculateTraditionalRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := make([]float64, len(model.MarriedSeperateTaxRates))

	for idx, taxRate := range model.MarriedSeperateTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedSeperate[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedSeperate
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedSeperate) CalculateRoth(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation.CalculateRoth(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := make([]float64, len(model.MarriedSeperateTaxRates))

	for idx, taxRate := range model.MarriedSeperateTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedSeperate[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedSeperate
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedSeperate) CalculateRothRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation.CalculateRothRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedSeperate := make([]float64, len(model.MarriedSeperateTaxRates))

	for idx, taxRate := range model.MarriedSeperateTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedSeperate[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedSeperate
}
