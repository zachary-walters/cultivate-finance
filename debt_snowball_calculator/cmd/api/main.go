package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

type RequestError struct {
	Err error `json:"error"`
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func reqError(err error) []byte {
	data, _ := json.Marshal(RequestError{
		Err: err,
	})

	return data
}

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

	nc.Subscribe("calculate_debt_snowball", func(msg *nats.Msg) {
		log.Println("Got task request on:", msg.Subject)

		model := calculator.Model{}

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

func calculateDatakey(d []byte) (any, error) {
	data := struct {
		Datakey string  `json:"datakey"`
		Value   float64 `json:"value"`
	}{}

	decoder := json.NewDecoder(bytes.NewReader(d))
	var input calculator.Input

	err := decoder.Decode(&input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(input)

	if len(input.Debts) <= 0 {
		return data, nil
	}

	calculation, exists := calculations[input.Datakey]
	if !exists {
		// send no exists message reply to nats
		log.Println("calculation doesn't exist")
	}

	model := calculator.NewModel(input)

	calculationData := calculation.Calculate(model)

	data.Value = calculationData

	return data, nil
}

var calculations = map[string]calculator.Calculation{
	"TOTAL_BEGINNING_DEBT":  calculator.NewTotalBeginningDebt(),
	"TOTAL_MINIMUM_PAYMENT": calculator.NewTotalMinimumPayment(),
}
