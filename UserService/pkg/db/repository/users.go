package repository

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
	"user_service/pkg/db"
	"user_service/pkg/entities"
	"user_service/pkg/utils"
)

const MaxStrikes = 3

func CreateUser(user entities.User) (entities.User, error) {

	err := utils.CheckUsersInfo(user)
	if err != nil {
		return user, err
	}

	user.Id = uuid.NewV4()
	user.Role = entities.RegisteredUser
	user.Password, _ = utils.HashPassword(user.Password)
	if result := db.Db.Create(&user); result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func FindUserByEmailAndPassword(email string, password string) (entities.User, error) {
	user, err := FindUserByEmail(email)
	if err != nil {
		return user, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return user, errors.New(fmt.Sprintf("wrong password for email '%s'", user.Email))
	}

	if user.Banned {
		return user, errors.New("user with given credentials is banned")
	}

	return user, nil
}

func FindUserByEmail(email string) (entities.User, error) {
	var user entities.User

	err := db.Db.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, fmt.Errorf("there is no user by search condition email='%s'", email)
	}

	return user, nil
}

func FindUserById(id uuid.UUID) (entities.User, error) {
	var user entities.User

	err := db.Db.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, fmt.Errorf("there is no user by search condition id='%s'", id)
	}

	return user, nil
}

func CheckUserOwnership(id uuid.UUID, vehicleId uuid.UUID) (entities.User, error) {
	user, err := FindUserById(id)
	if err != nil {
		return user, err
	}

	vehicle, err := GetVehicleByIdAndUserId(id, vehicleId)
	if err != nil {
		return user, err
	}

	user.Vehicles = append(user.Vehicles, vehicle)

	return user, nil
}

func GetAllRegisteredUsers() []entities.User {
	var users []entities.User
	db.Db.Table("users").Where("role = ?", entities.RegisteredUser).Find(&users)
	return users
}

func StrikeUser(id uuid.UUID) (string, error) {
	user, err := FindUserById(id)
	if err != nil {
		return "", err
	}

	if user.Strikes == MaxStrikes {
		return "", fmt.Errorf("this user is already banned by %d committed strikes", MaxStrikes)
	}

	if user.Strikes+1 == MaxStrikes {
		return banUser(user)
	} else {
		incrementStrikes(user)
		return getStrikeMessage(user), nil
	}
}

func banUser(user entities.User) (string, error) {
	user.Strikes++
	user.Banned = true
	user.BannedAt = uint64(time.Now().UnixMilli())
	user.BannedUntil = uint64(time.Now().AddDate(0, 1, 0).UnixMilli())
	db.Db.Save(&user)
	return fmt.Sprintf("user is banned during %d strikes", MaxStrikes), nil
}

func incrementStrikes(user entities.User) {
	user.Strikes++
	db.Db.Save(&user)
}

func getStrikeMessage(user entities.User) string {
	switch user.Strikes {
	case 1:
		return fmt.Sprintf("user id='%s' has now 1 strike", user.Id)
	default:
		return fmt.Sprintf("user id='%s' has now %d strikes", user.Id, user.Strikes)
	}
}
