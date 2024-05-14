package calculator

type TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation SequenceCalculation

type TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint struct {
	IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation
}

func NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint() TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint {
	return TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint{
		IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation: NewIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint(),
	}
}

func (c TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateTraditional(model Model) []float64 {
	incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint := c.IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation.CalculateTraditional(model)

	taxesOwedPerBracketAfterStandardDudectionAndContributions := make([]float64, len(model.MarriedJointTaxRates))

	for idx, taxRate := range model.MarriedJointTaxRates {
		taxesOwedPerBracketAfterStandardDudectionAndContributions[idx] = incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDudectionAndContributions
}

func (c TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateTraditionalRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint := c.IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation.CalculateTraditionalRetirement(model)

	taxesOwedPerBracketAfterStandardDudectionAndContributions := make([]float64, len(model.MarriedJointTaxRates))

	for idx, taxRate := range model.MarriedJointTaxRates {
		taxesOwedPerBracketAfterStandardDudectionAndContributions[idx] = incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDudectionAndContributions
}

func (c TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateRoth(model Model) []float64 {
	incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint := c.IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation.CalculateRoth(model)

	taxesOwedPerBracketAfterStandardDudectionAndContributions := make([]float64, len(model.MarriedJointTaxRates))

	for idx, taxRate := range model.MarriedJointTaxRates {
		taxesOwedPerBracketAfterStandardDudectionAndContributions[idx] = incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDudectionAndContributions
}

func (c TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateRothRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint := c.IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation.CalculateRothRetirement(model)

	taxesOwedPerBracketAfterStandardDudectionAndContributions := make([]float64, len(model.MarriedJointTaxRates))

	for idx, taxRate := range model.MarriedJointTaxRates {
		taxesOwedPerBracketAfterStandardDudectionAndContributions[idx] = incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDudectionAndContributions
}
