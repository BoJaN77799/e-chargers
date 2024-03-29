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

	router.HandleFunc("/api/users/login", handlers.Login).Methods("POST")
	router.HandleFunc("/api/users", handlers.AddUser).Methods("POST")
	router.HandleFunc("/api/users", handlers.FindAllUsers).Methods("GET")
	router.HandleFunc("/api/users/vehicles", handlers.AddVehicle).Methods("POST")
	router.HandleFunc("/api/users/vehicles/{username}", handlers.GetVehicles).Methods("GET")
	router.HandleFunc("/api/users/vehicles/{name}", handlers.DeleteVehicle).Methods("DELETE")

	router.HandleFunc("/api/users/exist/{username}/{vehicleId}", handlers.CheckIfUserExistWithVehicle).Methods("GET")
	router.HandleFunc("/api/users/exist/{username}", handlers.CheckIfUserExist).Methods("GET")

	router.HandleFunc("/api/users/{username}", handlers.GetUsersInfo).Methods("GET")
	router.HandleFunc("/api/users/strike/{username}", handlers.StrikeUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":50001", router))
}
