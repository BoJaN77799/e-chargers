package router

import (
	"fmt"
	"log"
	"net/http"
	"recension_service/pkg/db"
	"recension_service/pkg/handlers"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleRequests(port int) {

	db.Init() // initialized data base

	router := mux.NewRouter() // init router

	router.HandleFunc("/api/recensions", handlers.AddRecension).Methods("POST")
	router.HandleFunc("/api/recensions", handlers.FindAllRecensions).Methods("GET")
	router.HandleFunc("/api/recensions/{username}", handlers.FindAllRecensionsFromUser).Methods("GET")
	router.HandleFunc("/api/recensions", handlers.CancelRecension).Methods("DELETE")
	router.HandleFunc("/api/recensions/{recension_id}", handlers.BanRecension).Methods("DELETE")
	router.HandleFunc("/api/recensions/charger/{charger_id}", handlers.FindAllRecensionsOfCharger).Methods("GET")

	fmt.Println("RecensionService is running on port: " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
