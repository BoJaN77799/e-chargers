package utils

import (
	"errors"
	"regexp"
	"user_service/pkg/models"
)

const EmailRegex = "^[a-zA-Z\\d_+&*-]+(?:\\.[a-zA-Z\\d_+&*-]+)*@(?:[a-zA-Z\\d-]+\\.)+[a-zA-Z]{2,7}$"
const PassRegex = "^(?=.*[A-Za-z])(?=.*\\d)(?=.*[$@$!%*#?&])[A-Za-z\\d$@$!%*#?&]{8,}$"

func CheckUsersInfo(user models.User) error {

	if !checkRegexPattern(EmailRegex, user.Email) {
		return errors.New("email not matching pattern user@example.com")
	}

	if len(user.Username) == 0 {
		return errors.New("username is empty")
	}

	if len(user.Password) == 0 {
		return errors.New("password is empty")
	}

	if len(user.Username) < 3 && len(user.Username) > 20 {
		return errors.New("username is minimum 3 chars, maximum 20 chars")
	}

	//if !checkRegexPattern(PassRegex, user.Password) {
	//	return errors.New("password should be at least 8 characters long and should contain one number, one character and one special character")
	//}

	return nil
}

func checkRegexPattern(regex string, value string) bool {
	result, _ := regexp.MatchString(regex, value)
	return result
}

func CheckVehicleInfo(vehicle models.Vehicle) error {

	if len(vehicle.Name) == 0 {
		return errors.New("name is empty")
	}

	return nil
}
