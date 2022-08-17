package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"user_service/pkg/db/repository"
	"user_service/pkg/models"
	"user_service/pkg/utils"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	utils.OKResponse(w)
	json.NewEncoder(w).Encode("Hello World from UserService")
}

// This should be stored as an environment variable
var jwtKey = []byte("my_ultra_secret_key")

func Login(w http.ResponseWriter, r *http.Request) {

	var loginDTO models.LoginDTO
	json.NewDecoder(r.Body).Decode(&loginDTO)

	user, err := repository.FindUserByUsernameAndPassword(loginDTO.Username, loginDTO.Password)

	if err != nil {
		utils.UnauthorizedResponse(w)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := models.Claims{Email: user.Email, Username: user.Username, Role: user.Role.String(), Id: user.ID, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, _ := token.SignedString(jwtKey)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(models.UserTokenState{Token: tokenString, ExpiredAt: expirationTime.String()})
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	_, err = repository.CreateUser(user)
	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "username") {
			utils.BadRequestResponse(w, "user with given username already exists")
		}
		if strings.Contains(err.Error(), "email") {
			utils.BadRequestResponse(w, "user with given email already exists")
		}
		return
	}

	utils.CreatedResponse(w)
	json.NewEncoder(w).Encode("user successfully created")
}

func AddVehicle(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var vehicleDTO models.VehicleDTO
	json.Unmarshal(body, &vehicleDTO)

	_, err = repository.CreateVehicle(vehicleDTO)
	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "name") {
			utils.BadRequestResponse(w, "vehicle with given name already exists")
		}
		return
	}
	utils.CreatedResponse(w)
	json.NewEncoder(w).Encode("vehicle successfully created")
}

func GetVehicles(w http.ResponseWriter, r *http.Request) {
	var vehiclesDTO []models.VehicleDTO

	params := mux.Vars(r)
	username, _ := params["username"]

	var err error
	var user models.User
	user, err = repository.FindUserByUsername(username)

	if err != nil {
		utils.BadRequestResponse(w, "user with given username doesn't exist")
		json.NewEncoder(w).Encode(vehiclesDTO)
		return
	}

	vehicles := repository.GetAllVehicles(user.ID)

	for _, vehicle := range vehicles {
		vehiclesDTO = append(vehiclesDTO, vehicle.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(vehiclesDTO)
}

func DeleteVehicle(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	name, _ := params["name"]

	var err error
	err = repository.DeleteVehicle(name)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode("vehicle with name: " + name + " successfully deleted")
}
