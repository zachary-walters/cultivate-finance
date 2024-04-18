package calculator

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

type ChartData struct {
	AgeSequence      map[int]float64 `json:"age_sequence,omitempty"`
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
