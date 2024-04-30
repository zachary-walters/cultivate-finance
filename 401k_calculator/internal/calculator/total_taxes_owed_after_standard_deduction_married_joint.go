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

func (c TotalTaxesOwedAfterStandardDeductionMarriedJoint) Calculate(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := c.TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation.Calculate(model)

	totalTaxesOwedAfterStandardDeductionMarriedJoint := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedJoint {
		totalTaxesOwedAfterStandardDeductionMarriedJoint += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedJoint
}

func (c TotalTaxesOwedAfterStandardDeductionMarriedJoint) CalculateRetirement(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionMarriedJoint := c.TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation.CalculateRetirement(model)

	totalTaxesOwedAfterStandardDeductionMarriedJoint := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionMarriedJoint {
		totalTaxesOwedAfterStandardDeductionMarriedJoint += value
	}

	return totalTaxesOwedAfterStandardDeductionMarriedJoint
}
