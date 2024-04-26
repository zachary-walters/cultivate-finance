package calculator

type IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation SequenceCalculation

type IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold struct {
	IncomePerBracketAfterStandardDeductionAndContributionsCalculation
}

func NewIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold() IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold {
	return IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
		IncomePerBracketAfterStandardDeductionAndContributionsCalculation: NewIncomePerBracketAfterStandardDeductionAndContributions(),
	}
}

func (c IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) Calculate(model Model) []float64 {
	return c.IncomePerBracketAfterStandardDeductionAndContributionsCalculation.Calculate(model, model.HeadOfHouseholdTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateRetirement(model Model) []float64 {
	return c.IncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateRetirement(model, model.HeadOfHouseholdTaxRates)
}