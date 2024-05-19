package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"

	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type data struct {
	TraditionalValue           any `json:"traditional_value,omitempty"`
	TraditionalRetirementValue any `json:"traditional_retirement_value,omitempty"`
	RothValue                  any `json:"roth_value,omitempty"`
	RothRetirementValue        any `json:"roth_retirement_value,omitempty"`
}

func main() {
	var nc *nats.Conn
	var err error
	r := chi.NewRouter()

	uri := os.Getenv("NATS_URI")
	for i := 0; i < 5; i++ {
		nc, err = nats.Connect(uri)
		if err == nil {
			break
		}

		log.Println("Waiting before connecting to NATS at:", uri)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatal("Error establishing connection to NATS:", err)
	}
	log.Println("Connected to NATS at:", nc.ConnectedUrl())

	nc.Subscribe("calculate_all_401k", func(msg *nats.Msg) {
		log.Println("Got task request on:", msg.Subject)
		model, err := calculateModel(msg.Data)
		if err != nil {

		}

		d, err := json.Marshal(model)
		if err != nil {

		}

		nc.Publish(msg.Reply, d)
	})

	nc.Subscribe("calculate_401k_by_datakey", func(msg *nats.Msg) {
		log.Println("Got task request on:", msg.Subject)

		calculationData, err := calculateDatakey(msg.Data)
		if err != nil {

		}

		d, err := json.Marshal(calculationData)
		if err != nil {

		}

		nc.Publish(msg.Reply, d)

	})

	// r.Get("/", ns.ping)
	// r.Post("/{datakey}", ns.calculateByDatakey)
	// r.Post("/calculate_all", ns.calculateAll)

	log.Println("Server listening on port: ", os.Getenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r); err != nil {
		log.Fatal(err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func calculateModel(d []byte) (map[string]data, error) {
	decoder := json.NewDecoder(bytes.NewReader(d))
	var input calculator.Input

	err := decoder.Decode(&input)
	if err != nil {
		return nil, err
	}

	model := calculator.NewModel(input)

	wg := &sync.WaitGroup{}
	ch := make(chan calculator.CalculationData, len(calculations))
	for datakey, calculation := range calculations {
		wg.Add(1)
		go calculator.CalculateAsync(wg, ch, datakey, calculation, model)
	}
	wg.Wait()

	close(ch)

	modelMap := map[string]data{}
	for len(ch) > 0 {
		calculationData := <-ch
		modelMap[calculationData.Datakey] = data{
			TraditionalValue:           calculationData.TraditionalValue,
			TraditionalRetirementValue: calculationData.TraditionalRetirementValue,
			RothValue:                  calculationData.RothValue,
			RothRetirementValue:        calculationData.RothRetirementValue,
		}
	}

	return modelMap, nil
}

func calculateDatakey(d []byte) (any, error) {
	decoder := json.NewDecoder(bytes.NewReader(d))
	var input calculator.Input

	err := decoder.Decode(&input)
	if err != nil {
		return nil, err
	}

	log.Println("ARST", input, "ZXCV", input.Datakey)
	calculation, exists := calculations[input.Datakey]
	if !exists {
		// send no exists message reply to nats
		log.Println("calculation doesn't exist")
	}

	model := calculator.NewModel(input)
	calculationData := calculator.CalculateSynchronous(model, calculation, input.Datakey)

	data := struct {
		Datakey                    string `json:"datakey"`
		TraditionalValue           any    `json:"traditional_value,omitempty"`
		TraditionalRetirementValue any    `json:"traditional_retirement_value,omitempty"`
		RothValue                  any    `json:"roth_value,omitempty"`
		RothRetirementValue        any    `json:"roth_retirement_value,omitempty"`
	}{
		Datakey:                    input.Datakey,
		TraditionalValue:           calculationData.TraditionalValue,
		TraditionalRetirementValue: calculationData.TraditionalRetirementValue,
		RothValue:                  calculationData.RothValue,
		RothRetirementValue:        calculationData.RothRetirementValue,
	}

	return data, nil
}

// func calculateByDatakey(w http.ResponseWriter, r *http.Request) {
// 	datakey := strings.ToUpper(chi.URLParam(r, "datakey"))

// 	calculation, exists := calculations[datakey]
// 	if !exists {
// 		w.Write([]byte(fmt.Sprintf("the given datakey does not exist: %s", datakey)))
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	var input calculator.Input

// 	err := decoder.Decode(&input)
// 	if err != nil {
// 		panic(err)
// 	}

// 	model := calculator.NewModel(input)

// 	calculationData := calculator.CalculateSynchronous(model, calculation, datakey)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(struct {
// 		Datakey                    string `json:"datakey"`
// 		TraditionalValue           any    `json:"traditional_value,omitempty"`
// 		TraditionalRetirementValue any    `json:"traditional_retirement_value,omitempty"`
// 		RothValue                  any    `json:"roth_value,omitempty"`
// 		RothRetirementValue        any    `json:"roth_retirement_value,omitempty"`
// 	}{
// 		Datakey:                    datakey,
// 		TraditionalValue:           calculationData.TraditionalValue,
// 		TraditionalRetirementValue: calculationData.TraditionalRetirementValue,
// 		RothValue:                  calculationData.RothValue,
// 		RothRetirementValue:        calculationData.RothRetirementValue,
// 	})

// 	if err != nil {
// 		panic(err)
// 	}
// }

var calculations = map[string]any{
	"ADJUSTED_GROSS_INCOME":                                                               calculator.NewAdjustedGrossIncome(),
	"ANNUAL_GROWTH_LESS_INFLATION":                                                        calculator.NewAnnualGrowthLessInflation(),
	"ANNUAL_RETIREMENT_ACCOUNT_DISBURSEMENT":                                              calculator.NewAnnualRetirementAccountDisbursement(),
	"ANNUAL_TAX_SAVINGS_WITH_CONTRIBUTION":                                                calculator.NewAnnualTaxSavingsWithContribution(),
	"BALANCES_ROTH_MATCHING_NET_CONTRIBUTIONS":                                            calculator.NewBalancesRothMatchingNetContributions(),
	"BALANCES_TRADITIONAL":                                                                calculator.NewBalancesTraditional(),
	"COMBINED_RETIREMENT_INCOME":                                                          calculator.NewCombinedRetirementIncome(),
	"EFFECTIVE_TAX_RATE_ON_GROSS":                                                         calculator.NewEffectiveTaxRateOnGross(),
	"EQUIVALENT_ROTH_CONTRIBUTIONS":                                                       calculator.NewEquivalentRothContributions(),
	"HALF_OF_SOCIAL_SECURITY":                                                             calculator.NewHalfOfSocialSecurity(),
	"INCOME_AFTER_STANDARD_DEDUCTION":                                                     calculator.NewIncomeAfterStandardDeduction(),
	"INCOME_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS":                                   calculator.NewIncomeAfterStandardDeductionAndContributions(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION":                                         calculator.NewIncomePerBracketAfterStandardDeduction(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS":                       calculator.NewIncomePerBracketAfterStandardDeductionAndContributions(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_HEAD_OF_HOUSEHOLD":     calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_JOINT":         calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_SEPERATE":      calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeperate(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_SINGLE":                calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsSingle(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD":                       calculator.NewIncomePerBracketAfterStandardDeductionHeadOfHousehold(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_MARRIED_JOINT":                           calculator.NewIncomePerBracketAfterStandardDeductionMarriedJoint(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_MARRIED_SEPERATE":                        calculator.NewIncomePerBracketAfterStandardDeductionMarriedSeperate(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_SINGLE":                                  calculator.NewIncomePerBracketAfterStandardDeductionSingle(),
	"NET_DISTRIBUTION_AFTER_TAXES":                                                        calculator.NewNetDistributionAfterTaxes(),
	"ROTH_OR_TRADITIONAL_DECISION":                                                        calculator.NewRothOrTraditionalDecision(),
	"SOCIAL_SECURITY_TAXABLE_INCOME":                                                      calculator.NewSocialSecurityTaxableIncome(),
	"SOCIAL_SECURITY_TAXABLE_INCOME_INDIVIDUAL":                                           calculator.NewSocialSecurityTaxableIncomeIndividual(),
	"SOCIAL_SECURITY_TAXABLE_INCOME_JOINT":                                                calculator.NewSocialSecurityTaxableIncomeJoint(),
	"STANDARD_DEDUCTION":                                                                  calculator.NewStandardDeduction(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION":                                     calculator.NewTaxesOwedPerBracketAfterStandardDeduction(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS":                   calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributions(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_HEAD_OF_HOUSEHOLD": calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_JOINT":     calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_SEPERATE":  calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_SINGLE":            calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD":                   calculator.NewTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_MARRIED_JOINT":                       calculator.NewTaxesOwedPerBracketAfterStandardDeductionMarriedJoint(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_MARRIED_SEPERATE":                    calculator.NewTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_SINGLE":                              calculator.NewTaxesOwedPerBracketAfterStandardDeductionSingle(),
	"TAX_ON_TRADITIONAL_IRA_WITHDRAWAL":                                                   calculator.NewTaxOnTraditionalIRAWithdrawal(),
	"TAX_RATE_OF_SAVINGS":                                                                 calculator.NewTaxRateOfSavings(),
	"TOP_TIER_TAX_RATE":                                                                   calculator.NewTopTierTaxRate(),
	"TOTAL_ANNUAL_RETIREMENT_INCOME_BEFORE_TAX":                                           calculator.NewTotalAnnualRetirementIncomeBeforeTax(),
	"TOTAL_ANNUAL_RETIREMENT_INCOME_BEFORE_TAX_LESS_TAX_ON_TRADITIONAL_IRA_WITHDRAWAL":    calculator.NewTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal(),
	"TOTAL_CONTRIBUTIONS":                                                                 calculator.NewTotalContributions(),
	"TOTAL_DISBURSEMENTS":                                                                 calculator.NewTotalDisbursements(),
	"TOTAL_INTEREST":                                                                      calculator.NewTotalInterest(),
	"TOTAL_TAXABLE_INCOME":                                                                calculator.NewTotalTaxableIncome(),
	"TOTAL_TAXABLE_INCOME_AFTER_STANDARD_DEDUCTION":                                       calculator.NewTotalTaxableIncomeAfterStandardDeduction(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION":                                           calculator.NewTotalTaxesOwedAfterStandardDeduction(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS":                         calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributions(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_HEAD_OF_HOUSEHOLD":       calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_JOINT":           calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_SEPERATE":        calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeperate(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_SINGLE":                  calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsSingle(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD":                         calculator.NewTotalTaxesOwedAfterStandardDeductionHeadOfHousehold(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_MARRIED_JOINT":                             calculator.NewTotalTaxesOwedAfterStandardDeductionMarriedJoint(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_MARRIED_SEPERATE":                          calculator.NewTotalTaxesOwedAfterStandardDeductionMarriedSeperate(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_SINGLE":                                    calculator.NewTotalTaxesOwedAfterStandardDeductionSingle(),
}
