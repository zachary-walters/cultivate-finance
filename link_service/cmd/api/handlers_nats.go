package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
)

type NatsHandler struct {
	PGDB *sqlx.DB
	NC   *nats.Conn
}

func (h *NatsHandler) GenerateAll(msg *nats.Msg) {
	log.Println("Got task request on:", msg.Subject)
	ctx := context.Background()
	err := generateAll(h.PGDB, ctx)
	if err != nil {
		err := h.NC.Publish(msg.Reply, []byte(err.Error()))
		if err != nil {
			panic(err)
		}
		return
	}
}

func (h *NatsHandler) GetLink(msg *nats.Msg) {
	log.Println("Got task request on:", msg.Subject)
	ctx := context.Background()

	link, err := getLink(h.PGDB, ctx, string(msg.Data))
	if err != nil && err == sql.ErrNoRows {
		err := h.NC.Publish(msg.Reply, []byte("Shared link not found"))
		if err != nil {
			panic(err)
		}
		return
	}

	if err != nil {
		err := h.NC.Publish(msg.Reply, []byte("Internal Server Error 500"))
		if err != nil {
			panic(err)
		}
		return
	}

	linkBytes, err := json.Marshal(link)
	if err != nil {
		panic(err)
	}

	err = h.NC.Publish(msg.Reply, linkBytes)
	if err != nil {
		panic(err)
	}
}

func (h *NatsHandler) SaveLink(msg *nats.Msg) {
	log.Println("Got task request on:", msg.Subject)
	ctx := context.Background()
	link, err := saveLink(h.PGDB, ctx)
	if err != nil {
		err := h.NC.Publish(msg.Reply, []byte("Internal Server Error 500"))
		if err != nil {
			panic(err)
		}
		return
	}

	linkBytes, err := json.Marshal(link)
	if err != nil {
		panic(err)
	}

	err = h.NC.Publish(msg.Reply, linkBytes)
	if err != nil {
		panic(err)
	}
}

func (h *NatsHandler) UpdateExpiredLinks(msg *nats.Msg) {
	log.Println("Got task request on:", msg.Subject)
	ctx := context.Background()
	err := updateExpiredLinks(h.PGDB, ctx)
	if err != nil {
		err := h.NC.Publish(msg.Reply, []byte("Internal Server Error 500"))
		if err != nil {
			panic(err)
		}
		return
	}
}
