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

func (c MonthSequence) Calculate(model *Model) []float64 {
	debtPayoffMonth := c.DebtPayoffMonthCalculation.Calculate(model)
	sequence := []float64{}

	for i := 1; i <= int(debtPayoffMonth)+2; i++ {
		sequence = append(sequence, float64(i))
	}

	return sequence
}
