package calculator

type TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation SequenceCalculation

type TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate struct {
	IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation
}

func NewTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate() TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate {
	return TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate{
		IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation: NewIncomePerBracketAfterStandardDeductionMarriedSeparate(),
	}
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate) CalculateTraditional(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateTraditional(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedSeparate := make([]float64, len(model.MarriedSeparateTaxRates))

	for idx, taxRate := range model.MarriedSeparateTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedSeparate[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedSeparate
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate) CalculateTraditionalRetirement(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateTraditionalRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedSeparate := make([]float64, len(model.MarriedSeparateTaxRates))

	for idx, taxRate := range model.MarriedSeparateTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedSeparate[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedSeparate
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate) CalculateRoth(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateRoth(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedSeparate := make([]float64, len(model.MarriedSeparateTaxRates))

	for idx, taxRate := range model.MarriedSeparateTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedSeparate[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedSeparate
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate) CalculateRothRetirement(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateRothRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedSeparate := make([]float64, len(model.MarriedSeparateTaxRates))

	for idx, taxRate := range model.MarriedSeparateTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedSeparate[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedSeparate
}
