package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"user_service/pkg/db"
	"user_service/pkg/handlers"
)

const UrlBase = "/api/usr"

func HandleRequests(port int) {

	db.Init()

	router := mux.NewRouter()

	// Auth endpoint
	router.HandleFunc(UrlBase+"/auth", handlers.Auth).Methods("GET")
	router.HandleFunc(UrlBase+"/auth/login", handlers.Login).Methods("POST")
	router.HandleFunc(UrlBase+"/auth/register", handlers.Registration).Methods("POST")

	// User endpoints
	router.HandleFunc(UrlBase+"/users", handlers.FindAllUsers).Methods("GET")

	router.HandleFunc(UrlBase+"/users/exist/{id}/{vehicleId}", handlers.CheckIfUserExistWithVehicle).Methods("GET")
	router.HandleFunc(UrlBase+"/users/exist/{id}", handlers.CheckIfUserExist).Methods("GET")

	router.HandleFunc(UrlBase+"/users/{id}", handlers.GetUsersInfo).Methods("GET")
	router.HandleFunc(UrlBase+"/users/strike/{id}/{recension_id}", handlers.StrikeUser).Methods("GET")

	// Vehicles endpoints
	router.HandleFunc(UrlBase+"/vehicles", handlers.AddVehicle).Methods("POST")
	router.HandleFunc(UrlBase+"/vehicles", handlers.GetVehicles).Methods("GET")
	router.HandleFunc(UrlBase+"/vehicles/{id}", handlers.DeleteVehicle).Methods("DELETE")

	fmt.Println("UserService is running on port: " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
