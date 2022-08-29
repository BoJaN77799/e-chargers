package utils

import (
	"bytes"
	"errors"
	"net/http"
	"time"
)

func Authorize(req *http.Request, role string) error {
	bearer := req.Header["Authorization"]
	if bearer == nil {
		return errors.New("no authorization")
	}

	token := bearer[0]

	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}

	request, _ := http.NewRequest(http.MethodGet, BaseUserServicePath.Next().Host+"/auth/"+role, bytes.NewBufferString(""))
	request.Header.Set("Authorization", token)

	response, err := client.Do(request)
	if err != nil {
		return errors.New("error sending request")
	}

	if response.StatusCode != 200 {
		return errors.New("unauthorized")
	}

	return nil
}

func Authenticate(req *http.Request) (int, error) {
	//bearer := req.Header["Authorization"]
	//if bearer == nil {
	//	return http.StatusBadRequest, errors.New("no token bearer")
	//}
	//tokenStr := strings.Split(bearer[0], " ")[1]
	//
	//timeout := time.Duration(5 * time.Second)
	//client := http.Client{Timeout: timeout}
	//
	//endpoint := UserServiceRoot + _AuthenticateApi
	//request, _ := http.NewRequest(http.MethodGet, endpoint, bytes.NewBufferString(""))
	//request.Header.Set("Authorization", "Bearer "+tokenStr)
	//
	//response, err := client.Do(request)
	//if err != nil {
	//	return http.StatusGatewayTimeout, errors.New("error sending request")
	//}
	//
	//if response.StatusCode != 200 {
	//	return http.StatusUnauthorized, errors.New("not authenticated")
	//}

	return http.StatusOK, nil
}
