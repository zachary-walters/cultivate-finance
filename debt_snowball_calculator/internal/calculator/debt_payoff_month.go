package calculator

type DebtPayoffMonthCalculation Calculation

type DebtPayoffMonth struct {
	SnowballCalculation
}

func NewDebtPayoffMonth() *DebtPayoffMonth {
	return &DebtPayoffMonth{
		SnowballCalculation: NewSnowball(),
	}
}

func (c *DebtPayoffMonth) Calculate(model *Model) float64 {
	snowball := c.SnowballCalculation.Calculate(model)

	if len(snowball) <= 0 {
		return 0
	}

	lastDebtSequence := snowball[len(snowball)-1]

	if len(lastDebtSequence.Months) <= 0 {
		return 0
	}

	lastMonth := lastDebtSequence.Months[len(lastDebtSequence.Months)-1]

	return lastMonth
}
