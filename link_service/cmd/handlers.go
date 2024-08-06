package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	services "github.com/zachary-walters/cultivate-finance/link_service/internal"
)

type NatsHandler struct {
	PGDB *sqlx.DB
}
type HTTPHandler struct {
	PGDB *sqlx.DB
}

func (h *NatsHandler) Generate(msg *nats.Msg) {
}

func (h *HTTPHandler) GenerateAll(w http.ResponseWriter, r *http.Request) {
	linkService := services.NewLinkService(h.PGDB, r.Context())

	err := linkService.GenerateAll(4, "abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *HTTPHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	linkService := services.NewLinkService(h.PGDB, r.Context())
	slug := chi.URLParam(r, "slug")

	link, err := linkService.GetLink(slug)
	if err != nil && err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Shared link not found"))
		if err != nil {
			panic(err)
		}
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(link)
	if err != nil {
		panic(err)
	}
}

func (h *HTTPHandler) SaveLink(w http.ResponseWriter, r *http.Request) {
	linkService := services.NewLinkService(h.PGDB, r.Context())

	input := services.Input{
		Link: "somelink",
	}

	link := services.Link{
		Link: &input.Link,
	}

	link, err := linkService.SaveLink(link)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(link)
	if err != nil {
		panic(err)
	}
}

func (h *HTTPHandler) UpdateExpiredLinks(w http.ResponseWriter, r *http.Request) {
	linkService := services.NewLinkService(h.PGDB, r.Context())

	err := linkService.UpdateExpiredLinks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
