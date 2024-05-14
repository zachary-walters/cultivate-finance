package calculator

type TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation SequenceCalculation

type TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle struct {
	IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation
}

func NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle() TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle {
	return TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle{
		IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation: NewIncomePerBracketAfterStandardDeductionAndContributionsSingle(),
	}
}

func (c TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle) CalculateTraditional(model Model) []float64 {
	incomePerBracketAfterStandardDeductionAndContributionsSingle := c.IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation.CalculateTraditional(model)

	taxesOwedPerBracketAfterStandardDudectionAndContributions := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDudectionAndContributions[idx] = incomePerBracketAfterStandardDeductionAndContributionsSingle[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDudectionAndContributions
}

func (c TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle) CalculateTraditionalRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeductionAndContributionsSingle := c.IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation.CalculateTraditionalRetirement(model)

	taxesOwedPerBracketAfterStandardDudectionAndContributions := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDudectionAndContributions[idx] = incomePerBracketAfterStandardDeductionAndContributionsSingle[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDudectionAndContributions
}

func (c TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle) CalculateRoth(model Model) []float64 {
	incomePerBracketAfterStandardDeductionAndContributionsSingle := c.IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation.CalculateRoth(model)

	taxesOwedPerBracketAfterStandardDudectionAndContributions := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDudectionAndContributions[idx] = incomePerBracketAfterStandardDeductionAndContributionsSingle[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDudectionAndContributions
}

func (c TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle) CalculateRothRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeductionAndContributionsSingle := c.IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation.CalculateRothRetirement(model)

	taxesOwedPerBracketAfterStandardDudectionAndContributions := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDudectionAndContributions[idx] = incomePerBracketAfterStandardDeductionAndContributionsSingle[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDudectionAndContributions
}
