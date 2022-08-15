package router

import (
	"ApiGateway/pkg/handlers/ChargerService"
	"ApiGateway/pkg/handlers/UserService"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	router := mux.NewRouter()

	// UserService
	router.HandleFunc("/api/users/hello", UserService.HelloWorld).Methods("GET")
	router.HandleFunc("/api/users/login", UserService.Login).Methods("POST") // login
	router.HandleFunc("/api/users", UserService.AddUser).Methods("POST")     // register

	// ChargerService
	router.HandleFunc("/api/chargers", ChargerService.AddCharger).Methods("POST")
	router.HandleFunc("/api/chargers", ChargerService.GetAllChargers).Methods("GET")
	router.HandleFunc("/api/chargers/search", ChargerService.SearchChargers).Methods("POST")

	log.Fatal(http.ListenAndServe(":50000", router))
}
