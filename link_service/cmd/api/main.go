package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/zachary-walters/cultivate-finance/link_service/internal/db"
)

func main() {
	r := chi.NewRouter()

	connectionSvc := db.NewConnectionService()

	nc := connectionSvc.ConnectNats()
	pg := connectionSvc.ConnectPostgres()

	connectionSvc.Migrate()

	nh := NatsHandler{
		PGDB:      pg,
	}

	hh := HTTPHandler{
		PGDB:      pg,
	}

	nc.Subscribe("generate", nh.Generate)
	r.Post("/generate_all", hh.GenerateAll)
	r.Post("/new", hh.SaveLink)
	r.Get("/cf/{slug}", hh.GetLink)

	log.Println("Server listening on port: ", os.Getenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r); err != nil {
		log.Println(err)
	}
}
