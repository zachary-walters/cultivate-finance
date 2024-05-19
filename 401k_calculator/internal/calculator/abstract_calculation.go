package calculator

import (
	"strconv"
	"sync"
)

type Calculation interface {
	CalculateTraditional(Model) float64
	CalculateTraditionalRetirement(Model) float64
	CalculateRoth(Model) float64
	CalculateRothRetirement(Model) float64
}

type SequenceCalculation interface {
	CalculateTraditional(Model) []float64
	CalculateTraditionalRetirement(Model) []float64
	CalculateRoth(Model) []float64
	CalculateRothRetirement(Model) []float64
}

type ChartCalculation interface {
	Calculate(Model) ChartData
}

type DecisionCalculation interface {
	Calculate(Model) string
}

type CalculationData struct {
	Datakey                    string
	TraditionalValue           any
	TraditionalRetirementValue any
	RothValue                  any
	RothRetirementValue        any
}

type ChartData struct {
	BeginningBalance map[int32]float64 `json:"beginning_balance,omitempty"`
	Contribution     map[int32]float64 `json:"contribution,omitempty"`
	Withdrawal       map[int32]float64 `json:"withdrawal,omitempty"`
	InterestEarned   map[int32]float64 `json:"interest_earned,omitempty"`
	EndingBalance    map[int32]float64 `json:"ending_balance,omitempty"`
	AfterTaxIncome   map[int32]float64 `json:"after_tax_income,omitempty"`
}

type Input struct {
	CurrentAge                int32   `json:"current_age"`
	CurrentFilingStatus       string  `json:"current_filing_status"`
	CurrentAnnualIncome       float64 `json:"current_annual_income"`
	AnnualContributionsPreTax float64 `json:"annual_contributions_pretax"`
	AnnualInvestmentGrowth    float64 `json:"annual_investment_growth"`
	RetirementAge             int32   `json:"retirement_age"`
	RetirementFilingStatus    string  `json:"retirement_filing_status"`
	YearlyWithdrawal          float64 `json:"yearly_withdrawal"`
	// Extended
	WorkIncome                              float64 `json:"work_income"`
	QualifiedDividends                      float64 `json:"qualified_dividends"`
	OtherLongTermCapitalGains               float64 `json:"other_long_term_capital_gains"`
	PensionIncome                           float64 `json:"pension_income"`
	RentalNetIncome                         float64 `json:"rental_net_income"`
	AnnuityIncome                           float64 `json:"annuity_income"`
	SocialSecurity                          float64 `json:"social_security"`
	OtherTaxableIncome                      float64 `json:"other_taxable_income"`
	EstimatedTaxPercentIncreaseAtRetirement float64 `json:"estimated_tax_percent_increase_at_retirement"`
	StandardDeductionIncreaseDecrease       float64 `json:"standard_deduction_increase_decrease"`

	Datakey string `json:"datakey,omitempty"`
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
	SocialSecurityTaxRatesIndividual     []TaxRate `json:"social_security_tax_rates_individual"`
	SocialSecurityTaxRatesJoint          []TaxRate `json:"social_security_tax_rates_joint"`
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
		SocialSecurityTaxRatesIndividual:     Constants.SocialSecurityTaxRatesIndividual,
		SocialSecurityTaxRatesJoint:          Constants.SocialSecurityTaxRatesJoint,
		STANDARD_DEDUCTION_SINGLE:            Constants.STANDARD_DEDUCTION_SINGLE,
		STANDARD_DEDUCTION_MARRIED_JOINT:     Constants.STANDARD_DEDUCTION_MARRIED_JOINT,
		STANDARD_DEDUCTION_MARRIED_SEPERATE:  Constants.STANDARD_DEDUCTION_MARRIED_SEPERATE,
		STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD: Constants.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD,
	}
}

func coalesce[T int | float64](number T) T {
	// if NaN or negative
	if number != number || number < 0 {
		return 0
	}

	return number
}

func CalculateSynchronous(model Model, calculation any, datakey string) CalculationData {
	calc, isCalculation := calculation.(Calculation)
	seq, isSequenceCalculation := calculation.(SequenceCalculation)
	chart, isChartCalculation := calculation.(ChartCalculation)
	decision, isDecisionCalculation := calculation.(DecisionCalculation)

	calculationData := CalculationData{
		Datakey: datakey,
	}

	if isDecisionCalculation {
		calculationData.TraditionalValue = decision.Calculate(model)
		calculationData.TraditionalRetirementValue = nil
		calculationData.RothValue = nil
		calculationData.RothRetirementValue = nil
	} else if isSequenceCalculation {
		calculationData.TraditionalValue = seq.CalculateTraditional(model)
		calculationData.TraditionalRetirementValue = seq.CalculateTraditionalRetirement(model)
		calculationData.RothValue = seq.CalculateRoth(model)
		calculationData.RothRetirementValue = seq.CalculateRothRetirement(model)
	} else if isChartCalculation {
		calculationData.TraditionalValue = chart.Calculate(model)
		calculationData.TraditionalRetirementValue = nil
		calculationData.RothValue = nil
		calculationData.RothRetirementValue = nil
	} else if isCalculation {
		calculationData.TraditionalValue = calc.CalculateTraditional(model)
		calculationData.TraditionalRetirementValue = calc.CalculateTraditionalRetirement(model)
		calculationData.RothValue = calc.CalculateRoth(model)
		calculationData.RothRetirementValue = calc.CalculateRothRetirement(model)
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
	chart, isChartCalculation := calculation.(ChartCalculation)
	decision, isDecisionCalculation := calculation.(DecisionCalculation)

	calculationData := CalculationData{
		Datakey: datakey,
	}

	if isDecisionCalculation {
		calculationData.TraditionalValue = decision.Calculate(model)
		calculationData.TraditionalRetirementValue = nil
		calculationData.RothValue = nil
		calculationData.RothRetirementValue = nil
	} else if isSequenceCalculation {
		calculationData.TraditionalValue = translateFloatSlice(seq.CalculateTraditional(model))
		calculationData.TraditionalRetirementValue = translateFloatSlice(seq.CalculateTraditionalRetirement(model))
		calculationData.RothValue = translateFloatSlice(seq.CalculateRoth(model))
		calculationData.RothRetirementValue = translateFloatSlice(seq.CalculateRothRetirement(model))
	} else if isChartCalculation {
		calculationData.TraditionalValue = translateChartData(chart.Calculate(model))
		calculationData.TraditionalRetirementValue = nil
		calculationData.RothValue = nil
		calculationData.RothRetirementValue = nil
	} else if isCalculation {
		calculationData.TraditionalValue = calc.CalculateTraditional(model)
		calculationData.TraditionalRetirementValue = calc.CalculateTraditionalRetirement(model)
		calculationData.RothValue = calc.CalculateRoth(model)
		calculationData.RothRetirementValue = calc.CalculateRothRetirement(model)
	}

	return calculationData
}

func CalculateAsyncWasm(wg *sync.WaitGroup, ch chan CalculationData, datakey string, calculation any, model Model) {
	defer wg.Done()
	calculationData := CalculateSynchronousWasm(model, calculation, datakey)

	ch <- calculationData
}

func translateFloatSlice(s []float64) []interface{} {
	x := make([]interface{}, len(s))

	for i := range s {
		x[i] = s[i]
	}

	return x
}

func translateChartMap(m map[int32]float64) map[string]interface{} {
	n := map[string]interface{}{}

	for k, v := range m {
		n[strconv.Itoa(int(k))] = v
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
