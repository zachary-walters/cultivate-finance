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

func (c IncomeAfterStandardDeductionAndContributions) CalculateTraditional(model *Model) float64 {
	incomeAfterStandardDeduction := c.IncomeAfterStandardDeductionCalculation.CalculateTraditional(model)
	currentAnnualContribution := model.Input.AnnualContributionsPreTax

	return incomeAfterStandardDeduction - currentAnnualContribution
}

func (c IncomeAfterStandardDeductionAndContributions) CalculateTraditionalRetirement(model *Model) float64 {
	return c.CalculateTraditional(model)
}

func (c IncomeAfterStandardDeductionAndContributions) CalculateRoth(model *Model) float64 {
	return c.CalculateTraditional(model)
}

func (c IncomeAfterStandardDeductionAndContributions) CalculateRothRetirement(model *Model) float64 {
	return c.CalculateTraditional(model)
}
