package calculator

type DebtPayoffMonthCalculation Calculation

type DebtPayoffMonth struct {
	SnowballAvalancheCalculation
}

func NewDebtPayoffMonth() *DebtPayoffMonth {
	return &DebtPayoffMonth{
		SnowballAvalancheCalculation: NewSnowballAvalanche(),
	}
}

func (c *DebtPayoffMonth) CalculateSnowball(model Model) float64 {
	snowball := c.SnowballAvalancheCalculation.CalculateSnowball(model)

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

func (c *DebtPayoffMonth) CalculateAvalanche(model Model) float64 {
	snowball := c.SnowballAvalancheCalculation.CalculateAvalanche(model)

	payoffMonth := 0.0

	if len(snowball) <= 0 {
		return payoffMonth
	}

	for _, debtSequence := range snowball {
		if len(debtSequence.Months) <= 0 {
			continue
		}

		if debtSequence.Months[len(debtSequence.Months)-1] > payoffMonth {
			payoffMonth = debtSequence.Months[len(debtSequence.Months)-1]
		}
	}

	return payoffMonth
}
