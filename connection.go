package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// getConection obtiene una conexión a la BD
func getConection() *sql.DB {
	dsn := "postgres://golang:golang@127.0.0.1:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
