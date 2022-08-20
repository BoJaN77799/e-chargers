package utils

import (
	"io"
	"net/http"
	"net/url"

	roundRobinScheduler "github.com/hlts2/round-robin"
)

var BaseUserServicePath, _ = roundRobinScheduler.New(
	&url.URL{Host: "http://localhost:50001/api/users"},
)

var BaseChargerServicePath, _ = roundRobinScheduler.New(
	&url.URL{Host: "http://localhost:50002/api/chargers"},
)

var BaseReservationServicePath, _ = roundRobinScheduler.New(
	&url.URL{Host: "http://localhost:50003/api/reservations"},
)

var BaseReportsServicePath, _ = roundRobinScheduler.New(
	&url.URL{Host: "http://localhost:50003/api/reports"},
)

func DelegateResponse(response *http.Response, w http.ResponseWriter) {
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}

func SetupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
