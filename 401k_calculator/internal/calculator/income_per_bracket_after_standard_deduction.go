package calculator

type IncomePerBracketAfterStandardDeductionCalculation SequenceCalculation

type IncomePerBracketAfterStandardDeduction struct {
	IncomePerBracketAfterStandardDeductionSingleCalculation
	IncomePerBracketAfterStandardDeductionMarriedJointCalculation
	IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation
	IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation
}

func NewIncomePerBracketAfterStandardDeduction() IncomePerBracketAfterStandardDeduction {
	return IncomePerBracketAfterStandardDeduction{
		IncomePerBracketAfterStandardDeductionSingleCalculation:          NewIncomePerBracketAfterStandardDeductionSingle(),
		IncomePerBracketAfterStandardDeductionMarriedJointCalculation:    NewIncomePerBracketAfterStandardDeductionMarriedJoint(),
		IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation: NewIncomePerBracketAfterStandardDeductionMarriedSeparate(),
		IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: NewIncomePerBracketAfterStandardDeductionHeadOfHousehold(),
	}
}

func (c IncomePerBracketAfterStandardDeduction) CalculateTraditional(model *Model) []float64 {
	incomePerBracketAfterStandardDeductionSingle := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateTraditional(model)
	incomePerBracketAfterStandardDeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateTraditional(model)
	incomePerBracketAfterStandardDeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateTraditional(model)
	incomePerBracketAfterStandardDeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditional(model)

	switch model.Input.CurrentFilingStatus {
	case "single":
		return incomePerBracketAfterStandardDeductionSingle
	case "married-joint":
		return incomePerBracketAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return incomePerBracketAfterStandardDeductionMarriedSeparate
	case "head-of-household":
		return incomePerBracketAfterStandardDeductionHeadOfHousehold
	default:
		return incomePerBracketAfterStandardDeductionSingle
	}
}

func (c IncomePerBracketAfterStandardDeduction) CalculateTraditionalRetirement(model *Model) []float64 {
	incomePerBracketAfterStandardDeductionSingle := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateTraditionalRetirement(model)
	incomePerBracketAfterStandardDeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateTraditionalRetirement(model)
	incomePerBracketAfterStandardDeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateTraditionalRetirement(model)
	incomePerBracketAfterStandardDeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditionalRetirement(model)

	switch model.Input.RetirementFilingStatus {
	case "single":
		return incomePerBracketAfterStandardDeductionSingle
	case "married-joint":
		return incomePerBracketAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return incomePerBracketAfterStandardDeductionMarriedSeparate
	case "head-of-household":
		return incomePerBracketAfterStandardDeductionHeadOfHousehold
	default:
		return incomePerBracketAfterStandardDeductionSingle
	}
}

func (c IncomePerBracketAfterStandardDeduction) CalculateRoth(model *Model) []float64 {
	incomePerBracketAfterStandardDeductionSingle := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateRoth(model)
	incomePerBracketAfterStandardDeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateRoth(model)
	incomePerBracketAfterStandardDeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateRoth(model)
	incomePerBracketAfterStandardDeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRoth(model)

	switch model.Input.CurrentFilingStatus {
	case "single":
		return incomePerBracketAfterStandardDeductionSingle
	case "married-joint":
		return incomePerBracketAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return incomePerBracketAfterStandardDeductionMarriedSeparate
	case "head-of-household":
		return incomePerBracketAfterStandardDeductionHeadOfHousehold
	default:
		return incomePerBracketAfterStandardDeductionSingle
	}
}

func (c IncomePerBracketAfterStandardDeduction) CalculateRothRetirement(model *Model) []float64 {
	incomePerBracketAfterStandardDeductionSingle := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateRothRetirement(model)
	incomePerBracketAfterStandardDeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateRothRetirement(model)
	incomePerBracketAfterStandardDeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateRothRetirement(model)
	incomePerBracketAfterStandardDeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRothRetirement(model)

	switch model.Input.RetirementFilingStatus {
	case "single":
		return incomePerBracketAfterStandardDeductionSingle
	case "married-joint":
		return incomePerBracketAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return incomePerBracketAfterStandardDeductionMarriedSeparate
	case "head-of-household":
		return incomePerBracketAfterStandardDeductionHeadOfHousehold
	default:
		return incomePerBracketAfterStandardDeductionSingle
	}
}
