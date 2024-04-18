package main

import (
// "embed"
// "encoding/json"
// "html/template"
// "net/http"
// "strconv"

// "github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

/*
// go:embed templates/index.html
// go:embed templates/inputform.html
// go:embed constants.json
// var res embed.FS
*/
func main() {
	// calculator.AnnualGrowthLessInflation(calculator.Model{})

	// tmpl := make(map[string]*template.Template)
	// tmpl["index"] = template.Must(template.ParseFS(res, "templates/index.html", "templates/inputform.html"))
	// tmpl["inputform"] = template.Must(template.ParseFS(res, "templates/inputform.html"))

	// homeFunc := func(w http.ResponseWriter, r *http.Request) {
	// 	inputMap := map[string]*calculator.Input{
	// 		"input": getInput(false, r),
	// 	}

	// 	tmpl["index"].Execute(w, inputMap)
	// }

	// calculateFunc := func(w http.ResponseWriter, r *http.Request) {
	// 	inputMap := map[string]*calculator.Input{
	// 		"input": getInput(true, r),
	// 	}

	// 	tmpl["index"].Execute(w, inputMap)
	// }

	// http.HandleFunc("/", homeFunc)
	// http.HandleFunc("/calculate/", calculateFunc)

}

// func loadConstants() *calculator.Constants {
// 	constants := calculator.Constants{}

// 	j, err := res.ReadFile("constants.json")
// 	if err != nil {
// 		panic(err)
// 	}

// 	json.Unmarshal(j, &constants)

// 	return &constants
// }

// func getInput(calculate bool, r *http.Request) *calculator.Input {
// 	if calculate {
// 		currentAge, _ := strconv.Atoi(r.PostFormValue("current-age"))
// 		currentFilingStatus := r.PostFormValue("current-filing-status")
// 		currentAnnualIncome, _ := strconv.ParseFloat(r.PostFormValue("current-annual-income"), 64)
// 		annualInvestmentGrowth, _ := strconv.ParseFloat(r.PostFormValue("annual-investment-growth"), 64)
// 		retirementAge, _ := strconv.Atoi(r.PostFormValue("retirement-age"))
// 		retirementFilingStatus := r.PostFormValue("retirement-filing-status")
// 		yearlyWithdrawal, _ := strconv.ParseFloat(r.PostFormValue("yearly-withdrawal"), 64)

// 		return &calculator.Input{
// 			CurrentAge:             currentAge,
// 			CurrentFilingStatus:    currentFilingStatus,
// 			CurrentAnnualIncome:    currentAnnualIncome,
// 			AnnualInvestmentGrowth: annualInvestmentGrowth,
// 			RetirementAge:          retirementAge,
// 			RetirementFilingStatus: retirementFilingStatus,
// 			YearlyWithdrawal:       yearlyWithdrawal,
// 		}
// 	}

// 	return &calculator.Input{
// 		CurrentAge:                30,
// 		CurrentFilingStatus:       "single",
// 		CurrentAnnualIncome:       60000,
// 		AnnualContributionsPreTax: 23000,
// 		AnnualInvestmentGrowth:    8,
// 		RetirementAge:             65,
// 		RetirementFilingStatus:    "single",
// 		YearlyWithdrawal:          60000,
// 	}
// }
