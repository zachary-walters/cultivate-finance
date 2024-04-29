package main

import (
	"embed"
	"log"
	"net/http"
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
	YearlyWithdrawal          float64 `json:"yearly_withdrawal"`
}

func main() {
	templates, err := template.New("").
		ParseFS(res,
			"templates/401k_calculator.html",
			"templates/401k_calculator_input_form.html",
			"templates/401k_calculator_decision.html",
			"templates/401k_calculator_calculations.html")
	if err != nil {
		log.Fatal(err)
	}

	home := func(w http.ResponseWriter, r *http.Request) {
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
		})
	}

	http.HandleFunc("/", home)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("/assets"))))

	log.Fatal(http.ListenAndServe(":8662", nil))
}
