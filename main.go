package main

import (
	"log"
	"net/http"

	"github.com/gogineni1998/oolio-assignment-backend/configuration"
	"github.com/gogineni1998/oolio-assignment-backend/database"
	"github.com/gogineni1998/oolio-assignment-backend/server"
)

func Run() error {
	log.Println("Starting server on", configuration.Address)
	defer database.DisconnectDB(configuration.DBClient)
	return http.ListenAndServe(configuration.Address, server.NewServer())
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
