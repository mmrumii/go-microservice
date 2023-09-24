package main

import (
	"log"

	"github.com/mmrumii/go-microservice/internal/database"
	"github.com/mmrumii/go-microservice/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize databae Client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
