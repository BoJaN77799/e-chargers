package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"user_service/pkg/db/repository"
	"user_service/pkg/models"
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

func Login(w http.ResponseWriter, r *http.Request) {

	var loginDTO models.LoginDTO
	json.NewDecoder(r.Body).Decode(&loginDTO)

	user, err := repository.FindUserByUsernameAndPassword(loginDTO.Username, loginDTO.Password)

	if err != nil {
		if strings.Contains(err.Error(), "banned") {
			utils.UnauthorizedResponse(w)
			json.NewEncoder(w).Encode(fmt.Sprintf("you are banned by strikes until %d", user.BannedUntil))
			return
		}
		utils.UnauthorizedResponse(w)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := models.Claims{Email: user.Email, Username: user.Username, Role: user.Role.String(), Id: user.ID, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, _ := token.SignedString(utils.SECRET)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	utils.OKResponse(w)
	json.NewEncoder(w).Encode(models.UserTokenState{Token: tokenString, ExpiredAt: expirationTime.String()})
}

func Registration(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	_, err = repository.CreateUser(user)
	if err != nil {
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
