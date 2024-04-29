package calculator

import (
	"strconv"
	"sync"
)

type Calculation interface {
	Calculate(Model) float64
	CalculateRetirement(Model) float64
}

type SequenceCalculation interface {
	Calculate(Model) []float64
	CalculateRetirement(Model) []float64
}

type ChartCalculation interface {
	Calculate(Model) ChartData
}

type DecisionCalculation interface {
	Calculate(Model) string
	CalculateRetirement(Model) string
}

type ChartData struct {
	BeginningBalance map[int]float64 `json:"beginning_balance,omitempty"`
	Contribution     map[int]float64 `json:"contribution,omitempty"`
	Withdrawal       map[int]float64 `json:"withdrawal,omitempty"`
	InterestEarned   map[int]float64 `json:"interest_earned,omitempty"`
	EndingBalance    map[int]float64 `json:"ending_balance,omitempty"`
	AfterTaxIncome   map[int]float64 `json:"after_tax_income,omitempty"`
}

type Input struct {
	CurrentAge                int     `json:"current_age"`
	CurrentFilingStatus       string  `json:"current_filing_status"`
	CurrentAnnualIncome       float64 `json:"current_annual_income"`
	AnnualContributionsPreTax float64 `json:"annual_contributions_pretax"`
	AnnualInvestmentGrowth    float64 `json:"annual_investment_growth"`
	RetirementAge             int     `json:"retirement_age"`
	RetirementFilingStatus    string  `json:"retirement_filing_status"`
	YearlyWithdrawal          float64 `json:"yearly_withdrawal"`
}

type TaxRate struct {
	Cap  float64 `json:"cap"`
	Rate float64 `json:"rate"`
}

type Model struct {
	Input                                Input
	SingleTaxRates                       []TaxRate `json:"single_tax_rates"`
	MarriedJointTaxRates                 []TaxRate `json:"married_joint_tax_rates"`
	MarriedSeperateTaxRates              []TaxRate `json:"married_seperate_tax_rates"`
	HeadOfHouseholdTaxRates              []TaxRate `json:"head_of_household_tax_rates"`
	STANDARD_DEDUCTION_SINGLE            float64   `json:"standard_deduction_single"`
	STANDARD_DEDUCTION_MARRIED_JOINT     float64   `json:"standard_deduction_married_joint"`
	STANDARD_DEDUCTION_MARRIED_SEPERATE  float64   `json:"standard_deduction_married_seperate"`
	STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD float64   `json:"standard_deduction_head_of_household"`
}

func NewModel(input Input) Model {
	return Model{
		Input:                                input,
		SingleTaxRates:                       Constants.SingleTaxRates,
		MarriedJointTaxRates:                 Constants.MarriedJointTaxRates,
		MarriedSeperateTaxRates:              Constants.MarriedSeperateTaxRates,
		HeadOfHouseholdTaxRates:              Constants.HeadOfHouseholdTaxRates,
		STANDARD_DEDUCTION_SINGLE:            Constants.STANDARD_DEDUCTION_SINGLE,
		STANDARD_DEDUCTION_MARRIED_JOINT:     Constants.STANDARD_DEDUCTION_MARRIED_JOINT,
		STANDARD_DEDUCTION_MARRIED_SEPERATE:  Constants.STANDARD_DEDUCTION_MARRIED_SEPERATE,
		STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD: Constants.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD,
	}
}

func CalculateSynchronous(model Model, calculation any) (value any, retirementValue any) {
	calc, isCalculation := calculation.(Calculation)
	seq, isSequenceCalculation := calculation.(SequenceCalculation)
	chart, isChartCalculation := calculation.(ChartCalculation)
	decision, isDecisionCalculation := calculation.(DecisionCalculation)

	if isCalculation {
		value = calc.Calculate(model)
		retirementValue = calc.CalculateRetirement(model)
	} else if isSequenceCalculation {
		value = seq.Calculate(model)
		retirementValue = seq.CalculateRetirement(model)
	} else if isChartCalculation {
		value = chart.Calculate(model)
		retirementValue = nil
	} else if isDecisionCalculation {
		value = decision.Calculate(model)
		retirementValue = decision.CalculateRetirement(model)
	}

	return value, retirementValue
}

func CalculateAsync(wg *sync.WaitGroup, ch chan<- map[string]any, rch chan<- map[string]any, datakey string, calculation any, model Model) {
	value, retirementValue := CalculateSynchronous(model, calculation)

	ch <- map[string]any{datakey: value}
	rch <- map[string]any{datakey: retirementValue}

	wg.Done()
}

func CalculateSynchronousWasm(model Model, calculation any) (value any, retirementValue any) {
	calc, isCalculation := calculation.(Calculation)
	seq, isSequenceCalculation := calculation.(SequenceCalculation)
	chart, isChartCalculation := calculation.(ChartCalculation)
	decision, isDecisionCalculation := calculation.(DecisionCalculation)

	if isCalculation {
		value = calc.Calculate(model)
		retirementValue = calc.CalculateRetirement(model)
	} else if isSequenceCalculation {
		value = translateFloatSlice(seq.Calculate(model))
		retirementValue = translateFloatSlice(seq.CalculateRetirement(model))
	} else if isChartCalculation {
		value = translateChartData(chart.Calculate(model))
		retirementValue = nil
	} else if isDecisionCalculation {
		value = decision.Calculate(model)
		retirementValue = decision.CalculateRetirement(model)
	}
	return value, retirementValue
}

func CalculateAsyncWasm(wg *sync.WaitGroup, ch chan<- map[string]any, rch chan<- map[string]any, datakey string, calculation any, model Model) {
	value, retirementValue := CalculateSynchronousWasm(model, calculation)

	ch <- map[string]any{datakey: value}
	rch <- map[string]any{datakey: retirementValue}

	wg.Done()
}

