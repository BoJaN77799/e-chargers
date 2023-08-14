package UserService

import (
	"ApiGateway/pkg/handlers"
	"ApiGateway/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func AddVehicle(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	URL := utils.BaseUserServicePath.Next().Host + "/vehicles"
	response, err := handlers.DoRequestWithToken(r, http.MethodPost, URL, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func GetVehicles(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	URL := utils.BaseUserServicePath.Next().Host + "/vehicles"
	response, err := handlers.DoRequestWithToken(r, http.MethodGet, URL, nil)
	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func DeleteVehicle(w http.ResponseWriter, r *http.Request) {
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

	URL := utils.BaseUserServicePath.Next().Host + "/vehicles/" + id
	response, err := handlers.DoRequestWithToken(r, http.MethodDelete, URL, nil)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
