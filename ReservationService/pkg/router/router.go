package router

import (
	"log"
	"net/http"
	"reservation_service/pkg/db"
	"reservation_service/pkg/handlers"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	db.Init() // initialized data base

	router := mux.NewRouter() // init router

	router.HandleFunc("/api/reservations", handlers.AddReservation).Methods("POST")
	router.HandleFunc("/api/reservations", handlers.FindAllReservations).Methods("GET")
	router.HandleFunc("/api/reservations/{username}", handlers.FindAllReservationsFromUser).Methods("GET")
	router.HandleFunc("/api/reservations", handlers.CancelReservation).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":50003", router))
}
