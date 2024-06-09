package calculator

type IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation SequenceCalculation

type IncomePerBracketAfterStandardDeductionMarriedSeparate struct {
	AbstractIncomePerBracketAfterStandardDeductionCalculation
}

func NewIncomePerBracketAfterStandardDeductionMarriedSeparate() IncomePerBracketAfterStandardDeductionMarriedSeparate {
	return IncomePerBracketAfterStandardDeductionMarriedSeparate{
		AbstractIncomePerBracketAfterStandardDeductionCalculation: NewAbstractIncomePerBracketAfterStandardDeduction(),
	}
}

func (c IncomePerBracketAfterStandardDeductionMarriedSeparate) CalculateTraditional(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditional(model, model.MarriedSeparateTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionMarriedSeparate) CalculateTraditionalRetirement(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditionalRetirement(model, model.MarriedSeparateTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionMarriedSeparate) CalculateRoth(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRoth(model, model.MarriedSeparateTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionMarriedSeparate) CalculateRothRetirement(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRothRetirement(model, model.MarriedSeparateTaxRates)
}
