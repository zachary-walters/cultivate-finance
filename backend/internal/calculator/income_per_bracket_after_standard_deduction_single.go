package calculator

type IncomePerBracketAfterStandardDeductionSingleCalculation SequenceCalculation

type IncomePerBracketAfterStandardDeductionSingle struct {
	IncomePerBracketAfterStandardDeductionCalculation
}

func NewIncomePerBracketAfterStandardDeductionSingle() IncomePerBracketAfterStandardDeductionSingle {
	return IncomePerBracketAfterStandardDeductionSingle{
		IncomePerBracketAfterStandardDeductionCalculation: NewIncomePerBracketAfterStandardDeduction(),
	}
}

func (c IncomePerBracketAfterStandardDeductionSingle) Calculate(model Model) []float64 {
	return c.IncomePerBracketAfterStandardDeductionCalculation.Calculate(model, model.SingleTaxRates)
}

func (c IncomePerBracketAfterStandardDeductionSingle) CalculateRetirement(model Model) []float64 {
	return c.IncomePerBracketAfterStandardDeductionCalculation.CalculateRetirement(model, model.SingleTaxRates)
}
