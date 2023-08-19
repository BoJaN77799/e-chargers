package UserService

import (
	"ApiGateway/pkg/handlers"
	"ApiGateway/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
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

	URL := utils.BaseUserServicePath.Next().Host + "/users/" + id
	response, err := handlers.DoRequestWithToken(r, http.MethodGet, URL, nil)
	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func StrikeUser(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	id, _ := params["id"]
	recensionId, _ := params["recension_id"]

	URL := utils.BaseUserServicePath.Next().Host + "/strike/" + id + "/" + recensionId
	response, err := handlers.DoRequestWithToken(r, http.MethodPost, URL, nil)
	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	URL := utils.BaseUserServicePath.Next().Host + "/users"
	response, err := handlers.DoRequestWithToken(r, http.MethodGet, URL, nil)
	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
