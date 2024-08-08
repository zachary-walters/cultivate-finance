package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type HTTPHandler struct {
	PGDB *sqlx.DB
}

func (h *HTTPHandler) GenerateAll(w http.ResponseWriter, r *http.Request) {
	err := generateAll(h.PGDB, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *HTTPHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	link, err := getLink(h.PGDB, r.Context(), slug)
	if err != nil && err == sql.ErrNoRows {
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
	link, err := saveLink(h.PGDB, r.Context())
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
	err := updateExpiredLinks(h.PGDB, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
