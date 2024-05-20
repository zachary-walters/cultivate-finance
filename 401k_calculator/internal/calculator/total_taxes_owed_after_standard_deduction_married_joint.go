package calculator

type TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation Calculation

type TotalTaxesOwedAfterStandardDeductionMarriedJoint struct {
	TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation
}

func NewTotalTaxesOwedAfterStandardDeductionMarriedJoint() TotalTaxesOwedAfterStandardDeductionMarriedJoint {
	return TotalTaxesOwedAfterStandardDeductionMarriedJoint{
		TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation: NewTaxesOwedPerBracketAfterStandardDeductionMarriedJoint(),
	}
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedJoint) CalculateTraditional(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := c.TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation.CalculateTraditional(model)

	totalTaxesOwedAfterStandardDeductionMarriedJoint := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedJoint {
		totalTaxesOwedAfterStandardDeductionMarriedJoint += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedJoint
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedJoint) CalculateTraditionalRetirement(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := c.TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation.CalculateTraditionalRetirement(model)

	totalTaxesOwedAfterStandardDeductionMarriedJoint := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedJoint {
		totalTaxesOwedAfterStandardDeductionMarriedJoint += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedJoint
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedJoint) CalculateRoth(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := c.TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation.CalculateRoth(model)

	totalTaxesOwedAfterStandardDeductionMarriedJoint := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedJoint {
		totalTaxesOwedAfterStandardDeductionMarriedJoint += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedJoint
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedJoint) CalculateRothRetirement(model *Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := c.TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation.CalculateRothRetirement(model)

	totalTaxesOwedAfterStandardDeductionMarriedJoint := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedJoint {
		totalTaxesOwedAfterStandardDeductionMarriedJoint += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedJoint
}
