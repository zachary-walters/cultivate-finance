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

func (c TaxesOwedPerBracketAfterStandardDeductionSingle) Calculate(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionSingleCalculation.Calculate(model)

	taxesOwedPerBracketAfterStandardDeductionSingle := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDeductionSingle[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionSingle
}

func (c TaxesOwedPerBracketAfterStandardDeductionSingle) CalculateRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionSingle := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDeductionSingle[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionSingle
}
