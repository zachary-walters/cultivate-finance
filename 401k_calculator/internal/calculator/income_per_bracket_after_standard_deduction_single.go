package calculator

type IncomePerBracketAfterStandardDeductionSingleCalculation SequenceCalculation

type IncomePerBracketAfterStandardDeductionSingle struct {
	AbstractIncomePerBracketAfterStandardDeductionCalculation
}

func NewIncomePerBracketAfterStandardDeductionSingle() IncomePerBracketAfterStandardDeductionSingle {
	return IncomePerBracketAfterStandardDeductionSingle{
		AbstractIncomePerBracketAfterStandardDeductionCalculation: NewAbstractIncomePerBracketAfterStandardDeduction(),
	}
}

func (c IncomePerBracketAfterStandardDeductionSingle) CalculateTraditional(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditional(model, model.SingleTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionSingle) CalculateTraditionalRetirement(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditionalRetirement(model, model.SingleTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionSingle) CalculateRoth(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRoth(model, model.SingleTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionSingle) CalculateRothRetirement(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRothRetirement(model, model.SingleTaxRates)
}
