package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type data struct {
	Value           any `json:"value,omitempty"`
	RetirementValue any `json:"retirement_value,omitempty"`
}

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/{datakey}", getCalculationByDatakey)
	r.Post("/calculate_all", calculateAll)
	http.ListenAndServe(":8660", r)
}

func calculateAll(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input calculator.Input

	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}

	model := calculator.NewModel(input)

	defaultCh := make(chan map[string]any, len(calculator.Calculations))
	retirementCh := make(chan map[string]any, len(calculator.Calculations))
	wg := &sync.WaitGroup{}
	for datakey, calculation := range calculator.Calculations {
		wg.Add(1)
		go calculator.CalculateAsync(wg, defaultCh, retirementCh, datakey, calculation, model)
	}
	wg.Wait()

	close(defaultCh)
	close(retirementCh)

	m := map[string]any{}
	rm := map[string]any{}

	for i := 0; i < len(calculator.Calculations)*2; i++ {
		select {
		case data := <-defaultCh:
			for datakey, value := range data {
				m[datakey] = value
			}
		case data := <-retirementCh:
			for datakey, value := range data {
				rm[datakey] = value
			}
		}
	}

	modelMap := map[string]data{}
	for datakey := range m {
		modelMap[datakey] = data{
			Value:           m[datakey],
			RetirementValue: rm[datakey],
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(modelMap)
	if err != nil {
		panic(err)
	}
}

func getCalculationByDatakey(w http.ResponseWriter, r *http.Request) {
	datakey := strings.ToUpper(chi.URLParam(r, "datakey"))

	calculation, exists := calculator.Calculations[datakey]
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

	model := calculator.NewModel(input)

	value, retirementValue := calculator.CalculateSynchronous(model, calculation)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
}
