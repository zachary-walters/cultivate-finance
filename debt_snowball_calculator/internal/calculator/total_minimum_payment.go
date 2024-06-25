package calculator

type TotalMinimumPaymentCalculation Calculation

type TotalMinimumPayment struct {
	AbstractCalculation
}

func NewTotalMinimumPayment() *TotalMinimumPayment {
	return &TotalMinimumPayment{}
}

func (c TotalMinimumPayment) Calculate(model *Model) float64 {
	total := 0.0

	for _, debt := range model.Input.Debts {
		total += debt.MinimumPayment
	}

	return c.SanitizeToZero(total)
}
