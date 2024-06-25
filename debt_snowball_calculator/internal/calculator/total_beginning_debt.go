package calculator

type TotalBeginningDebtCalculation Calculation

type TotalBeginningDebt struct {
	AbstractCalculation
}

func NewTotalBeginningDebt() *TotalBeginningDebt {
	return &TotalBeginningDebt{}
}

func (c *TotalBeginningDebt) Calculate(model *Model) float64 {
	total := 0.0

	for _, debt := range model.Input.Debts {
		total += debt.Amount
	}

	return c.SanitizeToZero(total)
}
