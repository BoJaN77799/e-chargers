package main

import (
	"charger_service/pkg/router"
	"log"
)

func main() {

	log.Println("ChargerService starting up....")
	router.HandleRequests()
}
