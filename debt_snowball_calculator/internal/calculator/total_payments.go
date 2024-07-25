package calculator

type TotalPaymentsCalculation Calculation

type TotalPayments struct {
	AbstractCalculation
	MonthlySequencePaymentsCalculation
	ValidDebtsCalculation
}

func NewTotalPayments() *TotalPayments {
	return &TotalPayments{
		MonthlySequencePaymentsCalculation: NewMonthlySequencePayments(),
	}
}

func (c *TotalPayments) CalculateSnowball(model Model) float64 {
	monthlySequencePayments := c.MonthlySequencePaymentsCalculation.CalculateSnowball(model)

	total := 0.0
	for _, payment := range monthlySequencePayments {
		total += payment
	}

	return c.SanitizeToZero(total)
}

func (c *TotalPayments) CalculateAvalanche(model Model) float64 {
	monthlySequencePayments := c.MonthlySequencePaymentsCalculation.CalculateAvalanche(model)

	total := 0.0
	for _, payment := range monthlySequencePayments {
		total += payment
	}

	return c.SanitizeToZero(total)
}
