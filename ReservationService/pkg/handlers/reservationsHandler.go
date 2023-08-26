package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"reservation_service/pkg/db/repository"
	"reservation_service/pkg/entities"
	"reservation_service/pkg/utils"
)

func AddReservation(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var reservationDTO entities.ReservationDTO
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

	var reservationsDTO []entities.ReservationDTO

	reservations := repository.GetAllReservations()

	for _, reservation := range reservations {
		reservationsDTO = append(reservationsDTO, reservation.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(reservationsDTO)
}

func FindAllReservationsFromUser(w http.ResponseWriter, r *http.Request) {

	var reservationsDTO []entities.ReservationDTO

	userId, err := utils.GetIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	reservations := repository.GetAllReservationsFromUser(userId)

	for _, reservation := range reservations {
		reservationsDTO = append(reservationsDTO, reservation.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(reservationsDTO)
}

func CancelReservation(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetIdFromPathParams(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	err = repository.CancelReservation(id)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode("reservation successfully canceled")
}

func FindAllReservationsInPeriod(w http.ResponseWriter, r *http.Request) {

	var reservationsDTO []entities.ReservationDTO

	params := mux.Vars(r)
	dateFrom, _ := params["date_from"]
	dateTo, _ := params["date_to"]

	dates, err := utils.ConvertDateFromAndDateTo(dateFrom, dateTo)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	reservations := repository.GetAllReservationsInPeriod(dates.DateFrom, dates.DateFrom)

	for _, reservation := range reservations {
		reservationsDTO = append(reservationsDTO, reservation.ToDTO())
	}

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(reservationsDTO)
}
