//go:build wasm
// +build wasm

package main

import (
	"sync"
	"syscall/js"

	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

func main() {
	wait := make(chan struct{}, 0)
	js.Global().Set("calculateAll", js.FuncOf(calculateAll))
	<-wait
}

func calculateAll(this js.Value, args []js.Value) interface{} {
	debts := []calculator.Debt{}

	for i := range args[1].Int() {
		debts = append(debts, calculator.Debt{
			Name:           args[0].Index(i).Get("name").String(),
			Amount:         args[0].Index(i).Get("amount").Float(),
			MinimumPayment: args[0].Index(i).Get("minimum_payment").Float(),
			AnnualInterest: args[0].Index(i).Get("interest").Float(),
		})
	}

	model := calculator.NewModel(calculator.Input{
		ExtraMonthlyPayment:     args[2].Float(),
		OneTimeImmediatePayment: args[3].Float(),
		Debts:                   debts,
	})

	wg := &sync.WaitGroup{}
	ch := make(chan calculator.CalculationData, len(calculations))
	for datakey, calculation := range calculations {
		wg.Add(1)
		go calculator.CalculateAsyncWasm(wg, ch, datakey, calculation, model)
	}
	wg.Wait()

	modelMap := map[string]interface{}{}

	for len(ch) > 0 {
		calculationData := <-ch
		modelMap[calculationData.Datakey] = calculationData.Value
	}

	close(ch)

	return js.ValueOf(modelMap)
}

var calculations = map[string]any{
	"DEBT_PAYOFF_MONTH":         calculator.NewDebtPayoffMonth(),
	"MONTHLY_SEQUENCE_BALANCES": calculator.NewMonthlySequenceBalances(),
	"MONTHLY_SEQUENCE_PAYMENTS": calculator.NewMonthlySequencePayments(),
	"SNOWBALL":                  calculator.NewSnowball(),
	"TOTAL_BEGINNING_DEBT":      calculator.NewTotalBeginningDebt(),
	"TOTAL_INTEREST":            calculator.NewTotalInterest(),
	"TOTAL_MINIMUM_PAYMENT":     calculator.NewTotalMinimumPayment(),
	"TOTAL_PAYMENTS":            calculator.NewTotalPayments(),
}
