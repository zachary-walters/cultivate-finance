package calculator

type TotalBeginningDebtCalculation Calculation

type TotalBeginningDebt struct {
	AbstractCalculation
	ValidDebtsCalculation
}

func NewTotalBeginningDebt() *TotalBeginningDebt {
	return &TotalBeginningDebt{
		ValidDebtsCalculation: NewValidDebts(),
	}
}

func (c *TotalBeginningDebt) CalculateSnowball(model Model) float64 {
	validDebts := c.ValidDebtsCalculation.CalculateSnowball(model)

	total := 0.0

	for _, debt := range validDebts {
		total += debt.Amount
	}

	return c.SanitizeToZero(total)
}

func (c *TotalBeginningDebt) CalculateAvalanche(model Model) float64 {
	return c.CalculateSnowball(model)
}
