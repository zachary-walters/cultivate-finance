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

func (c IncomePerBracketAfterStandardDeductionSingle) Calculate(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.Calculate(model, model.SingleTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionSingle) CalculateRetirement(model Model) []float64 {
	return c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRetirement(model, model.SingleTaxRates)
}
