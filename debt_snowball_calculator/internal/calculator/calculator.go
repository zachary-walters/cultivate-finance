package calculator

type Calculation interface {
	Calculate(*Model) float64
}

type SequenceCalculation interface {
	Calculate(*Model) []float64
}

type AbstractCalculation struct{}

func (c *AbstractCalculation) SanitizeToZero(i interface{}) float64 {
	switch v := i.(type) {
	case float64:
		if v < 0.0 {
			return 0.0
		}
		return v
	case int:
		if v < 0 {
			return 0
		}
		return float64(v)
	default:
		return 0
	}
}

type Debt struct {
	Name           string  `json:"name"`
	Amount         float64 `json:"amount"`
	AnnualInterest float64 `json:"annual_interest"`
	MinimumPayment float64 `json:"minimum_payment"`
}

type Input struct {
	ExtraMonthlyPayment     float64 `json:"extra_monthly_payment"`
	OneTimeImmediatePayment float64 `json:"one_time_immediate_payment"`
	Debts                   []Debt  `json:"debts"`
	Datakey                 string  `json:"datakey"`
}

type Model struct {
	Input Input
}

func NewModel(input Input) *Model {
	return &Model{
		Input: input,
	}
}

type DebtSequence struct {
	Debt     Debt
	Months   []float64
	Payments []float64
	Balances []float64
}

type DebtSequences []DebtSequence
