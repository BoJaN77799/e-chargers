package router

import (
	"log"
	"net/http"
	"recension_service/pkg/db"
	"recension_service/pkg/handlers"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	db.Init() // initialized data base

	router := mux.NewRouter() // init router

	router.HandleFunc("/api/recensions", handlers.AddRecension).Methods("POST")
	router.HandleFunc("/api/recensions", handlers.FindAllRecensions).Methods("GET")
	router.HandleFunc("/api/recensions/{username}", handlers.FindAllRecensionsFromUser).Methods("GET")
	router.HandleFunc("/api/recensions", handlers.CancelRecension).Methods("DELETE")
	router.HandleFunc("/api/recensions/charger/{charger_id}", handlers.FindAllRecensionsOfCharger).Methods("GET")

	log.Fatal(http.ListenAndServe(":50005", router))
}
