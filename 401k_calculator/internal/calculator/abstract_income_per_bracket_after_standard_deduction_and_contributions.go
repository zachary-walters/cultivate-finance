package calculator

type AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation interface {
	CalculateTraditional(*Model, []TaxRate) []float64
	CalculateTraditionalRetirement(*Model, []TaxRate) []float64
	CalculateRoth(*Model, []TaxRate) []float64
	CalculateRothRetirement(*Model, []TaxRate) []float64
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

func (c AbstractIncomePerBracketAfterStandardDeductionAndContributions) CalculateTraditional(model *Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeductionAndContributions := c.IncomeAfterStandardDeductionAndContributionsCalculation.CalculateTraditional(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeductionAndContributions))
	}

	return values
}

func (c AbstractIncomePerBracketAfterStandardDeductionAndContributions) CalculateTraditionalRetirement(model *Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeductionAndContributions := c.IncomeAfterStandardDeductionAndContributionsCalculation.CalculateTraditionalRetirement(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeductionAndContributions))
	}

	return values
}

func (c AbstractIncomePerBracketAfterStandardDeductionAndContributions) CalculateRoth(model *Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeductionAndContributions := c.IncomeAfterStandardDeductionAndContributionsCalculation.CalculateRoth(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeductionAndContributions))
	}

	return values
}

func (c AbstractIncomePerBracketAfterStandardDeductionAndContributions) CalculateRothRetirement(model *Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeductionAndContributions := c.IncomeAfterStandardDeductionAndContributionsCalculation.CalculateRothRetirement(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.AbstractIncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeductionAndContributions))
	}

	return values
}
