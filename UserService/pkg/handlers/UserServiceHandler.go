package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"user_service/pkg/db/repository"
	"user_service/pkg/entities"
	"user_service/pkg/utils"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	utils.OKResponse(w)

	message := "Hello World from UserService : " + r.URL.String()
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {

	var usersDTO []entities.UserReportDTO

	users := repository.GetAllUsers()

	for _, user := range users {
		usersDTO = append(usersDTO, user.ToReportDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(usersDTO)
}

func AddVehicle(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.BadRequestResponse(w, "vehicle body is missing")
		return
	}

	var vehicleDTO entities.VehicleDto
	json.Unmarshal(body, &vehicleDTO)

	userId, err := utils.GetUserIDFromToken(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	_, err = repository.CreateVehicle(vehicleDTO, userId)
	if err != nil {
		if strings.Contains(err.Error(), "name") {
			utils.BadRequestResponse(w, "vehicle with given name already exists")
		}
		return
	}
	utils.CreatedResponse(w)
	json.NewEncoder(w).Encode("vehicle successfully created")
}

func GetVehicles(w http.ResponseWriter, r *http.Request) {
	var vehiclesDTO []entities.VehicleDto

	username, err := utils.GetUsernameFromToken(r)
	if err != nil {
		utils.BadToken(w, err.Error())
		json.NewEncoder(w).Encode(vehiclesDTO)
		return
	}

	user, err := repository.FindUserByUsername(username)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	vehicles := repository.GetAllVehicles(user.Id)
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

func CheckIfUserExistWithVehicle(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	username, _ := params["username"]
	vehicleId, _ := params["vehicleId"]

	id, err := uuid.FromString(vehicleId)
	if err != nil {
		utils.BadRequestResponse(w, "vehicleId isn't proper uuid")
	}

	var user entities.User
	user, err = repository.CheckUserOwnership(username, id)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(user)
}

func CheckIfUserExist(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	username, _ := params["username"]

	_, err := repository.FindUserByUsername(username)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(fmt.Sprintf("user with username: %s exist", username))
}

func GetUsersInfo(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	username, _ := params["username"]

	user, err := repository.FindUserByUsername(username)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(user.ToUserProfileDTO())
}

func StrikeUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	username, _ := params["username"]
	recensionId, _ := params["recension_id"]

	message, err := repository.StrikeUser(username)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	vehicleIdUint, err := strconv.ParseUint(recensionId, 10, 32)
	err = repository.BanRecension(uint(vehicleIdUint))

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}
	utils.OKResponse(w)
	json.NewEncoder(w).Encode(message)
}
