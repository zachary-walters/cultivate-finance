package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

//go:embed constants.json
var res embed.FS

func main() {
	var model calculator.Model
	j, err := res.ReadFile("constants.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(j, &model)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/{datakey}", func(w http.ResponseWriter, r *http.Request) {
		datakey := strings.ToUpper(chi.URLParam(r, "datakey"))

		calculation, exists := calculations[datakey]
		if !exists {
			w.Write([]byte(fmt.Sprintf("the given datakey does not exist: %s", datakey)))
			return
		}

		decoder := json.NewDecoder(r.Body)
		var input calculator.Input

		err := decoder.Decode(&input)
		if err != nil {
			panic(err)
		}

		model.Input = input

		calc, isCalculation := calculation.(calculator.Calculation)
		seq, isSequenceCalculation := calculation.(calculator.SequenceCalculation)
		chart, isChartCalculation := calculation.(calculator.ChartCalculation)

		var value any
		var retirementValue any

		if isCalculation {
			value = calc.Calculate(model)
			retirementValue = calc.CalculateRetirement(model)
		} else if isSequenceCalculation {
			value = seq.Calculate(model)
			retirementValue = seq.CalculateRetirement(model)
		} else if isChartCalculation {
			value = chart.Calculate(model)
		} else {
			// return an error
		}

		// value := calculation.Calculate(model)
		// retirementValue := calculation.CalculateRetirement(model)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(struct {
			Datakey         string `json:"datakey"`
			Value           any    `json:"value,omitempty"`
			RetirementValue any    `json:"retirement_value,omitempty"`
		}{
			Datakey:         datakey,
			Value:           value,
			RetirementValue: retirementValue,
		})

		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8660", r)
}

var calculations = map[string]any{
	"ANNUAL_GROWTH_LESS_INFLATION":                                             calculator.NewAnnualGrowthLessInflation(),
	"ANNUAL_TAX_SAVINGS_WITH_CONTRIBUTION":                                     calculator.NewAnnualTaxSavingsWithContribution(),
	"BALANCES_ROTH_MATCHING_GROSS_CONTRIBUTIONS":                               calculator.NewBalancesRothMatchingGrossContributions(),
	"BALANCES_ROTH_MATCHING_NET_CONTRIBUTIONS":                                 calculator.NewBalancesRothMatchingNetContributions(),
	"BALANCES_TRADITIONAL":                                                     calculator.NewBalancesTraditional(),
	"EFFECTIVE_TAX_RATE_ON_GROSS":                                              calculator.NewEffectiveTaxRateOnGross(),
	"EQUIVALENT_ROTH_CONTRIBUTIONS":                                            calculator.NewEquivalentRothContributions(),
	"INCOME_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS":                        calculator.NewIncomeAfterStandardDeductionAndContributions(),
	"INCOME_AFTER_STANDARD_DEDUCTION":                                          calculator.NewIncomeAfterStandardDeduction(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_SINGLE":     calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsSingle(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_SINGLE":                       calculator.NewIncomePerBracketAfterStandardDeductionSingle(),
	"NET_DISTRIBUTION_AFTER_TAXES":                                             calculator.NewNetDistributionAfterTaxes(),
	"STANDARD_DEDUCTION":                                                       calculator.NewStandardDeduction(),
	"TAX_RATE_OF_SAVINGS":                                                      calculator.NewTaxRateOfSavings(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_SINGLE": calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_SINGLE":                   calculator.NewTaxesOwedPerBracketAfterStandardDeductionSingle(),
	"TOTAL_DISBURSEMENTS_AFTER_TAX":                                            calculator.NewTotalDisbursementsAfterTax(),
	"TOTAL_DISBURSEMENTS_ROTH_MATCHING_GROSS":                                  calculator.NewTotalDisbursementsRothMatchingGross(),
	"TOTAL_DISBURSEMENTS_ROTH_MATCHING_NET":                                    calculator.NewTotalDisbursementsRothMatchingNet(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_SINGLE":       calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsSingle(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS":              calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributions(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_SINGLE":                         calculator.NewTotalTaxesOwedAfterStandardDeductionSingle(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION":                                calculator.NewTotalTaxesOwedAfterStandardDeduction(),
}
