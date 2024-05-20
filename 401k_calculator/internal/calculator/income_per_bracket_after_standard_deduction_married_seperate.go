package calculator

type IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation SequenceCalculation

type IncomePerBracketAfterStandardDeductionMarriedSeperate struct {
	AbstractIncomePerBracketAfterStandardDeductionCalculation
}

func NewIncomePerBracketAfterStandardDeductionMarriedSeperate() IncomePerBracketAfterStandardDeductionMarriedSeperate {
	return IncomePerBracketAfterStandardDeductionMarriedSeperate{
		AbstractIncomePerBracketAfterStandardDeductionCalculation: NewAbstractIncomePerBracketAfterStandardDeduction(),
	}
}

func (c IncomePerBracketAfterStandardDeductionMarriedSeperate) CalculateTraditional(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditional(model, model.MarriedSeperateTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionMarriedSeperate) CalculateTraditionalRetirement(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditionalRetirement(model, model.MarriedSeperateTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionMarriedSeperate) CalculateRoth(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRoth(model, model.MarriedSeperateTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionMarriedSeperate) CalculateRothRetirement(model *Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRothRetirement(model, model.MarriedSeperateTaxRates)
}
