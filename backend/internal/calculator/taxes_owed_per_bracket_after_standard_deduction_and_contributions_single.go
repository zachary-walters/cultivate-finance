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

func (c TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle) Calculate(model Model) []float64 {
	incomePerBracketAfterStandardDeductionAndContributionsSingle := c.IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation.Calculate(model)

	taxesOwedPerBracketAfterStandardDudectionAndContributions := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDudectionAndContributions[idx] = incomePerBracketAfterStandardDeductionAndContributionsSingle[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDudectionAndContributions
}

func (c TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle) CalculateRetirement(model Model) []float64 {
	incomePerBracketAfterStandardDeductionAndContributionsSingle := c.IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation.Calculate(model)

	taxesOwedPerBracketAfterStandardDudectionAndContributions := make([]float64, len(model.SingleTaxRates))

	for idx, taxRate := range model.SingleTaxRates {
		taxesOwedPerBracketAfterStandardDudectionAndContributions[idx] = incomePerBracketAfterStandardDeductionAndContributionsSingle[idx] * taxRate.Rate
	}

	return taxesOwedPerBracketAfterStandardDudectionAndContributions
}
