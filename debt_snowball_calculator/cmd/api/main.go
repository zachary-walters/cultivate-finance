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
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

func main() {
	var nc *nats.Conn
	var err error
	r := chi.NewRouter()

	uri := os.Getenv("NATS_URI")
	for {
		nc, err = nats.Connect(uri)
		if err == nil {
			break
		}

		log.Println("Waiting before connecting to NATS at:", uri)
		time.Sleep(1 * time.Second)
	}
	log.Println("Connected to NATS at:", nc.ConnectedUrl())

	nc.Subscribe("calculate_all_debt_snowball", func(msg *nats.Msg) {
		log.Println("Got task request on:", msg.Subject)

		model, err := calculateModel(msg.Data)
		if err != nil {
			nc.Publish(msg.Reply, reqError(err))
			return
		}

		d, err := json.Marshal(model)
		if err != nil {
			nc.Publish(msg.Reply, reqError(err))
			return
		}

		nc.Publish(msg.Reply, d)
	})

	nc.Subscribe("calculate_debt_snowball_by_datakey", func(msg *nats.Msg) {
		log.Println("Got task request on:", msg.Subject)

		calculationData, err := calculateDatakey(msg.Data)
		if err != nil {
			nc.Publish(msg.Reply, reqError(err))
			return
		}

		d, err := json.Marshal(calculationData)
		if err != nil {
			nc.Publish(msg.Reply, reqError(err))
			return
		}

		nc.Publish(msg.Reply, d)
	})

	log.Println("Server listening on port: ", os.Getenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r); err != nil {
		log.Println(err)
	}
}

func calculateModel(d []byte) (map[string]calculator.CalculationData, error) {
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

	modelMap := map[string]calculator.CalculationData{}
	for len(ch) > 0 {
		calculationData := <-ch
		modelMap[calculationData.Datakey] = calculator.CalculationData{
			Datakey: calculationData.Datakey,
			Value:   calculationData.Value,
		}
	}

	return modelMap, nil
}

func calculateDatakey(d []byte) (any, error) {
	data := calculator.CalculationData{}

	decoder := json.NewDecoder(bytes.NewReader(d))
	var input calculator.Input

	err := decoder.Decode(&input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(input.Debts) <= 0 {
		return data, nil
	}

	calculation, exists := calculations[input.Datakey]
	if !exists {
		// send no exists message reply to nats
		log.Println("calculation does not exist for datakey: ", input.Datakey)
		return nil, NewError(fmt.Sprint("calculation does not exists for datakey: ", input.Datakey))
	}

	model := calculator.NewModel(input)

	data = calculator.CalculateSynchronous(model, calculation, input.Datakey)

	return data, nil
}

var calculations = map[string]any{
	"DEBT_PAYOFF_MONTH":         calculator.NewDebtPayoffMonth(),
	"MONTHLY_SEQUENCE_BALANCES": calculator.NewMonthlySequenceBalances(),
	"MONTHLY_SEQUENCE_PAYMENTS": calculator.NewMonthlySequencePayments(),
	"SNOWBALL":                  calculator.NewSnowball(),
	"TOTAL_BEGINNING_DEBT":      calculator.NewTotalBeginningDebt(),
	"TOTAL_INTEREST":            calculator.NewTotalInterest(),
	"TOTAL_MINIMUM_PAYMENT":     calculator.NewTotalMinimumPayment(),
	"TOTAL_PAYMENTS":            calculator.NewTotalPayments(),
}
