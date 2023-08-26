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

const UrlBase = "/api/res"

func HandleRequests(port int) {

	db.Init()

	router := mux.NewRouter()

	router.HandleFunc(UrlBase+"/reservations", handlers.AddReservation).Methods("POST")
	router.HandleFunc(UrlBase+"/reservations", handlers.FindAllReservations).Methods("GET")
	router.HandleFunc(UrlBase+"/reservations/{usedId}", handlers.FindAllReservationsFromUser).Methods("GET")
	router.HandleFunc(UrlBase+"/reservations/{id}", handlers.CancelReservation).Methods("DELETE")
	router.HandleFunc(UrlBase+"/reservations/{date_from}/{date_to}", handlers.FindAllReservationsInPeriod).Methods("GET")

	fmt.Println("ReservationService is running on port: " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
