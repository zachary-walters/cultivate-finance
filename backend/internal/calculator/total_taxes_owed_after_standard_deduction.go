package calculator

type TotalTaxesOwedAfterStandardDeductionCalculation Calculation

type TotalTaxesOwedAfterStandardDeduction struct {
	TotalTaxesOwedAfterStandardDeductionSingleCalculation
	TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation
}

func NewTotalTaxesOwedAfterStandardDeduction() TotalTaxesOwedAfterStandardDeduction {
	return TotalTaxesOwedAfterStandardDeduction{
		TotalTaxesOwedAfterStandardDeductionSingleCalculation:       NewTotalTaxesOwedAfterStandardDeductionSingle(),
		TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation: NewTotalTaxesOwedAfterStandardDeductionMarriedJoint(),
	}
}

func (c TotalTaxesOwedAfterStandardDeduction) Calculate(model Model) float64 {
	totalTaxesOwedPerBracketAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.Calculate(model)
	totalTaxesOwedPerBracketAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.Calculate(model)

	switch model.Input.CurrentFilingStatus {
	case "single":
		return totalTaxesOwedPerBracketAfterStandardDeductionSingle
	case "married-joint":
		return totalTaxesOwedPerBracketAfterStandardDeductionMarriedJoint
	default:
		return 0
	}
}

func (c TotalTaxesOwedAfterStandardDeduction) CalculateRetirement(model Model) float64 {
	totalTaxesOwedPerBracketAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.CalculateRetirement(model)
	totalTaxesOwedPerBracketAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.CalculateRetirement(model)

	switch model.Input.RetirementFilingStatus {
	case "single":
		return totalTaxesOwedPerBracketAfterStandardDeductionSingle
	case "married-joint":
		return totalTaxesOwedPerBracketAfterStandardDeductionMarriedJoint
	default:
		return 0
	}
}
