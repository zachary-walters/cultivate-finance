package calculator

type TaxesOwedPerBracketAfterStandardDeductionSingleCalculation SequenceCalculation

type TaxesOwedPerBracketAfterStandardDeductionSingle struct {
	IncomePerBracketAfterStandardDeductionSingleCalculation
}

func NewTaxesOwedPerBracketAfterStandardDeductionSingle() TaxesOwedPerBracketAfterStandardDeductionSingle {
	return TaxesOwedPerBracketAfterStandardDeductionSingle{
		IncomePerBracketAfterStandardDeductionSingleCalculation: NewIncomePerBracketAfterStandardDeductionSingle(),
	}
}

func (c TaxesOwedPerBracketAfterStandardDeductionSingle) CalculateTraditional(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateTraditional(model)

	taxesOwedPerBracketAfterStandardDeductionSingle := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDeductionSingle[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionSingle
}

func (c TaxesOwedPerBracketAfterStandardDeductionSingle) CalculateTraditionalRetirement(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateTraditionalRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionSingle := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDeductionSingle[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionSingle
}

func (c TaxesOwedPerBracketAfterStandardDeductionSingle) CalculateRoth(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateRoth(model)

	taxesOwedPerBracketAfterStandardDeductionSingle := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDeductionSingle[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionSingle
}

func (c TaxesOwedPerBracketAfterStandardDeductionSingle) CalculateRothRetirement(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateRothRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionSingle := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDeductionSingle[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionSingle
}
