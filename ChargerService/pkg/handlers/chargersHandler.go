package handlers

import (
	"charger_service/pkg/db/repository"
	"charger_service/pkg/models"
	"charger_service/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	utils.OKResponse(w)
	json.NewEncoder(w).Encode("Hello World from UserService")
}

func AddCharger(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var charger models.Charger
	json.Unmarshal(body, &charger)

	_, err = repository.CreateCharger(charger)
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
}

func FindAllChargers(w http.ResponseWriter, r *http.Request) {

	var chargersDTO []models.ChargerDTO

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

	var searchDTO models.SearchDTO
	json.Unmarshal(body, &searchDTO)

	var chargersDTO []models.ChargerDTO

	chargers := repository.SearchChargers(searchDTO)

	for _, charger := range chargers {
		chargersDTO = append(chargersDTO, charger.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(chargersDTO)
}

func CheckIfExistCharger(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	chargerId, _ := params["chargerId"]

	chargerIdUint, err := strconv.ParseUint(chargerId, 10, 32)
	if err != nil {
		utils.BadRequestResponse(w, "chargerId isn't proper uint")
		return
	}

	charger := repository.GetChargerById(uint(chargerIdUint))

	if charger.ID == 0 {
		utils.BadRequestResponse(w, "charger with given id doesn't exist")
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(charger.Capacity)
}

func FindChargerReport(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	chargerId, _ := params["chargerId"]

	chargerIdUint, err := strconv.ParseUint(chargerId, 10, 32)
	if err != nil {
		utils.BadRequestResponse(w, "chargerId isn't proper uint")
		return
	}

	charger := repository.GetChargerById(uint(chargerIdUint))

	if charger.ID == 0 {
		utils.BadRequestResponse(w, "charger with given id doesn't exist")
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(charger.ToReportDTO())
}
