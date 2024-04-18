package calculator

type IncomePerBracketAfterStandardDeductionCalculation interface {
	Calculate(Model, []TaxRate) []float64
	CalculateRetirement(Model, []TaxRate) []float64
}

type IncomePerBracketAfterStandardDeduction struct {
	IncomeAfterStandardDeductionCalculation
	IncomePerBracketCalculation
}

func NewIncomePerBracketAfterStandardDeduction() IncomePerBracketAfterStandardDeduction {
	return IncomePerBracketAfterStandardDeduction{
		IncomeAfterStandardDeductionCalculation: NewIncomeAfterStandardDeduction(),
		IncomePerBracketCalculation:             NewIncomePerBracket(),
	}
}

func (c IncomePerBracketAfterStandardDeduction) Calculate(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeduction := c.IncomeAfterStandardDeductionCalculation.Calculate(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.IncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeduction))
	}

	return values
}

func (c IncomePerBracketAfterStandardDeduction) CalculateRetirement(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeduction := c.IncomeAfterStandardDeductionCalculation.CalculateRetirement(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.IncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeduction))
	}

	return values
}
