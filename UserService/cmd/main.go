package main

import (
	"user_service/pkg/router"
)

func main() {
	router.HandleRequests(50001)
}
