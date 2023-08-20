package ReservationService

import (
	"ApiGateway/pkg/handlers"
	"ApiGateway/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func AddReservation(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	URL := utils.BaseReservationServicePath.Next().Host + "/reservations"
	response, err := handlers.DoRequestWithToken(r, http.MethodPost, URL, r.Body)

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

	URL := utils.BaseReservationServicePath.Next().Host + "/reservations"
	response, err := handlers.DoRequestWithToken(r, http.MethodGet, URL, nil)

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
	userId, exists := params["userId"]
	if !exists {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	URL := utils.BaseReservationServicePath.Next().Host + "/reservations/" + userId
	response, err := handlers.DoRequestWithToken(r, http.MethodGet, URL, nil)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func CancelReservation(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	URL := utils.BaseReservationServicePath.Next().Host + "/" + id
	response, err := handlers.DoRequestWithToken(r, http.MethodDelete, URL, nil)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
