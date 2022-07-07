package router

import (
	"log"
	"net/http"
	"user_service/pkg/db"
	"user_service/pkg/handlers"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	DB := db.Init()

	h := handlers.New(DB)

	router := mux.NewRouter()
	router.HandleFunc("/api/users/hello", handlers.HelloWorld).Methods("GET")
	router.HandleFunc("/api/users", h.AddUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":50001", router))
}
