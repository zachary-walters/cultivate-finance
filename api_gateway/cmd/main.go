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

	// 401k_calculator
	r.Post("/calculate_all_401k", ns.calculateAll401k)
	r.Post("/calculate_401k/{datakey}", ns.calculate401kDatakey)

	// debt_snowball_calculator
	r.Post("/calculate_all_debt_snowball", ns.calculateDebtSnowball)
	r.Post("/calculate_debt_snowball/{datakey}", ns.calculateDebtSnowballDatakey)

	// link_service
	r.Post("/link/generate_all", ns.linkGenerateAll)
	r.Get("/link/get_link/{slug}", ns.linkGetLink)
	r.Post("/link/save_link", ns.linkSaveLink)
	r.Post("/link/cleanup_expired", ns.linkUpdateExpiredLinks)

	log.Println("Server listening on port: ", os.Getenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r); err != nil {
		log.Fatal(err)
	}
}
