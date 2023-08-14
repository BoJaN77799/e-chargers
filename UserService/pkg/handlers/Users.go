package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"user_service/pkg/db/repository"
	"user_service/pkg/entities"
	"user_service/pkg/utils"
)

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	var usersDTO []entities.UserReportDTO
	users := repository.GetAllRegisteredUsers()
	for _, user := range users {
		usersDTO = append(usersDTO, user.ToReportDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(usersDTO)
}

func CheckIfUserExistWithVehicle(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	vehicleId, err := utils.GetVehicleIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	var user entities.User
	user, err = repository.CheckUserOwnership(id, vehicleId)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(user)
}

func CheckIfUserExist(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	_, err = repository.FindUserById(id)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(fmt.Sprintf("user with id: %s exist", id))
}

func GetUsersInfo(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	user, err := repository.FindUserById(id)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(user.ToUserProfileDTO())
}

func StrikeUser(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	message, err := repository.StrikeUser(id)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	params := mux.Vars(r)
	recensionId, _ := params["recension_id"]
	recensionIdUint, err := strconv.ParseUint(recensionId, 10, 32)
	err = repository.BanRecension(recensionIdUint)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}
	utils.OKResponse(w)
	json.NewEncoder(w).Encode(message)
}
