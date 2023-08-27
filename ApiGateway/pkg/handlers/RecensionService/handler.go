package RecensionService

import (
	"ApiGateway/pkg/handlers"
	"ApiGateway/pkg/handlers/UserService"
	"ApiGateway/pkg/models/RecensionService"
	"ApiGateway/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func AddRecension(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	URL := utils.BaseRecensionsServicePath.Next().Host + "/recensions"
	response, err := handlers.DoRequestWithToken(r, http.MethodPost, URL, r.Body)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func FindAllRecensionsOfCharger(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	chargerId, _ := params["chargerId"]

	URL := utils.BaseRecensionsServicePath.Next().Host + "/recensions/charger/" + chargerId
	response, err := handlers.DoRequest(http.MethodGet, URL, nil)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	defer response.Body.Close()
	var recensions []RecensionService.RecensionDTO
	if err := json.NewDecoder(response.Body).Decode(&recensions); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	recensionsWithUser, err := UserService.GetUsersBatch(recensions)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseData, err := json.Marshal(recensionsWithUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(responseData)

	utils.DelegateResponse(response, w)
}
