package UserService

import (
	"ApiGateway/pkg/utils"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("Api aaaa")
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
	response, err := http.Get(utils.BaseUserServicePath.Next().Host + "/hello")

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
