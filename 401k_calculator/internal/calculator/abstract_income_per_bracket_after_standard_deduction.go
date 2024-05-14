package calculator

type AbstractIncomePerBracketAfterStandardDeductionCalculation interface {
	CalculateTraditional(Model, []TaxRate) []float64
	CalculateTraditionalRetirement(Model, []TaxRate) []float64
	CalculateRoth(Model, []TaxRate) []float64
	CalculateRothRetirement(Model, []TaxRate) []float64
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

func (c AbstractIncomePerBracketAfterStandardDeduction) CalculateTraditional(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeduction := c.IncomeAfterStandardDeductionCalculation.CalculateTraditional(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeduction))
	}

	return values
}

func (c AbstractIncomePerBracketAfterStandardDeduction) CalculateTraditionalRetirement(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeduction := c.IncomeAfterStandardDeductionCalculation.CalculateTraditionalRetirement(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeduction))
	}

	return values
}

func (c AbstractIncomePerBracketAfterStandardDeduction) CalculateRoth(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeduction := c.IncomeAfterStandardDeductionCalculation.CalculateRoth(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeduction))
	}

	return values
}

func (c AbstractIncomePerBracketAfterStandardDeduction) CalculateRothRetirement(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeduction := c.IncomeAfterStandardDeductionCalculation.CalculateRothRetirement(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeduction))
	}

	return values
}
