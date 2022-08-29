package ReportsService

import (
	"ApiGateway/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func FindAllReservationsInPeriod(w http.ResponseWriter, r *http.Request) {

	//// auth
	//if err := utils.Authorize(r, "admin"); err != nil {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusUnauthorized)
	//	json.NewEncoder(w).Encode(err.Error())
	//	return
	//}

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	dateFrom, _ := params["date_from"]
	dateTo, _ := params["date_to"]

	response, err := http.Get(utils.BaseReportsServicePath.Next().Host + "/chargers/" + dateFrom + "/" + dateTo)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func FindAllUsersReport(w http.ResponseWriter, r *http.Request) {

	//// auth
	//if err := utils.Authorize(r, "admin"); err != nil {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusUnauthorized)
	//	json.NewEncoder(w).Encode(err.Error())
	//	return
	//}

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	response, err := http.Get(utils.BaseReportsServicePath.Next().Host + "/users")

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
