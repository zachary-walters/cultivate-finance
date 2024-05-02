package calculator

type AdjustedGrossIncomeTraditionalCalculation Calculation

type AdjustedGrossIncomeTraditional struct{}

func NewAdjustedGrossIncomeTraditional() AdjustedGrossIncomeTraditional {
	return AdjustedGrossIncomeTraditional{}
}

func (c AdjustedGrossIncomeTraditional) Calculate(model Model) float64 {
	return model.Input.WorkIncome +
		model.Input.PensionIncome +
		model.Input.RentalNetIncome +
		model.Input.AnnuityIncome +
		model.Input.YearlyWithdrawal
}

func (c AdjustedGrossIncomeTraditional) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
