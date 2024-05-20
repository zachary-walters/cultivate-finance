package calculator

type StandardDeductionCalculation Calculation

type StandardDeduction struct{}

func NewStandardDeduction() StandardDeduction {
	return StandardDeduction{}
}

func (c StandardDeduction) CalculateTraditional(model *Model) float64 {
	switch model.Input.CurrentFilingStatus {
	case "single":
		return model.STANDARD_DEDUCTION_SINGLE
	case "married-joint":
		return model.STANDARD_DEDUCTION_MARRIED_JOINT
	case "married-seperate":
		return model.STANDARD_DEDUCTION_MARRIED_SEPERATE
	case "head-of-household":
		return model.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD
	default:
		return model.STANDARD_DEDUCTION_SINGLE
	}
}

func (c StandardDeduction) CalculateTraditionalRetirement(model *Model) float64 {
	switch model.Input.RetirementFilingStatus {
	case "single":
		return model.STANDARD_DEDUCTION_SINGLE
	case "married-joint":
		return model.STANDARD_DEDUCTION_MARRIED_JOINT
	case "married-seperate":
		return model.STANDARD_DEDUCTION_MARRIED_SEPERATE
	case "head-of-household":
		return model.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD
	default:
		return model.STANDARD_DEDUCTION_SINGLE
	}
}

func (c StandardDeduction) CalculateRoth(model *Model) float64 {
	return c.CalculateTraditional(model)
}

func (c StandardDeduction) CalculateRothRetirement(model *Model) float64 {
	return c.CalculateTraditionalRetirement(model)
}
