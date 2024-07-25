package calculator

type DecisionCalculation interface {
	CalculateSnowball(Model) FinalDecision
	CalculateAvalanche(Model) FinalDecision
}

type Decision struct {
	DebtPayoffMonthCalculation
	TotalPaymentsCalculation
}

func NewDecision() *Decision {
	return &Decision{
		DebtPayoffMonthCalculation: NewDebtPayoffMonth(),
		TotalPaymentsCalculation:   NewTotalPayments(),
	}
}

func (c *Decision) CalculateSnowball(model Model) FinalDecision {
	debtPayoffMonthSnowball := c.DebtPayoffMonthCalculation.CalculateSnowball(model)
	debtPayoffMonthAvalanche := c.DebtPayoffMonthCalculation.CalculateAvalanche(model)

	totalPaymentsSnowball := c.TotalPaymentsCalculation.CalculateSnowball(model)
	totalPaymentsAvalanche := c.TotalPaymentsCalculation.CalculateAvalanche(model)
	totalPaymentsDifference := totalPaymentsSnowball - totalPaymentsAvalanche

	choice := "Avalanche"
	if totalPaymentsDifference == 0 {
		choice = "Either"
	} else if totalPaymentsDifference < 0 {
		choice = "Snowball"
	}

	return FinalDecision{
		Choice:                  choice,
		MonthDifference:         debtPayoffMonthSnowball - debtPayoffMonthAvalanche,
		TotalPaymentsDifference: totalPaymentsSnowball - totalPaymentsAvalanche,
	}
}

func (c *Decision) CalculateAvalanche(model Model) FinalDecision {
	return c.CalculateSnowball(model)
}
