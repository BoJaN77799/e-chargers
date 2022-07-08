package repository

import (
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
