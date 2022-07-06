package utils

import (
	"errors"
	"regexp"
	"user_service/pkg/models"
)

func CheckUsersInfo(user models.User) error {

	result, _ := regexp.MatchString("^[a-zA-Z\\d_+&*-]+(?:\\.[a-zA-Z\\d_+&*-]+)*@(?:[a-zA-Z\\d-]+\\.)+[a-zA-Z]{2,7}$", user.Email)

	if !result {
		return errors.New("email not matching pattern user@example.com")
	}

	if len(user.Username) == 0 {
		return errors.New("username is empty")
	}

	if len(user.Password) == 0 {
		return errors.New("password is empty")
	}

	return nil
}
