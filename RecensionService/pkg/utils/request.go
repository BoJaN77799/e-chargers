package utils

import (
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func getFromPathParams(r *http.Request, param string) (uuid.UUID, error) {
	params := mux.Vars(r)
	paramId, exists := params[param]
	if !exists {
		return uuid.UUID{}, fmt.Errorf("%s is missing from path params", param)
	}

	id, err := uuid.FromString(paramId)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("invalid %s (uuid) in path params", param)
	}
	return id, nil
}

func GetIdFromPathParams(r *http.Request) (uuid.UUID, error) {
	return getFromPathParams(r, "id")
}

func GetUserIdFromPathParams(r *http.Request) (uuid.UUID, error) {
	return getFromPathParams(r, "userId")
}

func GetChargerIdFromPathParams(r *http.Request) (uuid.UUID, error) {
	return getFromPathParams(r, "chargerId")
}
