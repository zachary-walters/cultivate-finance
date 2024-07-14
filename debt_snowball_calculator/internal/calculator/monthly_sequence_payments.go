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

func (c MonthlySequencePayments) CalculateSnowball(model *Model) []float64 {
	debtPayoffMonth := c.DebtPayoffMonthCalculation.CalculateSnowball(model)
	snowball := c.SnowballCalculation.CalculateSnowball(model)

	payments := []float64{}

	for i := 0; i < int(debtPayoffMonth); i++ {
		payment := 0.0
		for _, debtSequence := range snowball {
			if len(debtSequence.Payments) > i && debtSequence.IsValid() {
				payment += debtSequence.Payments[i]
			}
		}

		payments = append(payments, payment)
	}

	return payments
}

func (c MonthlySequencePayments) CalculateAvalanche(model *Model) []float64 {
	debtPayoffMonth := c.DebtPayoffMonthCalculation.CalculateAvalanche(model)
	snowball := c.SnowballCalculation.CalculateAvalanche(model)

	payments := []float64{}

	for i := 0; i < int(debtPayoffMonth); i++ {
		payment := 0.0
		for _, debtSequence := range snowball {
			if len(debtSequence.Payments) > i && debtSequence.IsValid() {
				payment += debtSequence.Payments[i]
			}
		}

		payments = append(payments, payment)
	}

	return payments
}
