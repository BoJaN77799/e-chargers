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

const UrlBase = "/api/rec"

func HandleRequests(port int) {

	db.Init()

	router := mux.NewRouter()

	router.HandleFunc(UrlBase+"/recensions", handlers.AddRecension).Methods("POST")
	router.HandleFunc(UrlBase+"/recensions", handlers.FindAllRecensions).Methods("GET")
	router.HandleFunc(UrlBase+"/recensions/{userId}", handlers.FindAllRecensionsFromUser).Methods("GET")
	router.HandleFunc(UrlBase+"/recensions/{id}", handlers.CancelRecension).Methods("DELETE")
	router.HandleFunc(UrlBase+"/recensions/{id}", handlers.BanRecension).Methods("POST")
	router.HandleFunc(UrlBase+"/recensions/charger/{charger_id}", handlers.FindAllRecensionsOfCharger).Methods("GET")

	fmt.Println("RecensionService is running on port: " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
