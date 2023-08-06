package main

import (
	"reservation_service/pkg/router"
)

func main() {
	router.HandleRequests(50003)
}
