package main

import (
	"log"
	"user_service/pkg/router"
)

func main() {

	log.Println("UserService starting up....")
	router.HandleRequests()
}
