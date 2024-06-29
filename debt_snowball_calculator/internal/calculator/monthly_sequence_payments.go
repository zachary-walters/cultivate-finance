package calculator

type MonthlySequencePaymentsCalculation SequenceCalculation

type MonthlySequencePayments struct {
	DebtPayoffMonthCalculation
	SnowballCalculation
}

func NewMonthlySequencePayments() *MonthlySequencePayments {
	return &MonthlySequencePayments{
		DebtPayoffMonthCalculation: NewDebtPayoffMonth(),
		SnowballCalculation:        NewSnowball(),
	}
}

func (c MonthlySequencePayments) Calculate(model *Model) []float64 {
	debtPayoffMonth := c.DebtPayoffMonthCalculation.Calculate(model)
	snowball := c.SnowballCalculation.Calculate(model)

	payments := []float64{}

	for i := 0; i < int(debtPayoffMonth); i++ {
		payment := 0.0
		for _, debtSequence := range snowball {
			if len(debtSequence.Payments) > i {
				payment += debtSequence.Payments[i]
			}
		}

		payments = append(payments, payment)
	}

	return payments
}
