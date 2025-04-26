package main

import (
	"log"
	"net/http"

	"github.com/gogineni1998/oolio-assignment-backend/configuration"
	"github.com/gogineni1998/oolio-assignment-backend/server"
)

func main() {

	log.Println("Starting server on", configuration.Address)

	err := http.ListenAndServe(configuration.Address, server.NewServer())
	if err != nil {
		log.Fatal(err)
	}
}