func translateFloatSlice(s []float64) []interface{} {
	x := make([]interface{}, len(s))

	for i := range s {
		x[i] = s[i]
	}

	return x
}

func translateChartMap(m map[int]float64) map[string]interface{} {
	n := map[string]interface{}{}

	for k, v := range m {
		n[strconv.Itoa(k)] = v
	}

	return n
}

func translateChartData(c ChartData) map[string]interface{} {
	chartData := map[string]interface{}{}

	chartData["beginning_balance"] = translateChartMap(c.BeginningBalance)
	chartData["contribution"] = translateChartMap(c.Contribution)
	chartData["withdrawal"] = translateChartMap(c.Withdrawal)
	chartData["interest_earned"] = translateChartMap(c.InterestEarned)
	chartData["ending_balance"] = translateChartMap(c.EndingBalance)
	chartData["after_tax_income"] = translateChartMap(c.AfterTaxIncome)

	return chartData
}

var Calculations = map[string]any{
	"ANNUAL_GROWTH_LESS_INFLATION":                                                        NewAnnualGrowthLessInflation(),
	"ANNUAL_TAX_SAVINGS_WITH_CONTRIBUTION":                                                NewAnnualTaxSavingsWithContribution(),
	"BALANCES_ROTH_MATCHING_GROSS_CONTRIBUTIONS":                                          NewBalancesRothMatchingGrossContributions(),
	"BALANCES_ROTH_MATCHING_NET_CONTRIBUTIONS":                                            NewBalancesRothMatchingNetContributions(),
	"BALANCES_TRADITIONAL":                                                                NewBalancesTraditional(),
	"EFFECTIVE_TAX_RATE_ON_GROSS":                                                         NewEffectiveTaxRateOnGross(),
	"EQUIVALENT_ROTH_CONTRIBUTIONS":                                                       NewEquivalentRothContributions(),
	"INCOME_AFTER_STANDARD_DEDUCTION":                                                     NewIncomeAfterStandardDeduction(),
	"INCOME_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS":                                   NewIncomeAfterStandardDeductionAndContributions(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_HEAD_OF_HOUSEHOLD":     NewIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_JOINT":         NewIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_SEPERATE":      NewIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeperate(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_SINGLE":                NewIncomePerBracketAfterStandardDeductionAndContributionsSingle(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS":                       NewIncomePerBracketAfterStandardDeductionAndContributions(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD":                       NewIncomePerBracketAfterStandardDeductionHeadOfHousehold(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_MARRIED_JOINT":                           NewIncomePerBracketAfterStandardDeductionMarriedJoint(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_MARRIED_SEPERATE":                        NewIncomePerBracketAfterStandardDeductionMarriedSeperate(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION_SINGLE":                                  NewIncomePerBracketAfterStandardDeductionSingle(),
	"INCOME_PER_BRACKET_AFTER_STANDARD_DEDUCTION":                                         NewIncomePerBracketAfterStandardDeduction(),
	"NET_DISTRIBUTION_AFTER_TAXES":                                                        NewNetDistributionAfterTaxes(),
	"ROTH_OR_TRADITIONAL_DECISION":                                                        NewRothOrTraditionalDecision(),
	"STANDARD_DEDUCTION":                                                                  NewStandardDeduction(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_HEAD_OF_HOUSEHOLD": NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_JOINT":     NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_SEPERATE":  NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_SINGLE":            NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS":                   NewTaxesOwedPerBracketAfterStandardDeductionAndContributions(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD":                   NewTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_MARRIED_JOINT":                       NewTaxesOwedPerBracketAfterStandardDeductionMarriedJoint(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_MARRIED_SEPERATE":                    NewTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION_SINGLE":                              NewTaxesOwedPerBracketAfterStandardDeductionSingle(),
	"TAXES_OWED_PER_BRACKET_AFTER_STANDARD_DEDUCTION":                                     NewTaxesOwedPerBracketAfterStandardDeduction(),
	"TAX_RATE_OF_SAVINGS":                                                                 NewTaxRateOfSavings(),
	"TOTAL_CONTRIBUTIONS":                                                                 NewTotalContributions(),
	"TOTAL_CONTRIBUTIONS_ROTH":                                                            NewTotalContributionsRoth(),
	"TOTAL_DISBURSEMENTS_AFTER_TAX":                                                       NewTotalDisbursementsAfterTax(),
	"TOTAL_DISBURSEMENTS_ROTH_MATCHING_GROSS":                                             NewTotalDisbursementsRothMatchingGross(),
	"TOTAL_DISBURSEMENTS_ROTH_MATCHING_NET":                                               NewTotalDisbursementsRothMatchingNet(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION":                                           NewTotalTaxesOwedAfterStandardDeduction(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS":                         NewTotalTaxesOwedAfterStandardDeductionAndContributions(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_HEAD_OF_HOUSEHOLD":       NewTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_JOINT":           NewTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_MARRIED_SEPERATE":        NewTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeperate(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS_SINGLE":                  NewTotalTaxesOwedAfterStandardDeductionAndContributionsSingle(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD":                         NewTotalTaxesOwedAfterStandardDeductionHeadOfHousehold(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_MARRIED_JOINT":                             NewTotalTaxesOwedAfterStandardDeductionMarriedJoint(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_MARRIED_SEPERATE":                          NewTotalTaxesOwedAfterStandardDeductionMarriedSeperate(),
	"TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION_SINGLE":                                    NewTotalTaxesOwedAfterStandardDeductionSingle(),
}
