package router

import (
	"charger_service/pkg/db"
	"charger_service/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	db.Init() // initialized data base

	router := mux.NewRouter() // init router
	router.HandleFunc("/api/chargers/hello", handlers.HelloWorld).Methods("GET")

	router.HandleFunc("/api/chargers", handlers.AddCharger).Methods("POST")
	router.HandleFunc("/api/chargers", handlers.FindAllChargers).Methods("GET")
	router.HandleFunc("/api/chargers/search", handlers.SearchChargers).Methods("POST")

	router.HandleFunc("/api/chargers/exist/{chargerId}", handlers.CheckIfExistCharger).Methods("GET")

	log.Fatal(http.ListenAndServe(":50002", router))
}
