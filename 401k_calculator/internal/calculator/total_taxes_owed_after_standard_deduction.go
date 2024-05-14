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

func (c TotalTaxesOwedAfterStandardDeduction) CalculateTraditional(model Model) float64 {
	totalTaxesOwedAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.CalculateTraditional(model)
	totalTaxesOwedAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.CalculateTraditional(model)
	totalTaxesOwedAfterStandardDeductionMarriedSeperate := c.TotalTaxesOwedAfterStandardDeductionMarriedSeperateCalculation.CalculateTraditional(model)
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := c.TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditional(model)

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

func (c TotalTaxesOwedAfterStandardDeduction) CalculateTraditionalRetirement(model Model) float64 {
	totalTaxesOwedAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.CalculateTraditionalRetirement(model)
	totalTaxesOwedAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.CalculateTraditionalRetirement(model)
	totalTaxesOwedAfterStandardDeductionMarriedSeperate := c.TotalTaxesOwedAfterStandardDeductionMarriedSeperateCalculation.CalculateTraditionalRetirement(model)
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := c.TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditionalRetirement(model)

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

func (c TotalTaxesOwedAfterStandardDeduction) CalculateRoth(model Model) float64 {
	totalTaxesOwedAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.CalculateRoth(model)
	totalTaxesOwedAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.CalculateRoth(model)
	totalTaxesOwedAfterStandardDeductionMarriedSeperate := c.TotalTaxesOwedAfterStandardDeductionMarriedSeperateCalculation.CalculateRoth(model)
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := c.TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRoth(model)

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

func (c TotalTaxesOwedAfterStandardDeduction) CalculateRothRetirement(model Model) float64 {
	totalTaxesOwedAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.CalculateRothRetirement(model)
	totalTaxesOwedAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.CalculateRothRetirement(model)
	totalTaxesOwedAfterStandardDeductionMarriedSeperate := c.TotalTaxesOwedAfterStandardDeductionMarriedSeperateCalculation.CalculateRothRetirement(model)
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := c.TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRothRetirement(model)

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
