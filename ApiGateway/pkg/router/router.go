package router

import (
	"ApiGateway/pkg/handlers/UserService"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const (
	Administrator    = "Administrator"
	UnregisteredUser = "UnregisteredUser"
	RegisteredUser   = "RegisteredUser"
)

func HandleRequests(port int) {
	//router := mux.NewRouter()

	//// ADMIN - AUTH
	//router.HandleFunc("/api/users/strike/{username}/{recension_id}", UserService.StrikeUser).Methods("GET")
	//// ADMIN - AUTH
	//router.HandleFunc("/api/users", UserService.FindAllUsers).Methods("GET")
	//// USER - AUTH
	//router.HandleFunc("/api/users/{username}", UserService.GetUserInfo).Methods("GET")
	//
	//// ChargerService
	//// ADMIN - AUTH
	//router.HandleFunc("/api/chargers", ChargerService.AddCharger).Methods("POST")
	//// FREE
	//router.HandleFunc("/api/chargers", ChargerService.GetAllChargers).Methods("GET")
	//// FREE
	//router.HandleFunc("/api/chargers/search", ChargerService.SearchChargers).Methods("POST")
	//// NO-USE
	//router.HandleFunc("/api/chargers/{chargerId}", ChargerService.GetChargerById).Methods("GET")
	//// FREE
	//router.HandleFunc("/api/chargers/{lon}/{lat}", ChargerService.FindClosestCharger).Methods("GET")
	//
	//// ReservationService
	//// USER - AUTH
	//router.HandleFunc("/api/reservations", ReservationService.AddReservation).Methods("POST")
	//// NO-USE
	//router.HandleFunc("/api/reservations", ReservationService.FindAllReservations).Methods("GET")
	//// USER - AUTH
	//router.HandleFunc("/api/reservations/{username}", ReservationService.FindAllReservationsFromUser).Methods("GET")
	//// USER - AUTH
	//router.HandleFunc("/api/reservations/{id}", ReservationService.CancelReservation).Methods("DELETE")
	//
	//// ReportsService
	//// ADMIN - AUTH
	//router.HandleFunc("/api/reports/{date_from}/{date_to}", ReportsService.FindAllReservationsInPeriod).Methods("GET")
	//// ADMIN - AUTH
	//router.HandleFunc("/api/reports/users", ReportsService.FindAllUsersReport).Methods("GET")
	//
	////RecensionsService
	//// USER - AUTH
	//router.HandleFunc("/api/recensions", RecensionService.AddRecension).Methods("POST")
	//// FREE
	//router.HandleFunc("/api/recensions/charger/{charger_id}", RecensionService.FindAllRecensionsOfCharger).Methods("GET")

	router := mux.NewRouter()

	router.HandleFunc("/api/auth/login", UserService.Login).Methods("POST")           // login
	router.HandleFunc("/api/auth/register", UserService.Registration).Methods("POST") // register

	router.Use(authenticationMiddleware)

	//// Protected routes
	// Administrator
	router.HandleFunc("/api/users", authorizationMiddleware(UserService.FindAllUsers, []string{Administrator})).Methods("GET")

	// Registered User
	router.HandleFunc("/api/users/vehicles", authorizationMiddleware(UserService.GetVehicles, []string{RegisteredUser})).Methods("GET")
	router.HandleFunc("/api/users/vehicles", authorizationMiddleware(UserService.AddVehicle, []string{RegisteredUser})).Methods("POST")
	router.HandleFunc("/api/users/vehicles/{id}", authorizationMiddleware(UserService.DeleteVehicle, []string{RegisteredUser})).Methods("DELETE")

	fmt.Println("ApiGateway is running on port: " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
