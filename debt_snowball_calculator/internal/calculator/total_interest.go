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

func (c TotalInterest) CalculateSnowball(model Model) float64 {
	totalBeginningDebt := c.TotalBeginningDebtCalculation.CalculateSnowball(model)
	totalPayments := c.TotalPaymentsCalculation.CalculateSnowball(model)

	return c.SanitizeToZero(totalPayments - totalBeginningDebt)
}

func (c TotalInterest) CalculateAvalanche(model Model) float64 {
	totalBeginningDebt := c.TotalBeginningDebtCalculation.CalculateAvalanche(model)
	totalPayments := c.TotalPaymentsCalculation.CalculateAvalanche(model)

	return c.SanitizeToZero(totalPayments - totalBeginningDebt)
}
