package calculator

type IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation SequenceCalculation

type IncomePerBracketAfterStandardDeductionHeadOfHousehold struct {
	AbstractIncomePerBracketAfterStandardDeductionCalculation
}

func NewIncomePerBracketAfterStandardDeductionHeadOfHousehold() IncomePerBracketAfterStandardDeductionHeadOfHousehold {
	return IncomePerBracketAfterStandardDeductionHeadOfHousehold{
		AbstractIncomePerBracketAfterStandardDeductionCalculation: NewAbstractIncomePerBracketAfterStandardDeduction(),
	}
}

func (c IncomePerBracketAfterStandardDeductionHeadOfHousehold) CalculateTraditional(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditional(model, model.HeadOfHouseholdTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionHeadOfHousehold) CalculateTraditionalRetirement(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditionalRetirement(model, model.HeadOfHouseholdTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionHeadOfHousehold) CalculateRoth(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRoth(model, model.HeadOfHouseholdTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionHeadOfHousehold) CalculateRothRetirement(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRothRetirement(model, model.HeadOfHouseholdTaxRates)
}
