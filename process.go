package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const DBCONN = "postgres://postgres:postgres@localhost:5432/process_db?sslmode=disable"

func main() {
	// Open and ping a database connection to the docker db image.
	// Use this db connection if you like, it requires that the docker image
	// is running (via docker compose up).
	db, err := sql.Open("postgres", DBCONN)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	// Add process code here.
	fmt.Printf("Process db connection: %s\n", DBCONN)
}
