package calculator

type TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation SequenceCalculation

type TaxesOwedPerBracketAfterStandardDeductionMarriedJoint struct {
	IncomePerBracketAfterStandardDeductionMarriedJointCalculation
}

func NewTaxesOwedPerBracketAfterStandardDeductionMarriedJoint() TaxesOwedPerBracketAfterStandardDeductionMarriedJoint {
	return TaxesOwedPerBracketAfterStandardDeductionMarriedJoint{
		IncomePerBracketAfterStandardDeductionMarriedJointCalculation: NewIncomePerBracketAfterStandardDeductionMarriedJoint(),
	}
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedJoint) Calculate(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.Calculate(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := make([]float64, len(model.MarriedJointTaxRates))

	for idx, taxRate := range model.MarriedJointTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedJoint[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedJoint
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedJoint) CalculateRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := make([]float64, len(model.MarriedJointTaxRates))

	for idx, taxRate := range model.MarriedJointTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedJoint[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedJoint
}
