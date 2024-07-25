package calculator

type MonthlySequenceBalancesCalculation SequenceCalculation

type MonthlySequenceBalances struct {
	DebtPayoffMonthCalculation
	SnowballAvalancheCalculation
	TotalBeginningDebtCalculation
}

func NewMonthlySequenceBalances() *MonthlySequenceBalances {
	return &MonthlySequenceBalances{
		DebtPayoffMonthCalculation:    NewDebtPayoffMonth(),
		SnowballAvalancheCalculation:  NewSnowballAvalanche(),
		TotalBeginningDebtCalculation: NewTotalBeginningDebt(),
	}
}

func (c *MonthlySequenceBalances) CalculateSnowball(model Model) []float64 {
	debtPayoffMonth := c.DebtPayoffMonthCalculation.CalculateSnowball(model)
	snowball := c.SnowballAvalancheCalculation.CalculateSnowball(model)
	totalBeginningDebt := c.TotalBeginningDebtCalculation.CalculateSnowball(model)

	balances := []float64{
		totalBeginningDebt,
	}

	for i := 0; i < int(debtPayoffMonth); i++ {
		balance := 0.0
		for _, debtSequence := range snowball {
			if len(debtSequence.Balances) > i {
				balance += debtSequence.Balances[i]
			}
		}

		balances = append(balances, balance)
	}

	return balances
}

func (c *MonthlySequenceBalances) CalculateAvalanche(model Model) []float64 {
	debtPayoffMonth := c.DebtPayoffMonthCalculation.CalculateAvalanche(model)
	snowball := c.SnowballAvalancheCalculation.CalculateAvalanche(model)
	totalBeginningDebt := c.TotalBeginningDebtCalculation.CalculateAvalanche(model)

	balances := []float64{
		totalBeginningDebt,
	}

	for i := 0; i < int(debtPayoffMonth); i++ {
		balance := 0.0
		for _, debtSequence := range snowball {
			if len(debtSequence.Balances) > i {
				balance += debtSequence.Balances[i]
			}
		}

		balances = append(balances, balance)
	}

	return balances
}
