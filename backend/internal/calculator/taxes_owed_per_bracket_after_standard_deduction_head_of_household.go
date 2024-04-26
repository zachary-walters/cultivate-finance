package calculator

type TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation SequenceCalculation

type TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold struct {
	IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation
}

func NewTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold() TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold {
	return TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold{
		IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: NewIncomePerBracketAfterStandardDeductionHeadOfHousehold(),
	}
}

func (c TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) Calculate(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.Calculate(model)

	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := make([]float64, len(model.HeadOfHouseholdTaxRates))

	for idx, taxRate := range model.HeadOfHouseholdTaxRates {
		taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold
}

func (c TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := make([]float64, len(model.HeadOfHouseholdTaxRates))

	for idx, taxRate := range model.HeadOfHouseholdTaxRates {
		taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold
}
