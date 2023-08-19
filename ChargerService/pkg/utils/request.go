package utils

import (
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

func getUUIDFromPathParams(r *http.Request, param string) (uuid.UUID, error) {
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

func getFloat32FromPathParams(r *http.Request, param string) (float64, error) {
	params := mux.Vars(r)
	paramId, exists := params[param]
	if !exists {
		return 0.0, fmt.Errorf("%s is missing from path params", param)
	}

	id, err := strconv.ParseFloat(paramId, 32)
	if err != nil {
		return 0.0, fmt.Errorf("invalid %s (float32) in path params", param)
	}
	return id, nil
}

func GetIdFromPathParams(r *http.Request) (uuid.UUID, error) {
	return getUUIDFromPathParams(r, "id")
}

func GetLonFromPathParams(r *http.Request) (float64, error) {
	return getFloat32FromPathParams(r, "lon")
}

func GetLatFromPathParams(r *http.Request) (float64, error) {
	return getFloat32FromPathParams(r, "lat")
}
