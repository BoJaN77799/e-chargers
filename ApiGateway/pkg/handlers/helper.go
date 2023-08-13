package handlers

import (
	"ApiGateway/pkg/utils"
	"io"
	"net/http"
	"time"
)

func DoRequestWithToken(r *http.Request, method string, URL string, body io.Reader) (*http.Response, error) {
	tokenString, err := utils.ExtractAccessTokenFromHeader(r)
	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: 5 * time.Second}
	request, _ := http.NewRequest(method, URL, body)
	request.Header.Set("Authorization", "Bearer "+tokenString)
	return client.Do(request)
}
