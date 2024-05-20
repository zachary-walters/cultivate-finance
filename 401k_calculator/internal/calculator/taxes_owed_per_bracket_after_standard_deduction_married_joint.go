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

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedJoint) CalculateTraditional(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateTraditional(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := make([]float64, len(model.MarriedJointTaxRates))

	for idx, taxRate := range model.MarriedJointTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedJoint[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedJoint
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedJoint) CalculateTraditionalRetirement(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateTraditionalRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := make([]float64, len(model.MarriedJointTaxRates))

	for idx, taxRate := range model.MarriedJointTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedJoint[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedJoint
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedJoint) CalculateRoth(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateRoth(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := make([]float64, len(model.MarriedJointTaxRates))

	for idx, taxRate := range model.MarriedJointTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedJoint[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedJoint
}

func (c TaxesOwedPerBracketAfterStandardDeductionMarriedJoint) CalculateRothRetirement(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateRothRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := make([]float64, len(model.MarriedJointTaxRates))

	for idx, taxRate := range model.MarriedJointTaxRates {
		taxesOwedPerBracketAfterStandardDeductionMarriedJoint[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionMarriedJoint
}
