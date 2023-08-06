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

func HandleRequests(port int) {

	db.Init() // initialized data base

	router := mux.NewRouter() // init router
	router.HandleFunc("/api/chargers/hello", handlers.HelloWorld).Methods("GET")

	router.HandleFunc("/api/chargers", handlers.AddCharger).Methods("POST")
	router.HandleFunc("/api/chargers", handlers.FindAllChargers).Methods("GET")
	router.HandleFunc("/api/chargers/search", handlers.SearchChargers).Methods("POST")
	router.HandleFunc("/api/chargers/exist/{chargerId}", handlers.CheckIfExistCharger).Methods("GET")
	router.HandleFunc("/api/chargers/report/{chargerId}", handlers.FindChargerReport).Methods("GET")
	router.HandleFunc("/api/chargers/{chargerId}", handlers.GetChargerByID).Methods("GET")
	router.HandleFunc("/api/chargers/{lon}/{lat}", handlers.FindClosestCharger).Methods("GET")

	fmt.Println("ChargerService is running on port: " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
