package calculator

type TotalTaxesOwedAfterStandardDeductionCalculation Calculation

type TotalTaxesOwedAfterStandardDeduction struct {
	TotalTaxesOwedAfterStandardDeductionSingleCalculation
	TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation
	TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation
	TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation
}

func NewTotalTaxesOwedAfterStandardDeduction() TotalTaxesOwedAfterStandardDeduction {
	return TotalTaxesOwedAfterStandardDeduction{
		TotalTaxesOwedAfterStandardDeductionSingleCalculation:          NewTotalTaxesOwedAfterStandardDeductionSingle(),
		TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation:    NewTotalTaxesOwedAfterStandardDeductionMarriedJoint(),
		TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation: NewTotalTaxesOwedAfterStandardDeductionMarriedSeparate(),
		TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation: NewTotalTaxesOwedAfterStandardDeductionHeadOfHousehold(),
	}
}

func (c TotalTaxesOwedAfterStandardDeduction) CalculateTraditional(model *Model) float64 {
	totalTaxesOwedAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.CalculateTraditional(model)
	totalTaxesOwedAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.CalculateTraditional(model)
	totalTaxesOwedAfterStandardDeductionMarriedSeparate := c.TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation.CalculateTraditional(model)
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := c.TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditional(model)

	switch model.Input.CurrentFilingStatus {
	case "single":
		return totalTaxesOwedAfterStandardDeductionSingle
	case "married-joint":
		return totalTaxesOwedAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return totalTaxesOwedAfterStandardDeductionMarriedSeparate
	case "head-of-household":
		return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
	default:
		return 0
	}
}

func (c TotalTaxesOwedAfterStandardDeduction) CalculateTraditionalRetirement(model *Model) float64 {
	totalTaxesOwedAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.CalculateTraditionalRetirement(model)
	totalTaxesOwedAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.CalculateTraditionalRetirement(model)
	totalTaxesOwedAfterStandardDeductionMarriedSeparate := c.TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation.CalculateTraditionalRetirement(model)
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := c.TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditionalRetirement(model)

	switch model.Input.RetirementFilingStatus {
	case "single":
		return totalTaxesOwedAfterStandardDeductionSingle
	case "married-joint":
		return totalTaxesOwedAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return totalTaxesOwedAfterStandardDeductionMarriedSeparate
	case "head-of-household":
		return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
	default:
		return 0
	}
}

func (c TotalTaxesOwedAfterStandardDeduction) CalculateRoth(model *Model) float64 {
	totalTaxesOwedAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.CalculateRoth(model)
	totalTaxesOwedAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.CalculateRoth(model)
	totalTaxesOwedAfterStandardDeductionMarriedSeparate := c.TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation.CalculateRoth(model)
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := c.TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRoth(model)

	switch model.Input.CurrentFilingStatus {
	case "single":
		return totalTaxesOwedAfterStandardDeductionSingle
	case "married-joint":
		return totalTaxesOwedAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return totalTaxesOwedAfterStandardDeductionMarriedSeparate
	case "head-of-household":
		return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
	default:
		return 0
	}
}

func (c TotalTaxesOwedAfterStandardDeduction) CalculateRothRetirement(model *Model) float64 {
	totalTaxesOwedAfterStandardDeductionSingle := c.TotalTaxesOwedAfterStandardDeductionSingleCalculation.CalculateRothRetirement(model)
	totalTaxesOwedAfterStandardDeductionMarriedJoint := c.TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation.CalculateRothRetirement(model)
	totalTaxesOwedAfterStandardDeductionMarriedSeparate := c.TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation.CalculateRothRetirement(model)
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := c.TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRothRetirement(model)

	switch model.Input.RetirementFilingStatus {
	case "single":
		return totalTaxesOwedAfterStandardDeductionSingle
	case "married-joint":
		return totalTaxesOwedAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return totalTaxesOwedAfterStandardDeductionMarriedSeparate
	case "head-of-household":
		return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
	default:
		return 0
	}
}
