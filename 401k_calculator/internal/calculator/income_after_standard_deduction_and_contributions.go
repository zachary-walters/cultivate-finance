package calculator

type IncomeAfterStandardDeductionAndContributionsCalculation Calculation

type IncomeAfterStandardDeductionAndContributions struct {
	IncomeAfterStandardDeductionCalculation
}

func NewIncomeAfterStandardDeductionAndContributions() IncomeAfterStandardDeductionAndContributions {
	return IncomeAfterStandardDeductionAndContributions{
		IncomeAfterStandardDeductionCalculation: NewIncomeAfterStandardDeduction(),
	}
}

func (c IncomeAfterStandardDeductionAndContributions) Calculate(model Model) float64 {
	incomeAfterStandardDeduction := c.IncomeAfterStandardDeductionCalculation.Calculate(model)
	currentAnnualContribution := model.Input.AnnualContributionsPreTax

	return incomeAfterStandardDeduction - currentAnnualContribution
}

func (c IncomeAfterStandardDeductionAndContributions) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
