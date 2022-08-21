package router

import (
	"ApiGateway/pkg/handlers/ChargerService"
	"ApiGateway/pkg/handlers/RecensionService"
	"ApiGateway/pkg/handlers/ReportsService"
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
	router.HandleFunc("/api/users/strike/{username}", UserService.StrikeUser).Methods("GET")
	router.HandleFunc("/api/users", UserService.FindAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{username}", UserService.GetUserInfo).Methods("GET")

	// ChargerService
	router.HandleFunc("/api/chargers", ChargerService.AddCharger).Methods("POST")
	router.HandleFunc("/api/chargers", ChargerService.GetAllChargers).Methods("GET")
	router.HandleFunc("/api/chargers/search", ChargerService.SearchChargers).Methods("POST")
	router.HandleFunc("/api/chargers/{chargerId}", ChargerService.GetChargerById).Methods("GET")

	// ReservationService
	router.HandleFunc("/api/reservations", ReservationService.AddReservation).Methods("POST")
	router.HandleFunc("/api/reservations", ReservationService.FindAllReservations).Methods("GET")
	router.HandleFunc("/api/reservations/{username}", ReservationService.FindAllReservationsFromUser).Methods("GET")
	router.HandleFunc("/api/reservations/{id}", ReservationService.CancelReservation).Methods("DELETE")

	// ReportsService
	router.HandleFunc("/api/reports/{date_from}/{date_to}", ReportsService.FindAllReservationsInPeriod).Methods("GET")
	router.HandleFunc("/api/reports/users", ReportsService.FindAllUsersReport).Methods("GET")

	//RecensionsService
	router.HandleFunc("/api/recensions", RecensionService.AddRecension).Methods("POST")
	router.HandleFunc("/api/recensions/charger/{charger_id}", RecensionService.FindAllRecensionsOfCharger).Methods("GET")

	log.Fatal(http.ListenAndServe(":50000", router))
}
