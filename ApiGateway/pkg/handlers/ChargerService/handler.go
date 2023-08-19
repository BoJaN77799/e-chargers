package ChargerService

import (
	"ApiGateway/pkg/handlers"
	"ApiGateway/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func AddCharger(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	URL := utils.BaseChargerServicePath.Next().Host + "/chargers"
	response, err := handlers.DoRequestWithToken(r, http.MethodPost, URL, r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func GetAllChargers(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	URL := utils.BaseChargerServicePath.Next().Host + "/chargers"
	response, err := handlers.DoRequestWithToken(r, http.MethodGet, URL, nil)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func SearchChargers(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	URL := utils.BaseChargerServicePath.Next().Host + "/chargers/search"
	response, err := handlers.DoRequestWithToken(r, http.MethodPost, URL, r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func GetChargerById(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	id, _ := params["id"]

	URL := utils.BaseChargerServicePath.Next().Host + "/chargers/" + id
	response, err := handlers.DoRequestWithToken(r, http.MethodGet, URL, nil)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func FindClosestCharger(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	lon, _ := params["lon"]
	lat, _ := params["lat"]

	URL := utils.BaseChargerServicePath.Next().Host + "/" + lon + "/" + lat
	response, err := handlers.DoRequestWithToken(r, http.MethodGet, URL, nil)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
