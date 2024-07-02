package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"
)

type server struct {
	nc *nats.Conn
}

type data401k struct {
	TraditionalValue           any `json:"traditional_value,omitempty"`
	TraditionalRetirementValue any `json:"traditional_retirement_value,omitempty"`
	RothValue                  any `json:"roth_value,omitempty"`
	RothRetirementValue        any `json:"roth_retirement_value,omitempty"`
}

type data401kDatakey struct {
	Datakey                    string `json:"datakey"`
	TraditionalValue           any    `json:"traditional_value,omitempty"`
	TraditionalRetirementValue any    `json:"traditional_retirement_value,omitempty"`
	RothValue                  any    `json:"roth_value,omitempty"`
	RothRetirementValue        any    `json:"roth_retirement_value,omitempty"`
}

type dataDebtSnowballDatakey struct {
	Datakey string `json:"datakey,omitempty"`
	Value   any    `json:"value,omitempty"`
}

func (s server) calculate401kDatakey(w http.ResponseWriter, r *http.Request) {
	datakey := strings.ToUpper(chi.URLParam(r, "datakey"))

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error parsing request body: ", err)
		return
	}

	input := struct {
		Datakey                   string  `json:"datakey"`
		CurrentAge                int     `json:"current_age"`
		CurrentFilingStatus       string  `json:"current_filing_status"`
		CurrentAnnualIncome       float64 `json:"current_annual_income"`
		AnnualContributionsPreTax float64 `json:"annual_contributions_pretax"`
		AnnualInvestmentGrowth    float64 `json:"annual_investment_growth"`
		RetirementAge             int     `json:"retirement_age"`
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
	}{
		Datakey: datakey,
	}

	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Println("Error unmarshalling body into inputs: ", err)
		return
	}

	reqData, err := json.Marshal(input)
	if err != nil {
		log.Println("Error marshalling inputs into bytes: ", err)
		return
	}

	response, err := s.nc.Request("calculate_401k_by_datakey", reqData, 5*time.Second)
	if err != nil {
		log.Println("Error making NATS request:", err)
	}

	data := new(data401kDatakey)
	err = json.Unmarshal(response.Data, data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}

func (s server) calculateAll401k(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error parsing request body: ", err)
		return
	}
	response, err := s.nc.Request("calculate_all_401k", body, 5*time.Second)
	if err != nil {
		log.Println("Error making NATS request:", err)
	}

	modelMap := map[string]data401k{}
	err = json.Unmarshal(response.Data, &modelMap)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(modelMap)
	if err != nil {
		panic(err)
	}
}

func (s server) calculateDebtSnowball(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error parsing request body: ", err)
		return
	}

	response, err := s.nc.Request("calculate_all_debt_snowball", body, 5*time.Second)
	if err != nil {
		log.Println("Error making NATS request:", err)
	}

	modelMap := map[string]dataDebtSnowballDatakey{}
	err = json.Unmarshal(response.Data, &modelMap)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(modelMap)
	if err != nil {
		panic(err)
	}
}

func (s server) calculateDebtSnowballDatakey(w http.ResponseWriter, r *http.Request) {
	datakey := strings.ToUpper(chi.URLParam(r, "datakey"))

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error parsing request body: ", err)
		return
	}

	input := struct {
		Debts []struct {
			Name           string  `json:"name"`
			Balance        float64 `json:"balance"`
			Amount         float64 `json:"amount"`
			AnnualInterest float64 `json:"annual_interest"`
			MimimumPayment float64 `json:"minimum_payment"`
		} `json:"debts"`
		ExtraMonthlyPayment     float64 `json:"extra_monthly_payment"`
		OneTimeImmediatePayment float64 `json:"one_time_immediate_payment"`
		Datakey                 string  `json:"datakey"`
	}{
		Datakey: datakey,
	}

	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Println("Error unmarshalling body into inputs: ", err)
		return
	}

	reqData, err := json.Marshal(input)
	if err != nil {
		log.Println("Error marshalling inputs into bytes: ", err)
		return
	}

	response, err := s.nc.Request("calculate_debt_snowball_by_datakey", reqData, 5*time.Second)
	if err != nil {
		log.Println("Error making NATS request:", err)
	}

	data := new(dataDebtSnowballDatakey)
	err = json.Unmarshal(response.Data, data)
	if err != nil {
		panic(err)
	}

	data.Datakey = datakey

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
