package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"
)

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

func main() {
	var ns server
	var err error
	r := chi.NewRouter()

	uri := os.Getenv("NATS_URI")
	for {
		ns.nc, err = nats.Connect(uri)
		if err == nil {
			break
		}

		log.Println("Waiting before connecting to NATS at:", uri)
		time.Sleep(1 * time.Second)
	}
	log.Println("Connected to NATS at:", ns.nc.ConnectedUrl())

	r.Post("/calculate_all_401k", ns.calculateAll401k)
	r.Post("/calculate_401k/{datakey}", ns.calculateDatakey)

	log.Println("Server listening on port: ", os.Getenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r); err != nil {
		log.Fatal(err)
	}
}
