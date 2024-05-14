package calculator

type AdjustedGrossIncomeCalculation Calculation

type AdjustedGrossIncome struct{}

func NewAdjustedGrossIncome() AdjustedGrossIncome {
	return AdjustedGrossIncome{}
}

func (c AdjustedGrossIncome) CalculateTraditional(model Model) float64 {
	return c.CalculateTraditionalRetirement(model)
}

func (c AdjustedGrossIncome) CalculateTraditionalRetirement(model Model) float64 {
	return model.Input.WorkIncome +
		model.Input.PensionIncome +
		model.Input.RentalNetIncome +
		model.Input.AnnuityIncome +
		model.Input.YearlyWithdrawal
}

func (c AdjustedGrossIncome) CalculateRoth(model Model) float64 {
	return c.CalculateRothRetirement(model)
}

func (c AdjustedGrossIncome) CalculateRothRetirement(model Model) float64 {
	return model.Input.WorkIncome +
		model.Input.PensionIncome +
		model.Input.RentalNetIncome +
		model.Input.AnnuityIncome
}
