package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"reservation_service/pkg/db/repository"
	"reservation_service/pkg/models"
	"reservation_service/pkg/utils"
)

func AddReservation(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var reservationDTO models.ReservationDTO
	json.Unmarshal(body, &reservationDTO)

	_, err = repository.CreateReservation(reservationDTO)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}
	utils.OKResponse(w)
	json.NewEncoder(w).Encode("reservation successfully created")
}

func FindAllReservations(w http.ResponseWriter, r *http.Request) {

	var reservationsDTO []models.ReservationDTO

	reservations := repository.GetAllReservations()

	for _, reservation := range reservations {
		reservationsDTO = append(reservationsDTO, reservation.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(reservationsDTO)
}

func FindAllReservationsFromUser(w http.ResponseWriter, r *http.Request) {

	var reservationsDTO []models.ReservationDTO

	params := mux.Vars(r)
	username, _ := params["username"]

	reservations := repository.GetAllReservationsFromUser(username)

	for _, reservation := range reservations {
		reservationsDTO = append(reservationsDTO, reservation.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(reservationsDTO)
}

func CancelReservation(w http.ResponseWriter, r *http.Request) {

	//params := mux.Vars(r)
	//username, _ := params["username"]
	//charger, _ := params["chargerId"]
	//chargerId, err := strconv.ParseUint(charger, 10, 32)
	//
	//vehicle, _ := params["vehicleId"]
	//vehicleId, err := strconv.ParseUint(vehicle, 10, 32)

	//if err != nil {
	//	utils.BadRequestResponse(w, "charger id isn't proper uint")
	//	return
	//}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var cancelReservation models.CancelReservation
	json.Unmarshal(body, &cancelReservation)

	err = repository.CancelReservation(cancelReservation.Username, cancelReservation.ChargerId, cancelReservation.VehicleId)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode("reservation successfully canceled")
}
