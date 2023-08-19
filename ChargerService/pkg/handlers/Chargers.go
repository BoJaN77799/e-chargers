package handlers

import (
	"charger_service/pkg/db/repository"
	"charger_service/pkg/entities"
	"charger_service/pkg/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	utils.OKResponse(w)
	json.NewEncoder(w).Encode("Hello World from UserService")
}

func AddCharger(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var charger entities.Charger
	json.Unmarshal(data, &charger)

	_, err = repository.CreateCharger(charger)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode("charger successfully created")
}

func FindAllChargers(w http.ResponseWriter, r *http.Request) {

	var chargersDTO []entities.ChargerDTO

	chargers := repository.GetAllChargers()

	for _, charger := range chargers {
		chargersDTO = append(chargersDTO, charger.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(chargersDTO)
}

func SearchChargers(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var searchDTO entities.SearchDTO
	json.Unmarshal(body, &searchDTO)

	var chargersDTO []entities.ChargerDTO

	chargers := repository.SearchChargers(searchDTO)

	for _, charger := range chargers {
		chargersDTO = append(chargersDTO, charger.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(chargersDTO)
}

func CheckIfExistCharger(w http.ResponseWriter, r *http.Request) {
	id, _ := utils.GetIdFromPathParams(r)

	charger, err := repository.GetChargerById(id)

	if err != nil {
		utils.BadRequestResponse(w, "charger with given id doesn't exist")
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(charger.ToReservationDTO())
}

func FindChargerReport(w http.ResponseWriter, r *http.Request) {
	id, _ := utils.GetIdFromPathParams(r)

	charger, err := repository.GetChargerById(id)

	if err != nil {
		utils.BadRequestResponse(w, "charger with given id doesn't exist")
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(charger.ToReportDTO())
}

func GetChargerByID(w http.ResponseWriter, r *http.Request) {
	id, _ := utils.GetIdFromPathParams(r)

	charger, err := repository.GetChargerById(id)

	if err != nil {
		utils.BadRequestResponse(w, "charger with given id doesn't exist")
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(charger.ToDTO())
}

func FindClosestCharger(w http.ResponseWriter, r *http.Request) {

	longitude, _ := utils.GetLonFromPathParams(r)
	latitude, _ := utils.GetLatFromPathParams(r)

	charger, err := repository.GetClosestChargerToCoordinates(longitude, latitude)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(charger.ToDTO())
}
