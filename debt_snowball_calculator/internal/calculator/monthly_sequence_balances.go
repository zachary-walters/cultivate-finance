package calculator

type MonthlySequenceBalancesCalculation SequenceCalculation

type MonthlySequenceBalances struct {
	DebtPayoffMonthCalculation
	SnowballCalculation
	TotalBeginningDebtCalculation
}

func NewMonthlySequenceBalances() *MonthlySequenceBalances {
	return &MonthlySequenceBalances{
		DebtPayoffMonthCalculation:    NewDebtPayoffMonth(),
		SnowballCalculation:           NewSnowball(),
		TotalBeginningDebtCalculation: NewTotalBeginningDebt(),
	}
}

func (c *MonthlySequenceBalances) Calculate(model *Model) []float64 {
	debtPayoffMonth := c.DebtPayoffMonthCalculation.Calculate(model)
	snowball := c.SnowballCalculation.Calculate(model)
	totalBeginningDebt := c.TotalBeginningDebtCalculation.Calculate(model)

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
