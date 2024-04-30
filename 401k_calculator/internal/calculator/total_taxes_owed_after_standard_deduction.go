package calculator

type TotalTaxesOwedAfterStandardDeductionCalculation Calculation

type TotalTaxesOwedAfterStandardDeduction struct {
	TotalTaxesOwedAfterStandardDeductionSingleCalculation
	TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation
	TotalTaxesOwedAfterStandardDeductionMarriedSeperateCalculation
	TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation
}

func NewTotalTaxesOwedAfterStandardDeduction() TotalTaxesOwedAfterStandardDeduction {
	return TotalTaxesOwedAfterStandardDeduction{
		TotalTaxesOwedAfterStandardDeductionSingleCalculation:          NewTotalTaxesOwedAfterStandardDeductionSingle(),
		TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation:    NewTotalTaxesOwedAfterStandardDeductionMarriedJoint(),
		TotalTaxesOwedAfterStandardDeductionMarriedSeperateCalculation: NewTotalTaxesOwedAfterStandardDeductionMarriedSeperate(),
		TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation: NewTotalTaxesOwedAfterStandardDeductionHeadOfHousehold(),
	}
}

func (c TotalTaxesOwedAfterStandardDeduction) Calculate(model Model) float64 {
	totalTaxesOwedAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.Calculate(model)
	totalTaxesOwedAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.Calculate(model)
	totalTaxesOwedAfterStandardDeductionMarriedSeperate := c.TotalTaxesOwedAfterStandardDeductionMarriedSeperateCalculation.Calculate(model)
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := c.TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation.Calculate(model)

	switch model.Input.CurrentFilingStatus {
	case "single":
		return totalTaxesOwedAfterStandardDeductionSingle
	case "married-joint":
		return totalTaxesOwedAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return totalTaxesOwedAfterStandardDeductionMarriedSeperate
	case "head-of-household":
		return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
	default:
		return 0
	}
}

func (c TotalTaxesOwedAfterStandardDeduction) CalculateRetirement(model Model) float64 {
	totalTaxesOwedAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.CalculateRetirement(model)
	totalTaxesOwedAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.CalculateRetirement(model)
	totalTaxesOwedAfterStandardDeductionMarriedSeperate := c.TotalTaxesOwedAfterStandardDeductionMarriedSeperateCalculation.CalculateRetirement(model)
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := c.TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRetirement(model)

	switch model.Input.RetirementFilingStatus {
	case "single":
		return totalTaxesOwedAfterStandardDeductionSingle
	case "married-joint":
		return totalTaxesOwedAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return totalTaxesOwedAfterStandardDeductionMarriedSeperate
	case "head-of-household":
		return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
	default:
		return 0
	}
}
