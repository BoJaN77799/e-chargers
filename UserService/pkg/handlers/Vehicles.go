package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"user_service/pkg/db/repository"
	"user_service/pkg/entities"
	"user_service/pkg/utils"
)

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

	id, err := utils.GetUserIDFromToken(r)
	if err != nil {
		utils.BadToken(w, err.Error())
		json.NewEncoder(w).Encode(vehiclesDTO)
		return
	}

	user, err := repository.FindUserById(id)
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

	id, err := utils.GetIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	err = repository.DeleteVehicle(id)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode("vehicle successfully deleted")
}
