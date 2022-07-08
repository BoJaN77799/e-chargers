package repository

import (
	"errors"
	"fmt"
	"user_service/pkg/db"
	"user_service/pkg/models"
	"user_service/pkg/utils"
)

func CreateUser(user models.User) (models.User, error) {

	err := utils.CheckUsersInfo(user)
	if err != nil {
		return user, err
	}

	// registered user
	user.Role = models.RegisteredUser

	// hashing password
	user.Password, _ = utils.HashPassword(user.Password)

	if result := db.Db.Create(&user); result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func FindUserByUsernameAndPassword(username string, password string) (models.User, error) {
	var user models.User

	db.Db.Table("users").Where("username = ?", username).First(&user)

	if user.Id == 0 {
		return user, errors.New("invalid username")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return user, errors.New(fmt.Sprintf("wrong password for username '%s'", user.Username))
	}

	//if time.Now().Before(user.BannedUntil) {
	//	return user, errors.New("You are banned until: " + user.BannedUntil.String())
	//}

	return user, nil
}
