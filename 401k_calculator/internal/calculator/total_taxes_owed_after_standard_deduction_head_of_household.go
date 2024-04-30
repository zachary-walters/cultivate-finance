package calculator

type TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation Calculation

type TotalTaxesOwedAfterStandardDeductionHeadOfHousehold struct {
	TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation
}

func NewTotalTaxesOwedAfterStandardDeductionHeadOfHousehold() TotalTaxesOwedAfterStandardDeductionHeadOfHousehold {
	return TotalTaxesOwedAfterStandardDeductionHeadOfHousehold{
		TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: NewTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold(),
	}
}

func (c TotalTaxesOwedAfterStandardDeductionHeadOfHousehold) Calculate(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := c.TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation.Calculate(model)

	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold {
		totalTaxesOwedAfterStandardDeductionHeadOfHousehold += value
	}

	return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
}

func (c TotalTaxesOwedAfterStandardDeductionHeadOfHousehold) CalculateRetirement(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := c.TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRetirement(model)

	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold {
		totalTaxesOwedAfterStandardDeductionHeadOfHousehold += value
	}

	return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
}
