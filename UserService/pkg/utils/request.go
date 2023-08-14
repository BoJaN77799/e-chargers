package utils

import (
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func GetIdFromPathParams(r *http.Request) (uuid.UUID, error) {
	params := mux.Vars(r)
	paramId, exists := params["id"]
	if !exists {
		return uuid.UUID{}, fmt.Errorf("id is missing from path params")
	}

	id, err := uuid.FromString(paramId)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("invalid id (uuid) in path params")
	}
	return id, nil
}
