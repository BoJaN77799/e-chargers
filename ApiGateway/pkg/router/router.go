package router

import (
	"ApiGateway/pkg/handlers/ChargerService"
	"ApiGateway/pkg/handlers/ReservationService"
	"ApiGateway/pkg/handlers/UserService"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {

	router := mux.NewRouter()

	// UserService
	router.HandleFunc("/api/users/hello", UserService.HelloWorld).Methods("GET")
	router.HandleFunc("/api/users/login", UserService.Login).Methods("POST") // login
	router.HandleFunc("/api/users", UserService.AddUser).Methods("POST")     // register
	router.HandleFunc("/api/users/vehicles", UserService.AddVehicle).Methods("POST")
	router.HandleFunc("/api/users/vehicles/{username}", UserService.GetVehicles).Methods("GET")
	router.HandleFunc("/api/users/vehicles/{name}", UserService.DeleteVehicle).Methods("DELETE")

	// ChargerService
	router.HandleFunc("/api/chargers", ChargerService.AddCharger).Methods("POST")
	router.HandleFunc("/api/chargers", ChargerService.GetAllChargers).Methods("GET")
	router.HandleFunc("/api/chargers/search", ChargerService.SearchChargers).Methods("POST")

	// ReservationService
	router.HandleFunc("/api/reservations", ReservationService.AddReservation).Methods("POST")
	router.HandleFunc("/api/reservations", ReservationService.FindAllReservations).Methods("GET")
	router.HandleFunc("/api/reservations/{username}", ReservationService.FindAllReservationsFromUser).Methods("GET")
	router.HandleFunc("/api/reservations", ReservationService.CancelReservation).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":50000", router))
}
