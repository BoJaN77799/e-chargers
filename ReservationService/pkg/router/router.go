package router

import (
	"fmt"
	"log"
	"net/http"
	"reservation_service/pkg/db"
	"reservation_service/pkg/handlers"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleRequests(port int) {

	db.Init() // initialized data base

	router := mux.NewRouter() // init router

	router.HandleFunc("/api/reservations", handlers.AddReservation).Methods("POST")
	router.HandleFunc("/api/reservations", handlers.FindAllReservations).Methods("GET")
	router.HandleFunc("/api/reservations/{usedId}", handlers.FindAllReservationsFromUser).Methods("GET")
	router.HandleFunc("/api/reservations/{id}", handlers.CancelReservation).Methods("DELETE")

	router.HandleFunc("/api/reservations/{date_from}/{date_to}", handlers.FindAllReservationsInPeriod).Methods("GET")

	fmt.Println("ReservationService is running on port: " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
