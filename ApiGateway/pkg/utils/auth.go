package utils

import (
	"bytes"
	"errors"
	"net/http"
	"time"
)

// SECRET This should be stored as an environment variable
var SECRET = []byte("my_ultra_secret_key")

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
