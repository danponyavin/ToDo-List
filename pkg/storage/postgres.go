package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type PostgresStorage struct {
	db *sqlx.DB
}

func NewPostgresStorage() *PostgresStorage {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5428"
	}
	connStr := fmt.Sprintf("host=%s port=%s user=postgres password=mypass dbname=test sslmode=disable", host, port)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStorage{db: db}
}
