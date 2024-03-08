package main

import (
	"fmt"
	"log"
)

const (
	dbDriver      = "PLACEHOLDER"
	dbSource      = "PLACEHOLDER"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	//connect to DB
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	//server
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start the server", err)
	}
}
