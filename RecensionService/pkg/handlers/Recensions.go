package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"recension_service/pkg/db/repository"
	"recension_service/pkg/entities"
	"recension_service/pkg/utils"
)

func AddRecension(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var recensionDTO entities.RecensionDTO
	json.Unmarshal(body, &recensionDTO)

	_, err = repository.CreateRecension(recensionDTO)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}
	utils.OKResponse(w)
	json.NewEncoder(w).Encode("recension successfully created")
}

func FindAllRecensions(w http.ResponseWriter, r *http.Request) {
	var recensionsDTO []entities.RecensionDTO
	recensions := repository.GetAllRecensions()
	for _, recension := range recensions {
		recensionsDTO = append(recensionsDTO, recension.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(recensionsDTO)
}

func FindAllRecensionsFromUser(w http.ResponseWriter, r *http.Request) {
	var recensionsDTO []entities.RecensionDTO
	userId, err := utils.GetUserIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		json.NewEncoder(w).Encode(recensionsDTO)
		return
	}

	recensions := repository.GetAllRecensionsFromUser(userId)
	for _, recension := range recensions {
		recensionsDTO = append(recensionsDTO, recension.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(recensionsDTO)
}

func CancelRecension(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	err = repository.CancelRecension(id)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode("recension successfully canceled")
}

func FindAllRecensionsOfCharger(w http.ResponseWriter, r *http.Request) {
	var recensionsDTO []entities.RecensionDTO
	chargerId, err := utils.GetChargerIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		json.NewEncoder(w).Encode(recensionsDTO)
		return
	}

	recensions := repository.GetAllRecensionsOfCharger(chargerId)
	for _, recension := range recensions {
		recensionsDTO = append(recensionsDTO, recension.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(recensionsDTO)
}

func BanRecension(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetUserIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	err = repository.BanRecension(id)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode("recension successfully canceled")
}
