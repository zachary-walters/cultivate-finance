package calculator

type TaxOnTraditionalIRAWithdrawalCalculation Calculation

type TaxOnTraditionalIRAWithdrawal struct {
	TotalTaxesOwedAfterStandardDeductionCalculation
}

func NewTaxOnTraditionalIRAWithdrawal() TaxOnTraditionalIRAWithdrawal {
	return TaxOnTraditionalIRAWithdrawal{
		TotalTaxesOwedAfterStandardDeductionCalculation: NewTotalTaxesOwedAfterStandardDeduction(),
	}
}

func (c TaxOnTraditionalIRAWithdrawal) CalculateTraditional(model *Model) float64 {
	return c.CalculateTraditionalRetirement(model)
}

func (c TaxOnTraditionalIRAWithdrawal) CalculateTraditionalRetirement(model *Model) float64 {
	totalTaxesOwedAfterStandardDeductionTraditional := c.TotalTaxesOwedAfterStandardDeductionCalculation.CalculateTraditionalRetirement(model)
	totalTaxesOwedAfterStandardDeductionRoth := c.TotalTaxesOwedAfterStandardDeductionCalculation.CalculateRothRetirement(model)

	return totalTaxesOwedAfterStandardDeductionTraditional - totalTaxesOwedAfterStandardDeductionRoth
}

func (c TaxOnTraditionalIRAWithdrawal) CalculateRoth(model *Model) float64 {
	return c.CalculateTraditionalRetirement(model)
}

func (c TaxOnTraditionalIRAWithdrawal) CalculateRothRetirement(model *Model) float64 {
	return c.CalculateTraditionalRetirement(model)
}
