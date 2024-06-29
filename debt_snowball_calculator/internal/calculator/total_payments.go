package calculator

type TotalPaymentsCalculation Calculation

type TotalPayments struct {
	AbstractCalculation
	MonthlySequencePaymentsCalculation
}

func NewTotalPayments() *TotalPayments {
	return &TotalPayments{
		MonthlySequencePaymentsCalculation: NewMonthlySequencePayments(),
	}
}

func (c *TotalPayments) Calculate(model *Model) float64 {
	monthlySequencePayments := c.MonthlySequencePaymentsCalculation.Calculate(model)

	total := 0.0
	for _, payment := range monthlySequencePayments {
		total += payment
	}

	return c.SanitizeToZero(total)
}
