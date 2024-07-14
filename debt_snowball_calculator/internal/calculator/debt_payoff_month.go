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

func (c *DebtPayoffMonth) CalculateSnowball(model *Model) float64 {
	snowball := c.SnowballCalculation.CalculateSnowball(model)

	if len(snowball) <= 0 {
		return 0
	}

	lastDebtSequence := snowball[len(snowball)-1]

	for i := len(snowball) - 1; i >= 0; i-- {
		if !snowball[i].Invalid {
			lastDebtSequence = snowball[i]
			break
		}
	}

	if len(lastDebtSequence.Months) <= 0 {
		return 0
	}

	lastMonth := lastDebtSequence.Months[len(lastDebtSequence.Months)-1]

	return lastMonth
}

func (c *DebtPayoffMonth) CalculateAvalanche(model *Model) float64 {
	snowball := c.SnowballCalculation.CalculateAvalanche(model)

	if len(snowball) <= 0 {
		return 0
	}

	lastDebtSequence := snowball[len(snowball)-1]

	for i := len(snowball) - 1; i >= 0; i-- {
		if !snowball[i].Invalid {
			lastDebtSequence = snowball[i]
			break
		}
	}

	if len(lastDebtSequence.Months) <= 0 {
		return 0
	}

	lastMonth := lastDebtSequence.Months[len(lastDebtSequence.Months)-1]

	return lastMonth
}
