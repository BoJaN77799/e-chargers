package UserService

import (
	"ApiGateway/pkg/models"
	"ApiGateway/pkg/utils"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}
	response, err := http.Get(utils.BaseUserServicePath.Next().Host + "/hello")

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

	var userDTO models.UserDTO
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
