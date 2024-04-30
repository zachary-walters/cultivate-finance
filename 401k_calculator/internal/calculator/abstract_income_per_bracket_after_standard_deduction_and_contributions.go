package calculator

type AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation interface {
	Calculate(Model, []TaxRate) []float64
	CalculateRetirement(Model, []TaxRate) []float64
}

type AbstractIncomePerBracketAfterStandardDeductionAndContributions struct {
	AbstractIncomePerBracketCalculation
	IncomeAfterStandardDeductionAndContributionsCalculation
}

func NewAbstractIncomePerBracketAfterStandardDeductionAndContributions() AbstractIncomePerBracketAfterStandardDeductionAndContributions {
	return AbstractIncomePerBracketAfterStandardDeductionAndContributions{
		AbstractIncomePerBracketCalculation:                     NewAbstractIncomePerBracket(),
		IncomeAfterStandardDeductionAndContributionsCalculation: NewIncomeAfterStandardDeductionAndContributions(),
	}
}

func (c AbstractIncomePerBracketAfterStandardDeductionAndContributions) Calculate(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeductionAndContributions := c.IncomeAfterStandardDeductionAndContributionsCalculation.Calculate(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeductionAndContributions))
	}

	return values
}

func (c AbstractIncomePerBracketAfterStandardDeductionAndContributions) CalculateRetirement(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeductionAndContributions := c.IncomeAfterStandardDeductionAndContributionsCalculation.CalculateRetirement(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeductionAndContributions))
	}

	return values
}
