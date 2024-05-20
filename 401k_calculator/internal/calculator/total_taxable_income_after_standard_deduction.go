package calculator

type TotalTaxableIncomeAfterStandardDeductionCalculation Calculation

type TotalTaxableIncomeAfterStandardDeduction struct {
	TotalTaxableIncomeCalculation
	StandardDeductionCalculation
}

func NewTotalTaxableIncomeAfterStandardDeduction() TotalTaxableIncomeAfterStandardDeduction {
	return TotalTaxableIncomeAfterStandardDeduction{
		TotalTaxableIncomeCalculation: NewTotalTaxableIncome(),
		StandardDeductionCalculation:  NewStandardDeduction(),
	}
}

func (c TotalTaxableIncomeAfterStandardDeduction) CalculateTraditional(model *Model) float64 {
	totalTaxableIncome := c.TotalTaxableIncomeCalculation.CalculateTraditional(model)
	standardDeduction := c.StandardDeductionCalculation.CalculateTraditional(model)

	return totalTaxableIncome - standardDeduction
}

func (c TotalTaxableIncomeAfterStandardDeduction) CalculateTraditionalRetirement(model *Model) float64 {
	totalTaxableIncome := c.TotalTaxableIncomeCalculation.CalculateTraditionalRetirement(model)
	standardDeduction := c.StandardDeductionCalculation.CalculateTraditionalRetirement(model)

	return totalTaxableIncome - standardDeduction
}

func (c TotalTaxableIncomeAfterStandardDeduction) CalculateRoth(model *Model) float64 {
	totalTaxableIncome := c.TotalTaxableIncomeCalculation.CalculateRoth(model)
	standardDeduction := c.StandardDeductionCalculation.CalculateRoth(model)

	return totalTaxableIncome - standardDeduction
}

func (c TotalTaxableIncomeAfterStandardDeduction) CalculateRothRetirement(model *Model) float64 {
	totalTaxableIncome := c.TotalTaxableIncomeCalculation.CalculateRothRetirement(model)
	standardDeduction := c.StandardDeductionCalculation.CalculateRothRetirement(model)

	return totalTaxableIncome - standardDeduction
}
