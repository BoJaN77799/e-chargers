package ReportsService

import (
	"ApiGateway/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func FindAllReservationsInPeriod(w http.ResponseWriter, r *http.Request) {
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
