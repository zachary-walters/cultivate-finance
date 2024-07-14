package calculator

type TotalMinimumPaymentCalculation Calculation

type TotalMinimumPayment struct {
	AbstractCalculation
}

func NewTotalMinimumPayment() *TotalMinimumPayment {
	return &TotalMinimumPayment{}
}

func (c TotalMinimumPayment) CalculateSnowball(model *Model) float64 {
	total := 0.0

	for _, debt := range model.Input.Debts {
		total += debt.MinimumPayment
	}

	return c.SanitizeToZero(total)
}

func (c TotalMinimumPayment) CalculateAvalanche(model *Model) float64 {
	return c.CalculateSnowball(model)
}
