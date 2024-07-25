package calculator

type MonthSequenceCalculation Calculation

type MonthSequence struct {
	DebtPayoffMonthCalculation
}

func NewMonthSequence() *MonthSequence {
	return &MonthSequence{
		DebtPayoffMonthCalculation: NewDebtPayoffMonth(),
	}
}

func (c MonthSequence) CalculateSnowball(model Model) []float64 {
	debtPayoffMonth := c.DebtPayoffMonthCalculation.CalculateSnowball(model)
	sequence := []float64{}

	for i := 1; i <= int(debtPayoffMonth)+2; i++ {
		sequence = append(sequence, float64(i))
	}

	return sequence
}

func (c MonthSequence) CalculateAvalanche(model Model) []float64 {
	debtPayoffMonth := c.DebtPayoffMonthCalculation.CalculateAvalanche(model)
	sequence := []float64{}

	for i := 1; i <= int(debtPayoffMonth)+2; i++ {
		sequence = append(sequence, float64(i))
	}

	return sequence
}
