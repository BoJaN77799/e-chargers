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
	"strconv"
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

func FindAllReservationsInPeriod(w http.ResponseWriter, r *http.Request) {

	var reservationsDTO []models.ReservationDTO

	params := mux.Vars(r)
	dateFrom, _ := params["date_from"]
	dateTo, _ := params["date_to"]

	dateFromUInt64, err := strconv.ParseUint(dateFrom, 10, 64)

	if err != nil {
		utils.BadRequestResponse(w, "dateFrom is not valid")
		return
	}
	dateToUInt64, err := strconv.ParseUint(dateTo, 10, 64)

	if err != nil {
		utils.BadRequestResponse(w, "dateTo is not valid")
		return
	}
	reservations := repository.GetAllReservationsInPeriod(dateFromUInt64, dateToUInt64)

	for _, reservation := range reservations {
		reservationsDTO = append(reservationsDTO, reservation.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(reservationsDTO)
}
