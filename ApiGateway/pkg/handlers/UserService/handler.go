package UserService

import (
	"ApiGateway/pkg/handlers"
	"ApiGateway/pkg/models/UserService"
	"ApiGateway/pkg/utils"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var loginDTO UserService.LoginDTO
	data, _ := ioutil.ReadAll(r.Body)
	json.NewDecoder(bytes.NewReader(data)).Decode(&loginDTO)

	req, _ := http.NewRequest(http.MethodPost, utils.BaseUserServicePath.Next().Host+"/auth/login", bytes.NewReader(data))
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

func AddUser(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var userDTO UserService.UserDTO
	data, _ := ioutil.ReadAll(r.Body)
	json.NewDecoder(bytes.NewReader(data)).Decode(&userDTO)

	req, _ := http.NewRequest(http.MethodPost, utils.BaseUserServicePath.Next().Host, bytes.NewReader(data))
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

func AddVehicle(w http.ResponseWriter, r *http.Request) {

	// auth
	if err := utils.Authorize(r, "user"); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var vehicleDTO UserService.VehicleDTO
	data, _ := ioutil.ReadAll(r.Body)
	json.NewDecoder(bytes.NewReader(data)).Decode(&vehicleDTO)

	req, _ := http.NewRequest(http.MethodPost, utils.BaseUserServicePath.Next().Host+"/vehicles", bytes.NewReader(data))
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

func GetVehicles(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	URL := utils.BaseUserServicePath.Next().Host + "/users/vehicles"
	response, err := handlers.DoRequestWithToken(r, http.MethodGet, URL, nil)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {

	// auth
	if err := utils.Authorize(r, "user"); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	username, _ := params["username"]

	response, err := http.Get(utils.BaseUserServicePath.Next().Host + "/" + username)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func StrikeUser(w http.ResponseWriter, r *http.Request) {

	// auth
	if err := utils.Authorize(r, "admin"); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	username, _ := params["username"]
	recensionId, _ := params["recension_id"]

	response, err := http.Get(utils.BaseUserServicePath.Next().Host + "/strike/" + username + "/" + recensionId)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {

	//// auth
	//if err := utils.Authorize(r, "user"); err != nil {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusUnauthorized)
	//	json.NewEncoder(w).Encode(err.Error())
	//	return
	//}

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	response, err := http.Get(utils.BaseUserServicePath.Next().Host)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func DeleteVehicle(w http.ResponseWriter, r *http.Request) {

	// auth
	if err := utils.Authorize(r, "user"); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	name, _ := params["name"]

	req, _ := http.NewRequest(http.MethodDelete, utils.BaseUserServicePath.Next().Host+"/vehicles/"+name, nil)
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
