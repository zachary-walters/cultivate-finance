package db

import (
	"embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

var pgURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	os.Getenv("PG_USER"),
	os.Getenv("PG_PASS"),
	os.Getenv("PG_HOST"),
	os.Getenv("PG_PORT"),
	os.Getenv("PG_DB"))

type connectionService struct{}

func NewConnectionService() *connectionService {
	return &connectionService{}
}

func (srv *connectionService) ConnectNats() *nats.Conn {
	var nc *nats.Conn
	var err error

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

	return nc
}

func (srv *connectionService) ConnectPostgres() *sqlx.DB {
	var pg *sqlx.DB
	var err error

	for {
		pg, err = sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("PG_HOST"),
			os.Getenv("PG_PORT"),
			os.Getenv("PG_USER"),
			os.Getenv("PG_PASS"),
			os.Getenv("PG_DB")))
		if err == nil {
			break
		}

		log.Println("Error connecting to postgres, trying again in: ", err)
		time.Sleep(2 * time.Second)
	}
	log.Println("Connected to POSTGRES.")

	return pg
}

func (svc *connectionService) Migrate() {
	d, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		fmt.Println(err)
		return
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", d, pgURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = migrations.Up()
	if err != nil && err.Error() != "no change" {
		fmt.Println(err)
		return
	}
}
