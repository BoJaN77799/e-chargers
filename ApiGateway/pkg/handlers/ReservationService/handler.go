package ReservationService

import (
	"ApiGateway/pkg/models/ReservationService"
	"ApiGateway/pkg/utils"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func AddReservation(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var chargerDTO ReservationService.ReservationDTO
	data, _ := ioutil.ReadAll(r.Body)
	json.NewDecoder(bytes.NewReader(data)).Decode(&chargerDTO)

	req, _ := http.NewRequest(http.MethodPost, utils.BaseReservationServicePath.Next().Host, bytes.NewReader(data))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func FindAllReservations(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	response, err := http.Get(utils.BaseReservationServicePath.Next().Host)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func FindAllReservationsFromUser(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	username, _ := params["username"]

	response, err := http.Get(utils.BaseReservationServicePath.Next().Host + "/" + username)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func CancelReservation(w http.ResponseWriter, r *http.Request) {

	// TODO Implement CancelReservation.... ne moze preko DTO mora preko pathVariabli
	//utils.SetupResponse(&w, r)
	//if r.Method == "OPTIONS" {
	//	return
	//}
	//
	//var cancelReservationDTO ReservationService.CancelReservationDTO
	//data, _ := ioutil.ReadAll(r.Body)
	//json.NewDecoder(bytes.NewReader(data)).Decode(&cancelReservationDTO)
	//
	//req, _ := http.NewRequest(http.MethodDelete, utils.BaseReservationServicePath.Next().Host, cancelReservationDTO)
	//req.Header.Set("Accept", "application/json")
	//req.Header.Set("Content-Type", "application/json")
	//client := &http.Client{}
	//response, err := client.Do(req)
	//
	//if err != nil {
	//	w.WriteHeader(http.StatusGatewayTimeout)
	//	return
	//}
	//
	//utils.DelegateResponse(response, w)
}
