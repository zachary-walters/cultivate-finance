package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"text/template"
)

//go:embed templates/*
var res embed.FS

type Input struct {
	CurrentAge                int     `json:"current_age"`
	CurrentFilingStatus       string  `json:"current_filing_status"`
	CurrentAnnualIncome       float64 `json:"current_annual_income"`
	AnnualContributionsPreTax float64 `json:"annual_contributions_pretax"`
	AnnualInvestmentGrowth    float64 `json:"annual_investment_growth"`
	RetirementAge             int     `json:"retirement_age"`
	RetirementFilingStatus    string  `json:"retirement_filing_status"`
	WorkIncome                float64 `json:"work_income"`
	QualifiedDividends        float64 `json:"qualified_dividends"`
	OtherLongTermCapitalGains float64 `json:"other_long_term_capital_gains"`
	PensionIncome             float64 `json:"pension_income"`
	RentalNetIncome           float64 `json:"rental_net_income"`
	AnnuityIncome             float64 `json:"annuity_income"`
	SocialSecurity            float64 `json:"social_security"`
	OtherTaxableIncome        float64 `json:"other_taxable_income"`
	YearlyWithdrawal          float64 `json:"yearly_withdrawal"`
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/wiki", wiki)
	http.HandleFunc("/roth_vs_traditional", rothVsTraditional)
	http.HandleFunc("/debt_snowball", debtSnowball)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("/assets"))))

	log.Fatal(http.ListenAndServe(":8662", nil))
}

func rothVsTraditional(w http.ResponseWriter, r *http.Request) {
	templates, err := template.New("").
		ParseFS(res,
			"templates/401k_calculator/401k_calculator.html",
			"templates/401k_calculator/401k_calculator_input_form.html",
			"templates/401k_calculator/401k_calculator_decision.html",
			"templates/401k_calculator/401k_calculator_contributions_interest_charts.html",
			"templates/401k_calculator/401k_calculator_area_chart.html",
			"templates/401k_calculator/401k_calculator_calculations.html",
			"templates/navigation_bar.html")
	if err != nil {
		log.Fatal(err)
	}

	templates.ExecuteTemplate(w, "401k_calculator.html", map[string]interface{}{
		"input": Input{
			CurrentAge:                35,
			CurrentFilingStatus:       "single",
			CurrentAnnualIncome:       60000,
			AnnualContributionsPreTax: 10000,
			AnnualInvestmentGrowth:    0.08,
			RetirementAge:             70,
			RetirementFilingStatus:    "single",
			YearlyWithdrawal:          60000,
		},
		"build_env": os.Getenv("BUILD_ENV"),
	})
}

func debtSnowball(w http.ResponseWriter, r *http.Request) {
	templates, err := template.New("").
		ParseFS(res,
			"templates/debt_snowball_calculator/debt_snowball_calculator.html",
			"templates/debt_snowball_calculator/debt_snowball_input_form.html",
			"templates/debt_snowball_calculator/debt_snowball_calculations.html",
			"templates/debt_snowball_calculator/debt_snowball_months_payoff_chart.html",
			"templates/debt_snowball_calculator/debt_snowball_payoff_over_time_chart.html",
			"templates/debt_snowball_calculator/debt_snowball_donut_chart.html",
			"templates/debt_snowball_calculator/debt_snowball_decision.html",
			"templates/debt_snowball_calculator/debt_snowball_debt_order.html",
			"templates/navigation_bar.html",
		)
	if err != nil {
		log.Fatal(err)
	}

	templates.ExecuteTemplate(w, "debt_snowball_calculator.html", map[string]interface{}{})
}

func home(w http.ResponseWriter, r *http.Request) {
	templates, err := template.New("").
		ParseFS(res,
			"templates/about.html",
			"templates/navigation_bar.html",
		)
	if err != nil {
		log.Fatal(err)
	}

	templates.ExecuteTemplate(w, "about.html", map[string]interface{}{
		"build_env": os.Getenv("BUILD_ENV"),
	})
}

func wiki(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://wiki.cultivatefinance.org", http.StatusSeeOther)
}
