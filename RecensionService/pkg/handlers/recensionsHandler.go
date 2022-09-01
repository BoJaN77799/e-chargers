package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"recension_service/pkg/db/repository"
	"recension_service/pkg/models"
	"recension_service/pkg/utils"
	"strconv"
)

func AddRecension(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var recensionDTO models.RecensionDTO
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

	var recensionsDTO []models.RecensionDTO

	recensions := repository.GetAllRecensions()

	for _, recension := range recensions {
		recensionsDTO = append(recensionsDTO, recension.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(recensionsDTO)
}

func FindAllRecensionsFromUser(w http.ResponseWriter, r *http.Request) {

	var recensionsDTO []models.RecensionDTO

	params := mux.Vars(r)
	username, _ := params["username"]

	recensions := repository.GetAllRecensionsFromUser(username)

	for _, recension := range recensions {
		recensionsDTO = append(recensionsDTO, recension.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(recensionsDTO)
}

func CancelRecension(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var cancelRecension models.CancelRecension
	json.Unmarshal(body, &cancelRecension)

	err = repository.CancelRecension(cancelRecension.Username, cancelRecension.ChargerId)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode("recension successfully canceled")
}

func FindAllRecensionsOfCharger(w http.ResponseWriter, r *http.Request) {

	var recensionsDTO []models.RecensionDTO
	params := mux.Vars(r)
	charger, _ := params["charger_id"]

	chargerId, err := strconv.ParseUint(charger, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recensions := repository.GetAllRecensionsOfCharger(uint(chargerId))

	for _, recension := range recensions {
		recensionsDTO = append(recensionsDTO, recension.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(recensionsDTO)
}

func BanRecension(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	recensionId, _ := params["recension_id"]

	recensionIdUint, err := strconv.ParseUint(recensionId, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = repository.BanRecension(uint(recensionIdUint))

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode("recension successfully canceled")
}
