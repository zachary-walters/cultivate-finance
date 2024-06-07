package calculator

type IncomeAfterStandardDeductionCalculation Calculation

type IncomeAfterStandardDeduction struct {
	StandardDeductionCalculation
	TotalTaxableIncomeCalculation
}

func NewIncomeAfterStandardDeduction() IncomeAfterStandardDeduction {
	return IncomeAfterStandardDeduction{
		StandardDeductionCalculation:  NewStandardDeduction(),
		TotalTaxableIncomeCalculation: NewTotalTaxableIncome(),
	}
}

func (c IncomeAfterStandardDeduction) CalculateTraditional(model *Model) float64 {
	standardDeduction := c.StandardDeductionCalculation.CalculateTraditional(model)
	currentAnnualIncome := model.Input.CurrentAnnualIncome

	return currentAnnualIncome - standardDeduction
}

func (c IncomeAfterStandardDeduction) CalculateTraditionalRetirement(model *Model) float64 {
	standardDeduction := c.StandardDeductionCalculation.CalculateTraditionalRetirement(model)
	combinedRetirementIncome := c.TotalTaxableIncomeCalculation.CalculateTraditionalRetirement(model)

	return coalesce(combinedRetirementIncome - standardDeduction)
}

func (c IncomeAfterStandardDeduction) CalculateRoth(model *Model) float64 {
	standardDeduction := c.StandardDeductionCalculation.CalculateRoth(model)
	currentAnnualIncome := model.Input.CurrentAnnualIncome

	return currentAnnualIncome - standardDeduction
}

func (c IncomeAfterStandardDeduction) CalculateRothRetirement(model *Model) float64 {
	standardDeduction := c.StandardDeductionCalculation.CalculateRothRetirement(model)
	combinedRetirementIncome := c.TotalTaxableIncomeCalculation.CalculateRothRetirement(model)

	return coalesce(combinedRetirementIncome - standardDeduction)
}
