package calculator

type IncomePerBracketAfterStandardDeductionCalculation SequenceCalculation

type IncomePerBracketAfterStandardDeduction struct {
	IncomePerBracketAfterStandardDeductionSingleCalculation
	IncomePerBracketAfterStandardDeductionMarriedJointCalculation
	IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation
	IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation
}

func NewIncomePerBracketAfterStandardDeduction() IncomePerBracketAfterStandardDeduction {
	return IncomePerBracketAfterStandardDeduction{
		IncomePerBracketAfterStandardDeductionSingleCalculation:          NewIncomePerBracketAfterStandardDeductionSingle(),
		IncomePerBracketAfterStandardDeductionMarriedJointCalculation:    NewIncomePerBracketAfterStandardDeductionMarriedJoint(),
		IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation: NewIncomePerBracketAfterStandardDeductionMarriedSeperate(),
		IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: NewIncomePerBracketAfterStandardDeductionHeadOfHousehold(),
	}
}

func (c IncomePerBracketAfterStandardDeduction) Calculate(model Model) []float64 {
	incomePerBracketAfterStandardDeductionSingle := c.IncomePerBracketAfterStandardDeductionSingleCalculation.Calculate(model)
	incomePerBracketAfterStandardDeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.Calculate(model)
	incomePerBracketAfterStandardDeductionMarriedSeperate := c.IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation.Calculate(model)
	incomePerBracketAfterStandardDeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.Calculate(model)

	switch model.Input.CurrentFilingStatus {
	case "single":
		return incomePerBracketAfterStandardDeductionSingle
	case "married-joint":
		return incomePerBracketAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return incomePerBracketAfterStandardDeductionMarriedSeperate
	case "head-of-household":
		return incomePerBracketAfterStandardDeductionHeadOfHousehold
	default:
		return nil
	}
}

func (c IncomePerBracketAfterStandardDeduction) CalculateRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeductionSingle := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateRetirement(model)
	incomePerBracketAfterStandardDeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateRetirement(model)
	incomePerBracketAfterStandardDeductionMarriedSeperate := c.IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation.CalculateRetirement(model)
	incomePerBracketAfterStandardDeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRetirement(model)

	switch model.Input.CurrentFilingStatus {
	case "single":
		return incomePerBracketAfterStandardDeductionSingle
	case "married-joint":
		return incomePerBracketAfterStandardDeductionMarriedJoint
	case "married-seperate":
		return incomePerBracketAfterStandardDeductionMarriedSeperate
	case "head-of-household":
		return incomePerBracketAfterStandardDeductionHeadOfHousehold
	default:
		return nil
	}
}
