package main

import (
	"log"
	"reservation_service/pkg/router"
)

func main() {

	log.Println("ReservationService starting up....")
	router.HandleRequests()
}
