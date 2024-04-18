package calculator

type StandardDeductionCalculation Calculation

type StandardDeduction struct{}

func NewStandardDeduction() StandardDeduction {
	return StandardDeduction{}
}

func (c StandardDeduction) Calculate(model Model) float64 {
	switch model.Input.CurrentFilingStatus {
	case "single":
		return model.STANDARD_DEDUCTION_SINGLE
	case "married_joint":
		return model.STANDARD_DEDUCTION_MARRIED_JOINT
	case "married_seperate":
		return model.STANDARD_DEDUCTION_MARRIED_SEPERATE
	case "head_of_household":
		return model.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD
	default:
		return model.STANDARD_DEDUCTION_SINGLE
	}
}

func (c StandardDeduction) CalculateRetirement(model Model) float64 {
	switch model.Input.RetirementFilingStatus {
	case "single":
		return model.STANDARD_DEDUCTION_SINGLE
	case "married_joint":
		return model.STANDARD_DEDUCTION_MARRIED_JOINT
	case "married_seperate":
		return model.STANDARD_DEDUCTION_MARRIED_SEPERATE
	case "head_of_household":
		return model.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD
	default:
		return model.STANDARD_DEDUCTION_SINGLE
	}
}
