package calculator

type IncomeAfterStandardDeductionCalculation Calculation

type IncomeAfterStandardDeduction struct {
	StandardDeductionCalculation
	TotalTaxableIncomeTraditionalCalculation
}

func NewIncomeAfterStandardDeduction() IncomeAfterStandardDeduction {
	return IncomeAfterStandardDeduction{
		StandardDeductionCalculation:                   NewStandardDeduction(),
		TotalTaxableIncomeTraditionalCalculation:       NewTotalTaxableIncomeTraditional(),
	}
}

func (c IncomeAfterStandardDeduction) Calculate(model Model) float64 {
	standardDeduction := c.StandardDeductionCalculation.Calculate(model)
	currentAnnualIncome := model.Input.CurrentAnnualIncome

	return currentAnnualIncome - standardDeduction
}

func (c IncomeAfterStandardDeduction) CalculateRetirement(model Model) float64 {
	standardDeduction := c.StandardDeductionCalculation.CalculateRetirement(model)

	combinedRetirementIncomeTraditional := c.TotalTaxableIncomeTraditionalCalculation.Calculate(model)

	return combinedRetirementIncomeTraditional - standardDeduction
}
