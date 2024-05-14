package calculator

type IncomePerBracketAfterStandardDeductionMarriedJointCalculation SequenceCalculation

type IncomePerBracketAfterStandardDeductionMarriedJoint struct {
	AbstractIncomePerBracketAfterStandardDeductionCalculation
}

func NewIncomePerBracketAfterStandardDeductionMarriedJoint() IncomePerBracketAfterStandardDeductionMarriedJoint {
	return IncomePerBracketAfterStandardDeductionMarriedJoint{
		AbstractIncomePerBracketAfterStandardDeductionCalculation: NewAbstractIncomePerBracketAfterStandardDeduction(),
	}
}

func (c IncomePerBracketAfterStandardDeductionMarriedJoint) CalculateTraditional(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditional(model, model.MarriedJointTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionMarriedJoint) CalculateTraditionalRetirement(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditionalRetirement(model, model.MarriedJointTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionMarriedJoint) CalculateRoth(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRoth(model, model.MarriedJointTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionMarriedJoint) CalculateRothRetirement(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRothRetirement(model, model.MarriedJointTaxRates)
}
