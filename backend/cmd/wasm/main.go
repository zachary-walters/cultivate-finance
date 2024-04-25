package main

import (
	"embed"
	"encoding/json"
	"sync"
	"syscall/js"

	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

//go:embed constants.json
var res embed.FS

var constants calculator.Model

type data struct {
	Value           any `json:"value,omitempty"`
	RetirementValue any `json:"retirement_value,omitempty"`
}

func main() {
	j, err := res.ReadFile("constants.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(j, &constants)

	wait := make(chan struct{}, 0)
	js.Global().Set("calculate", js.FuncOf(calculate))
	js.Global().Set("calculateAll", js.FuncOf(calculateAll))
	<-wait
}

func calculateAll(this js.Value, args []js.Value) interface{} {
	model := calculator.Model{
		Input: calculator.Input{
			CurrentAge:                args[0].Get("current_age").Int(),
			CurrentFilingStatus:       args[0].Get("current_filing_status").String(),
			CurrentAnnualIncome:       args[0].Get("current_annual_income").Float(),
			AnnualContributionsPreTax: args[0].Get("annual_contributions_pre_tax").Float(),
			AnnualInvestmentGrowth:    args[0].Get("annual_investment_growth").Float(),
			RetirementAge:             args[0].Get("retirement_age").Int(),
			RetirementFilingStatus:    args[0].Get("retirement_filing_status").String(),
			YearlyWithdrawal:          args[0].Get("yearly_withdrawal").Float(),
		},
		SingleTaxRates:                       constants.SingleTaxRates,
		MarriedJointTaxRates:                 constants.MarriedJointTaxRates,
		MarriedSeperateTaxRates:              constants.MarriedSeperateTaxRates,
		HeadOfHouseholdTaxRates:              constants.HeadOfHouseholdTaxRates,
		STANDARD_DEDUCTION_SINGLE:            constants.STANDARD_DEDUCTION_SINGLE,
		STANDARD_DEDUCTION_MARRIED_JOINT:     constants.STANDARD_DEDUCTION_MARRIED_JOINT,
		STANDARD_DEDUCTION_MARRIED_SEPERATE:  constants.STANDARD_DEDUCTION_MARRIED_SEPERATE,
		STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD: constants.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD,
	}

	defaultCh := make(chan map[string]any, len(calculator.Calculations))
	retirementCh := make(chan map[string]any, len(calculator.Calculations))
	wg := &sync.WaitGroup{}
	for datakey, calculation := range calculator.Calculations {
		wg.Add(1)
		go calculator.CalculateAsyncWasm(wg, defaultCh, retirementCh, datakey, calculation, model)
	}
	wg.Wait()

	close(defaultCh)
	close(retirementCh)

	m := map[string]any{}
	rm := map[string]any{}

	for i := 0; i < len(calculator.Calculations); i++ {
		select {
		case data := <-defaultCh:
			for datakey, value := range data {
				m[datakey] = value
			}
		}
	}

	for i := 0; i < len(calculator.Calculations); i++ {
		select {
		case data := <-retirementCh:
			for datakey, value := range data {
				rm[datakey] = value
			}
		}
	}

	modelMapDefault := map[string]interface{}{}
	modelMapRetired := map[string]interface{}{}
	for datakey := range m {
		modelMapDefault[datakey] = m[datakey]
		modelMapRetired[datakey] = rm[datakey]
	}

	return js.ValueOf(map[string]interface{}{
		"default": modelMapDefault,
		"retired": modelMapRetired,
	})
}

func calculate(this js.Value, args []js.Value) interface{} {
	model := calculator.Model{
		Input: calculator.Input{
			CurrentAge:                args[0].Get("current_age").Int(),
			CurrentFilingStatus:       args[0].Get("current_filing_status").String(),
			CurrentAnnualIncome:       args[0].Get("current_annual_income").Float(),
			AnnualContributionsPreTax: args[0].Get("annual_contributions_pre_tax").Float(),
			AnnualInvestmentGrowth:    args[0].Get("annual_investment_growth").Float(),
			RetirementAge:             args[0].Get("retirement_age").Int(),
			RetirementFilingStatus:    args[0].Get("retirement_filing_status").String(),
			YearlyWithdrawal:          args[0].Get("yearly_withdrawal").Float(),
		},
		SingleTaxRates:                       constants.SingleTaxRates,
		MarriedJointTaxRates:                 constants.MarriedJointTaxRates,
		MarriedSeperateTaxRates:              constants.MarriedSeperateTaxRates,
		HeadOfHouseholdTaxRates:              constants.HeadOfHouseholdTaxRates,
		STANDARD_DEDUCTION_SINGLE:            constants.STANDARD_DEDUCTION_SINGLE,
		STANDARD_DEDUCTION_MARRIED_JOINT:     constants.STANDARD_DEDUCTION_MARRIED_JOINT,
		STANDARD_DEDUCTION_MARRIED_SEPERATE:  constants.STANDARD_DEDUCTION_MARRIED_SEPERATE,
		STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD: constants.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD,
	}

	calculation, _ := calculator.Calculations[args[0].Get("datakey").String()]

	value, retirementValue := calculator.CalculateSynchronousWasm(model, calculation)

	return js.ValueOf(map[string]interface{}{
		"datakey":          args[0].Get("datakey").String(),
		"value":            value,
		"retirement_value": retirementValue,
	})
}
