package utils

import (
	"errors"
	"regexp"
	"user_service/pkg/entities"
)

const EmailRegex = "^[a-zA-Z\\d_+&*-]+(?:\\.[a-zA-Z\\d_+&*-]+)*@(?:[a-zA-Z\\d-]+\\.)+[a-zA-Z]{2,7}$"
const PassRegex = "^(?=.*[A-Za-z])(?=.*\\d)(?=.*[$@$!%*#?&])[A-Za-z\\d$@$!%*#?&]{8,}$"

func CheckUsersInfo(user entities.User) error {

	if !checkRegexPattern(EmailRegex, user.Email) {
		return errors.New("email not matching pattern user@example.com")
	}

	if len(user.Password) == 0 {
		return errors.New("password is empty")
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

func CheckVehicleInfo(vehicle entities.Vehicle) error {

	if len(vehicle.Name) == 0 {
		return errors.New("name is empty")
	}

	return nil
}
