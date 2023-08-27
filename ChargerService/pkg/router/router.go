package router

import (
	"charger_service/pkg/db"
	"charger_service/pkg/handlers"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const UrlBase = "/api/chr"

func HandleRequests(port int) {
	db.Init()

	router := mux.NewRouter()
	router.HandleFunc(UrlBase+"/chargers", handlers.FindAllChargers).Methods("GET")

	router.HandleFunc(UrlBase+"/chargers", handlers.AddCharger).Methods("POST")
	router.HandleFunc(UrlBase+"/chargers/search", handlers.SearchChargers).Methods("POST")
	router.HandleFunc(UrlBase+"/chargers/{id}", handlers.GetChargerByID).Methods("GET")
	router.HandleFunc(UrlBase+"/chargers/{lon}/{lat}", handlers.FindClosestCharger).Methods("GET")
	router.HandleFunc(UrlBase+"/chargers/report/{id}", handlers.FindChargerReport).Methods("GET")
	router.HandleFunc(UrlBase+"/chargers/reservation/{id}", handlers.CheckIfExistCharger).Methods("GET")

	fmt.Println("ChargerService is running on port: " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
