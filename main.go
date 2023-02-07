package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"github.com/gorilla/sessions"
	"crud/api"
	"net/http"
	"os"
	"context"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/crud?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nul {
		log.Println("Can not connect to db: ", err)
	}

	server := api.NewServer(context.Context)
	err = server.Start(serverAddress)
	if err != nil {
		log.Println("Can not start the server: ", err)
	}
}