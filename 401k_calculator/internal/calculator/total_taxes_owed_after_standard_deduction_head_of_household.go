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

func (c TotalTaxesOwedAfterStandardDeductionHeadOfHousehold) CalculateTraditional(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := c.TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditional(model)

	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold {
		totalTaxesOwedAfterStandardDeductionHeadOfHousehold += value
	}

	return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
}

func (c TotalTaxesOwedAfterStandardDeductionHeadOfHousehold) CalculateTraditionalRetirement(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := c.TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditionalRetirement(model)

	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold {
		totalTaxesOwedAfterStandardDeductionHeadOfHousehold += value
	}

	return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
}

func (c TotalTaxesOwedAfterStandardDeductionHeadOfHousehold) CalculateRoth(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := c.TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRoth(model)

	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold {
		totalTaxesOwedAfterStandardDeductionHeadOfHousehold += value
	}

	return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
}

func (c TotalTaxesOwedAfterStandardDeductionHeadOfHousehold) CalculateRothRetirement(model Model) float64 {
	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := c.TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRothRetirement(model)

	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := 0.0

	for _, value := range taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold {
		totalTaxesOwedAfterStandardDeductionHeadOfHousehold += value
	}

	return totalTaxesOwedAfterStandardDeductionHeadOfHousehold
}
