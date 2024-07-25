package calculator

import (
	"sync"
)

type Calculation interface {
	CalculateSnowball(Model) float64
	CalculateAvalanche(Model) float64
}

type SequenceCalculation interface {
	CalculateSnowball(Model) []float64
	CalculateAvalanche(Model) []float64
}

type SnowballCalculation interface {
	CalculateSnowball(Model) DebtSequences
	CalculateAvalanche(Model) DebtSequences
}

type DebtCalculation interface {
	CalculateSnowball(Model) []Debt
	CalculateAvalanche(Model) []Debt
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
	ID             string  `json:"id"`
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

func NewModel(input Input) Model {
	return Model{
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

type FinalDecision struct {
	Choice                  string  `json:"choice"`
	MonthDifference         float64 `json:"month_difference"`
	TotalPaymentsDifference float64 `json:"payments_difference"`
}

type CalculationData struct {
	Datakey   string `json:"datakey,omitempty"`
	Value     any    `json:"value,omitempty"`
	Avalanche any    `json:"avalanche,omitempty"`
}

func CalculateSynchronous(model Model, calculation any, datakey string) CalculationData {
	calc, isCalculation := calculation.(Calculation)
	seq, isSequenceCalculation := calculation.(SequenceCalculation)
	snowball, isSnowballCalculation := calculation.(SnowballCalculation)
	decisionCalculation, isDecisionCalculation := calculation.(DecisionCalculation)
	debtCalculation, isDebtCalculation := calculation.(DebtCalculation)

	calculationData := CalculationData{
		Datakey: datakey,
	}

	if isSequenceCalculation {
		calculationData.Value = seq.CalculateSnowball(model)
		calculationData.Avalanche = seq.CalculateAvalanche(model)
	} else if isCalculation {
		calculationData.Value = calc.CalculateSnowball(model)
		calculationData.Avalanche = calc.CalculateAvalanche(model)
	} else if isSnowballCalculation {
		calculationData.Value = snowball.CalculateSnowball(model)
		calculationData.Avalanche = snowball.CalculateAvalanche(model)
	} else if isDecisionCalculation {
		calculationData.Value = decisionCalculation.CalculateSnowball(model)
		calculationData.Avalanche = nil
	} else if isDebtCalculation {
		calculationData.Value = debtCalculation.CalculateSnowball(model)
		calculationData.Avalanche = nil
	}

	return calculationData
}

func CalculateAsync(wg *sync.WaitGroup, ch chan CalculationData, datakey string, calculation any, model Model) {
	defer wg.Done()
	calculationData := CalculateSynchronous(model, calculation, datakey)

	ch <- calculationData
}

func CalculateSynchronousWasm(model Model, calculation any, datakey string) CalculationData {
	calc, isCalculation := calculation.(Calculation)
	seq, isSequenceCalculation := calculation.(SequenceCalculation)
	snowball, isSnowballCalculation := calculation.(SnowballCalculation)
	decisionCalculation, isDecisionCalculation := calculation.(DecisionCalculation)
	debtCalculation, isDebtCalculation := calculation.(DebtCalculation)

	calculationData := CalculationData{
		Datakey: datakey,
	}

	if isSequenceCalculation {
		calculationData.Value = TranslateFloatSlice(seq.CalculateSnowball(model))
		calculationData.Avalanche = TranslateFloatSlice(seq.CalculateAvalanche(model))
	} else if isCalculation {
		calculationData.Value = calc.CalculateSnowball(model)
		calculationData.Avalanche = calc.CalculateAvalanche(model)
	} else if isSnowballCalculation {
		calculationData.Value = TranslateSnowball(snowball.CalculateSnowball(model))
		calculationData.Avalanche = TranslateSnowball(snowball.CalculateAvalanche(model))
	} else if isDecisionCalculation {
		calculationData.Value = TranslateDecision(decisionCalculation.CalculateSnowball(model))
		calculationData.Avalanche = nil
	} else if isDebtCalculation {
		calculationData.Value = TranslateDebts(debtCalculation.CalculateSnowball(model))
		calculationData.Avalanche = nil
	}

	return calculationData
}

func CalculateAsyncWasm(wg *sync.WaitGroup, ch chan CalculationData, datakey string, calculation any, model Model) {
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

func TranslateSnowball(s DebtSequences) []interface{} {
	x := []interface{}{}

	for idx, debtSequence := range s {
		x = append(x, map[string]interface{}{
			"sequence": idx,
			"balances": TranslateFloatSlice(debtSequence.Balances),
			"payments": TranslateFloatSlice(debtSequence.Payments),
			"months":   TranslateFloatSlice(debtSequence.Months),
			"debt": map[string]interface{}{
				"id":              debtSequence.Debt.ID,
				"name":            debtSequence.Debt.Name,
				"amount":          debtSequence.Debt.Amount,
				"minimum_payment": debtSequence.Debt.MinimumPayment,
				"annual_interest": debtSequence.Debt.AnnualInterest,
			},
		})
	}

	return x
}

func TranslateDebts(debts []Debt) []interface{} {
	x := []interface{}{}

	for _, debt := range debts {
		x = append(x, map[string]interface{}{
			"id":              debt.ID,
			"name":            debt.Name,
			"amount":          debt.Amount,
			"minimum_payment": debt.MinimumPayment,
			"annual_interest": debt.AnnualInterest,
		})
	}

	return x
}

func TranslateDecision(decision FinalDecision) map[string]interface{} {
	return map[string]interface{}{
		"choice":                    decision.Choice,
		"month_difference":          decision.MonthDifference,
		"total_payments_difference": decision.TotalPaymentsDifference,
	}
}
