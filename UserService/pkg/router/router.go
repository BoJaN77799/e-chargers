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

func HandleRequests(port int) {

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
	router.HandleFunc("/api/users/strike/{username}/{recension_id}", handlers.StrikeUser).Methods("GET")

	// authorization
	router.HandleFunc("/api/users/auth/admin", handlers.AuthAdmin).Methods("GET")
	router.HandleFunc("/api/users/auth/user", handlers.AuthUser).Methods("GET")

	fmt.Println("UserService is running on port: " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
