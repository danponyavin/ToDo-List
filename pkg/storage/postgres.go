package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	connStr := "host=localhost port=5428 user=postgres password=mypass dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStorage{db: db}
}
