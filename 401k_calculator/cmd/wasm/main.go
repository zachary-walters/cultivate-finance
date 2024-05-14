//go:build wasm
// +build wasm

package main

import (
	"sync"
	"syscall/js"

	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type data struct {
	Value           any `json:"value,omitempty"`
	RetirementValue any `json:"retirement_value,omitempty"`
}

func main() {
	wait := make(chan struct{}, 0)
	js.Global().Set("calculateAll", js.FuncOf(calculateAll))
	<-wait
}

func calculateAll(this js.Value, args []js.Value) interface{} {
	input := calculator.Input{
		CurrentAge:                args[0].Get("current_age").Int(),
		CurrentFilingStatus:       args[0].Get("current_filing_status").String(),
		CurrentAnnualIncome:       args[0].Get("current_annual_income").Float(),
		AnnualContributionsPreTax: args[0].Get("annual_contributions_pre_tax").Float(),
		AnnualInvestmentGrowth:    args[0].Get("annual_investment_growth").Float(),
		RetirementAge:             args[0].Get("retirement_age").Int(),
		RetirementFilingStatus:    args[0].Get("retirement_filing_status").String(),
		YearlyWithdrawal:          args[0].Get("yearly_withdrawal").Float(),
	}

	model := calculator.NewModel(input)

	wg := &sync.WaitGroup{}
	ch := make(chan calculator.CalculationData, len(calculations))
	for datakey, calculation := range calculations {
		wg.Add(1)
		go calculator.CalculateAsyncWasm(wg, ch, datakey, calculation, model)
	}
	wg.Wait()

	modelMapTraditional := map[string]interface{}{}
	modelMapTraditionalRetired := map[string]interface{}{}
	modelMapRoth := map[string]interface{}{}
	modelMapRothRetired := map[string]interface{}{}

	for len(ch) > 0 {
		calculationData := <-ch
		modelMapTraditional[calculationData.Datakey] = calculationData.TraditionalValue
		modelMapTraditionalRetired[calculationData.Datakey] = calculationData.TraditionalRetirementValue
		modelMapRoth[calculationData.Datakey] = calculationData.RothValue
		modelMapRothRetired[calculationData.Datakey] = calculationData.RothRetirementValue
	}

	return js.ValueOf(map[string]interface{}{
		"traditional":            modelMapTraditional,
		"traditional_retirement": modelMapTraditionalRetired,
		"roth":                   modelMapRoth,
		"roth_retirement":        modelMapRothRetired,
	})
}

var calculations = map[string]any{
	"ANNUAL_GROWTH_LESS_INFLATION":                      calculator.NewAnnualGrowthLessInflation(),
	"ANNUAL_TAX_SAVINGS_WITH_CONTRIBUTION":              calculator.NewAnnualTaxSavingsWithContribution(),
	"BALANCES_ROTH_MATCHING_NET_CONTRIBUTIONS":          calculator.NewBalancesRothMatchingNetContributions(),
	"BALANCES_TRADITIONAL":                              calculator.NewBalancesTraditional(),
	"EFFECTIVE_TAX_RATE_ON_GROSS":                       calculator.NewEffectiveTaxRateOnGross(),
	"EQUIVALENT_ROTH_CONTRIBUTIONS":                     calculator.NewEquivalentRothContributions(),
	"INCOME_AFTER_STANDARD_DEDUCTION":                   calculator.NewIncomeAfterStandardDeduction(),
	"INCOME_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS": calculator.NewIncomeAfterStandardDeductionAndContributions(),
	"NET_DISTRIBUTION_AFTER_TAXES":                      calculator.NewNetDistributionAfterTaxes(),
	"STANDARD_DEDUCTION":                                calculator.NewStandardDeduction(),
	"TAX_RATE_OF_SAVINGS":                               calculator.NewTaxRateOfSavings(),
	"TOTAL_CONTRIBUTIONS":                               calculator.NewTotalContributions(),
	"TOTAL_DISBURSEMENTS":                               calculator.NewTotalDisbursements(),
	"TOTAL_INTEREST":                                    calculator.NewTotalInterest(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION":         calculator.NewTotalTaxesOwedAfterStandardDeduction(),
}
