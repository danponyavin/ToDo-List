package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type PostgresStorage struct {
	db *sqlx.DB
}

func NewPostgresStorage() *PostgresStorage {
	connStr := "host=localhost port=5428 user=postgres password=mypass dbname=test sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStorage{db: db}
}
