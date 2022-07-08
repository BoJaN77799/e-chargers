package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"user_service/pkg/db/repository"
	"user_service/pkg/models"
	"user_service/pkg/utils"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	utils.OKResponse(w)
	json.NewEncoder(w).Encode("Hello World from UserService")
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	_, err = repository.CreateUser(user)
	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "username") {
			utils.BadRequestResponse(w, "user with given username already exists")
		}
		if strings.Contains(err.Error(), "email") {
			utils.BadRequestResponse(w, "user with given email already exists")
		}
		return
	}

	utils.CreatedResponse(w)
	json.NewEncoder(w).Encode("user successfully created")
}
