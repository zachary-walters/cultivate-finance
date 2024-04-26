package calculator

type IncomePerBracketAfterStandardDeductionMarriedJointCalculation SequenceCalculation

type IncomePerBracketAfterStandardDeductionMarriedJoint struct {
	IncomePerBracketAfterStandardDeductionCalculation
}

func NewIncomePerBracketAfterStandardDeductionMarriedJoint() IncomePerBracketAfterStandardDeductionMarriedJoint {
	return IncomePerBracketAfterStandardDeductionMarriedJoint{
		IncomePerBracketAfterStandardDeductionCalculation: NewIncomePerBracketAfterStandardDeduction(),
	}
}

func (c IncomePerBracketAfterStandardDeductionMarriedJoint) Calculate(model Model) []float64 {
	return c.IncomePerBracketAfterStandardDeductionCalculation.Calculate(model, model.MarriedJointTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionMarriedJoint) CalculateRetirement(model Model) []float64 {
	return c.IncomePerBracketAfterStandardDeductionCalculation.CalculateRetirement(model, model.MarriedJointTaxRates)
}
