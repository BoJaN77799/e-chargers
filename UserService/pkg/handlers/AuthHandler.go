package handlers

import (
	"net/http"
	"user_service/pkg/db/repository"
	"user_service/pkg/utils"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	username, err := utils.GetUsernameFromToken(r)
	if err != nil {
		utils.UnauthorizedResponse(w)
		return
	}
	_, err = repository.FindUserByUsername(username)
	if err != nil {
		utils.UnauthorizedResponse(w)
		return
	}
}
