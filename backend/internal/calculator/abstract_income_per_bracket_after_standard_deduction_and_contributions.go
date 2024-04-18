package calculator

type IncomePerBracketAfterStandardDeductionAndContributionsCalculation interface {
	Calculate(Model, []TaxRate) []float64
	CalculateRetirement(Model, []TaxRate) []float64
}

type IncomePerBracketAfterStandardDeductionAndContributions struct {
	IncomePerBracketCalculation
	IncomeAfterStandardDeductionAndContributionsCalculation
}

func NewIncomePerBracketAfterStandardDeductionAndContributions() IncomePerBracketAfterStandardDeductionAndContributions {
	return IncomePerBracketAfterStandardDeductionAndContributions{
		IncomePerBracketCalculation:                             NewIncomePerBracket(),
		IncomeAfterStandardDeductionAndContributionsCalculation: NewIncomeAfterStandardDeductionAndContributions(),
	}
}

func (c IncomePerBracketAfterStandardDeductionAndContributions) Calculate(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeductionAndContributions := c.IncomeAfterStandardDeductionAndContributionsCalculation.Calculate(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.IncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeductionAndContributions))
	}

	return values
}

func (c IncomePerBracketAfterStandardDeductionAndContributions) CalculateRetirement(model Model, taxRates []TaxRate) []float64 {
	incomeAfterStandardDeductionAndContributions := c.IncomeAfterStandardDeductionAndContributionsCalculation.CalculateRetirement(model)

	values := []float64{}

	for idx := range taxRates {
		values = append(values, c.IncomePerBracketCalculation.Calculate(taxRates, idx, incomeAfterStandardDeductionAndContributions))
	}

	return values
}
