package calculator

type TotalInterestCalculation Calculation

type TotalInterest struct {
	AbstractCalculation
	TotalBeginningDebtCalculation
	TotalPaymentsCalculation
}

func NewTotalInterest() *TotalInterest {
	return &TotalInterest{
		TotalBeginningDebtCalculation: NewTotalBeginningDebt(),
		TotalPaymentsCalculation:      NewTotalPayments(),
	}
}

func (c TotalInterest) Calculate(model *Model) float64 {
	totalBeginningDebt := c.TotalBeginningDebtCalculation.Calculate(model)
	totalPayments := c.TotalPaymentsCalculation.Calculate(model)

	return c.SanitizeToZero(totalPayments - totalBeginningDebt)
}
