package calculator

type AdjustedGrossIncomeRothCalculation Calculation

type AdjustedGrossIncomeRoth struct{}

func NewAdjustedGrossIncomeRoth() AdjustedGrossIncomeRoth {
	return AdjustedGrossIncomeRoth{}
}

func (c AdjustedGrossIncomeRoth) Calculate(model Model) float64 {
	return model.Input.WorkIncome +
		model.Input.PensionIncome +
		model.Input.RentalNetIncome +
		model.Input.AnnuityIncome
}

func (c AdjustedGrossIncomeRoth) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
