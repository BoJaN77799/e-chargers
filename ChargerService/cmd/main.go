package main

import (
	"charger_service/pkg/router"
)

func main() {
	router.HandleRequests(50002)
}
