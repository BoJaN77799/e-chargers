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

	router.HandleFunc("/api/users/login", UserService.Login).Methods("POST") // login
	router.HandleFunc("/api/users", UserService.AddUser).Methods("POST")     // register
	// USER - AUTH
	router.HandleFunc("/api/users/vehicles", UserService.AddVehicle).Methods("POST")
	// USER - AUTH
	router.HandleFunc("/api/users/vehicles/{username}", UserService.GetVehicles).Methods("GET")
	// USER - AUTH
	router.HandleFunc("/api/users/vehicles/{name}", UserService.DeleteVehicle).Methods("DELETE")
	// ADMIN - AUTH
	router.HandleFunc("/api/users/strike/{username}/{recension_id}", UserService.StrikeUser).Methods("GET")
	// ADMIN - AUTH
	router.HandleFunc("/api/users", UserService.FindAllUsers).Methods("GET")
	// USER - AUTH
	router.HandleFunc("/api/users/{username}", UserService.GetUserInfo).Methods("GET")

	// ChargerService
	// ADMIN - AUTH
	router.HandleFunc("/api/chargers", ChargerService.AddCharger).Methods("POST")
	// FREE
	router.HandleFunc("/api/chargers", ChargerService.GetAllChargers).Methods("GET")
	// FREE
	router.HandleFunc("/api/chargers/search", ChargerService.SearchChargers).Methods("POST")
	// NO-USE
	router.HandleFunc("/api/chargers/{chargerId}", ChargerService.GetChargerById).Methods("GET")
	// FREE
	router.HandleFunc("/api/chargers/{lon}/{lat}", ChargerService.FindClosestCharger).Methods("GET")

	// ReservationService
	// USER - AUTH
	router.HandleFunc("/api/reservations", ReservationService.AddReservation).Methods("POST")
	// NO-USE
	router.HandleFunc("/api/reservations", ReservationService.FindAllReservations).Methods("GET")
	// USER - AUTH
	router.HandleFunc("/api/reservations/{username}", ReservationService.FindAllReservationsFromUser).Methods("GET")
	// USER - AUTH
	router.HandleFunc("/api/reservations/{id}", ReservationService.CancelReservation).Methods("DELETE")

	// ReportsService
	// ADMIN - AUTH
	router.HandleFunc("/api/reports/{date_from}/{date_to}", ReportsService.FindAllReservationsInPeriod).Methods("GET")
	// ADMIN - AUTH
	router.HandleFunc("/api/reports/users", ReportsService.FindAllUsersReport).Methods("GET")

	//RecensionsService
	// USER - AUTH
	router.HandleFunc("/api/recensions", RecensionService.AddRecension).Methods("POST")
	// FREE
	router.HandleFunc("/api/recensions/charger/{charger_id}", RecensionService.FindAllRecensionsOfCharger).Methods("GET")

	log.Fatal(http.ListenAndServe(":50000", router))
}
