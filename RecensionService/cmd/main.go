package main

import (
	"log"
	"recension_service/pkg/router"
)

func main() {

	log.Println("RecensionService starting up....")
	router.HandleRequests()
}
