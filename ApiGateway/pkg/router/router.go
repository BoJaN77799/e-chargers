package router

import (
	"ApiGateway/pkg/handlers/UserService"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	router := mux.NewRouter()

	// UserService
	router.HandleFunc("/api/users/hello", UserService.HelloWorld).Methods("GET")
	router.HandleFunc("/api/users", UserService.AddUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":50000", router))
}
