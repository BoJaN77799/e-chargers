package RecensionService

import (
	"ApiGateway/pkg/handlers"
	"ApiGateway/pkg/utils"
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
	chargerId, _ := params["charger_id"]

	URL := utils.BaseRecensionsServicePath.Next().Host + "/recensions/" + chargerId
	response, err := handlers.DoRequestWithToken(r, http.MethodGet, URL, nil)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
