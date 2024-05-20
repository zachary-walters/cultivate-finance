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

func (c TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateTraditional(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditional(model)

	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := make([]float64, len(model.HeadOfHouseholdTaxRates))

	for idx, taxRate := range model.HeadOfHouseholdTaxRates {
		taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold
}

func (c TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateTraditionalRetirement(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditionalRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := make([]float64, len(model.HeadOfHouseholdTaxRates))

	for idx, taxRate := range model.HeadOfHouseholdTaxRates {
		taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold
}

func (c TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateRoth(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRoth(model)

	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := make([]float64, len(model.HeadOfHouseholdTaxRates))

	for idx, taxRate := range model.HeadOfHouseholdTaxRates {
		taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold
}

func (c TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateRothRetirement(model *Model) []float64 {
	incomePerBracketAfterStandardDeduction := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRothRetirement(model)

	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := make([]float64, len(model.HeadOfHouseholdTaxRates))

	for idx, taxRate := range model.HeadOfHouseholdTaxRates {
		taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold[idx] = incomePerBracketAfterStandardDeduction[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold
}
