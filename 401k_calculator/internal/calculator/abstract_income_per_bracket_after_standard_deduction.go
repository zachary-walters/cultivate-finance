package calculator

type AbstractIncomePerBracketAfterStandardDeductionCalculation interface {
	Calculate(Model, []TaxRate) []float64
	CalculateRetirement(Model, []TaxRate) []float64
}

type AbstractIncomePerBracketAfterStandardDeduction struct {
	IncomeAfterStandardDeductionCalculation
	AbstractIncomePerBracketCalculation
}

func NewAbstractIncomePerBracketAfterStandardDeduction() AbstractIncomePerBracketAfterStandardDeduction {
	return AbstractIncomePerBracketAfterStandardDeduction{
		IncomeAfterStandardDeductionCalculation: NewIncomeAfterStandardDeduction(),
		AbstractIncomePerBracketCalculation:     NewAbstractIncomePerBracket(),
	}
}

func (c AbstractIncomePerBracketAfterStandardDeduction) Calculate(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeduction := c.IncomeAfterStandardDeductionCalculation.Calculate(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeduction))
	}

	return values
}

func (c AbstractIncomePerBracketAfterStandardDeduction) CalculateRetirement(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeduction := c.IncomeAfterStandardDeductionCalculation.CalculateRetirement(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeduction))
	}

	return values
}
