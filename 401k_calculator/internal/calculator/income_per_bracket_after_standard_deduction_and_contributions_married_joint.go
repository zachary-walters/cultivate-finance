package calculator

type IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation SequenceCalculation

type IncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint struct {
	AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation
}

func NewIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint() IncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint {
	return IncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint{
		AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: NewAbstractIncomePerBracketAfterStandardDeductionAndContributions(),
	}
}

func (c IncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateTraditional(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateTraditional(model, model.MarriedJointTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateTraditionalRetirement(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateTraditionalRetirement(model, model.MarriedJointTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateRoth(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateRoth(model, model.MarriedJointTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateRothRetirement(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateRothRetirement(model, model.MarriedJointTaxRates)
}