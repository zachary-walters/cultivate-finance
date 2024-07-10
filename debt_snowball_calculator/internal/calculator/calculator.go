package calculator

import (
	"sync"
)

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
	Debt     Debt      `json:"debt,omitempty"`
	Months   []float64 `json:"months,omitempty"`
	Payments []float64 `json:"payments,omitempty"`
	Balances []float64 `json:"balances,omitempty"`
}

type DebtSequences []DebtSequence

type CalculationData struct {
	Datakey string `json:"datakey,omitempty"`
	Value   any    `json:"value,omitempty"`
}

func CalculateSynchronous(model *Model, calculation any, datakey string) CalculationData {
	calc, isCalculation := calculation.(Calculation)
	seq, isSequenceCalculation := calculation.(SequenceCalculation)
	snowball, isSnowballCalculation := calculation.(SnowballCalculation)

	calculationData := CalculationData{
		Datakey: datakey,
	}

	if isSequenceCalculation {
		calculationData.Value = seq.Calculate(model)
	} else if isCalculation {
		calculationData.Value = calc.Calculate(model)
	} else if isSnowballCalculation {
		calculationData.Value = snowball.Calculate(model)
	}

	return calculationData
}

func CalculateAsync(wg *sync.WaitGroup, ch chan CalculationData, datakey string, calculation any, model *Model) {
	defer wg.Done()
	calculationData := CalculateSynchronous(model, calculation, datakey)

	ch <- calculationData
}

func CalculateSynchronousWasm(model *Model, calculation any, datakey string) CalculationData {
	calc, isCalculation := calculation.(Calculation)
	seq, isSequenceCalculation := calculation.(SequenceCalculation)
	snowball, isSnowballCalculation := calculation.(SnowballCalculation)

	calculationData := CalculationData{
		Datakey: datakey,
	}

	if isSequenceCalculation {
		calculationData.Value = TranslateFloatSlice(seq.Calculate(model))
	} else if isCalculation {
		calculationData.Value = calc.Calculate(model)
	} else if isSnowballCalculation {
		calculationData.Value = TranslateSnowball(snowball.Calculate(model))
	}

	return calculationData
}

func CalculateAsyncWasm(wg *sync.WaitGroup, ch chan CalculationData, datakey string, calculation any, model *Model) {
	defer wg.Done()
	calculationData := CalculateSynchronousWasm(model, calculation, datakey)

	ch <- calculationData
}

func TranslateFloatSlice(s []float64) []interface{} {
	x := make([]interface{}, len(s))

	for i := range s {
		x[i] = s[i]
	}

	return x
}

func TranslateSnowball(s DebtSequences) map[string]interface{} {
	m := map[string]interface{}{}

	for _, debtSequence := range s {
		m[debtSequence.Debt.Name] = map[string]interface{}{
			"balances": TranslateFloatSlice(debtSequence.Balances),
			"payments": TranslateFloatSlice(debtSequence.Payments),
			"months":   TranslateFloatSlice(debtSequence.Months),
			"debt": map[string]interface{}{
				debtSequence.Debt.Name: map[string]interface{}{
					"name":            debtSequence.Debt.Name,
					"amount":          debtSequence.Debt.Amount,
					"minimum_payment": debtSequence.Debt.MinimumPayment,
					"annual_interest": debtSequence.Debt.AnnualInterest,
				},
			},
		}
	}

	return m
}
