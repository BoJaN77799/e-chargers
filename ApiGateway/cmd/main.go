package main

import "ApiGateway/pkg/router"

func main() {
	router.HandleRequests(50000)
}
