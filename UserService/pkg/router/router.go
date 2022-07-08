package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user_service/pkg/db"
	"user_service/pkg/handlers"
)

func HandleRequests() {

	db.Init() // initialized data base

	router := mux.NewRouter() // init router
	router.HandleFunc("/api/users/hello", handlers.HelloWorld).Methods("GET")

	router.HandleFunc("/api/users", handlers.AddUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":50001", router))
}
