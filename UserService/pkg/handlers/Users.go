package handlers

import (
	"encoding/json"
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

func FindUsersByUserIds(w http.ResponseWriter, r *http.Request) {
	var usersBatch entities.UsersBatchDTO
	json.NewDecoder(r.Body).Decode(&usersBatch)

	var usersDTO []entities.UserBaseInfoDTO
	users := repository.GetUsersByIds(usersBatch.UserIds)
	for _, user := range users {
		usersDTO = append(usersDTO, user.ToUserBaseInfoDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(usersDTO)
}

func GetUserWithVehicle(w http.ResponseWriter, r *http.Request) {

	userId, err := utils.GetIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	vehicleId, err := utils.GetVehicleIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	user, err := repository.GetUserVehicleByIdAndUserId(vehicleId, userId)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(user.ToUserReservationDTO())
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
