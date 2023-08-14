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

func FindUserByUsernameAndPassword(username string, password string) (entities.User, error) {
	var user entities.User

	err := db.Db.Table("users").Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return user, errors.New(fmt.Sprintf("wrong password for username '%s'", user.Username))
	}

	if user.Banned {
		return user, errors.New("user with given credentials is banned")
	}

	return user, nil
}

func FindUserByUsername(username string) (entities.User, error) {
	var user entities.User

	err := db.Db.Table("users").Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, fmt.Errorf("there is no user by search condition username='%s'", username)
	}

	return user, nil
}

func CheckUserOwnership(username string, vehicleId uuid.UUID) (entities.User, error) {
	var user entities.User

	err := db.Db.Table("users").Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, errors.New("invalid username")
	}

	var vehicle entities.Vehicle
	err = db.Db.Table("vehicles").Where("user_id = ? AND id = ?", user.Id, vehicleId).Find(&vehicle).Error
	if err != nil {
		return user, errors.New("user with given username isn't owner of given vehicle")
	}

	user.Vehicles = append(user.Vehicles, vehicle)

	return user, nil
}

func CreateVehicle(vehicleDTO entities.VehicleDto, userId uuid.UUID) (*entities.Vehicle, error) {
	vehicle := entities.Vehicle{
		Id:          uuid.NewV4(),
		Name:        vehicleDTO.Name,
		VehicleType: entities.StrToVehicleType(vehicleDTO.VehicleType),
		UserID:      userId,
	}

	if err := utils.CheckVehicleInfo(vehicle); err != nil {
		return nil, err
	}

	if result := db.Db.Create(&vehicle); result.Error != nil {
		return nil, result.Error
	}

	return &vehicle, nil
}

func GetAllVehicles(userId uuid.UUID) []entities.Vehicle {
	var vehicles []entities.Vehicle

	db.Db.Table("vehicles").Where("user_id = ?", userId).Find(&vehicles)

	return vehicles
}

func GetAllUsers() []entities.User {
	var users []entities.User
	db.Db.Table("users").Where("role = 2").Find(&users)
	return users
}

func DeleteVehicle(name string) error {

	var vehicle entities.Vehicle

	err := db.Db.Table("vehicles").Where("name = ?", name).Find(&vehicle).Error
	if err != nil {
		return errors.New("vehicle with given name doesn't exist")
	}

	return db.Db.Delete(&vehicle).Error
}

func StrikeUser(username string) (string, error) {
	var user entities.User

	err := db.Db.Table("users").Where("username = ?", username).First(&user).Error

	if err != nil {
		return "", errors.New("invalid username")
	}
	if user.Strikes == 3 {
		return "", errors.New("this user is already banned by 3 committed strikes")
	}

	if user.Strikes+1 == 3 {
		user.Strikes += 1
		user.Banned = true
		user.BannedAt = uint64(time.Now().UnixMilli())
		user.BannedUntil = uint64(time.Now().AddDate(0, 1, 0).UnixMilli())
		db.Db.Save(user)
		return "user is banned during 3 strikes", nil
	} else {
		user.Strikes += 1
		db.Db.Save(user)
		if user.Strikes == 1 {
			return fmt.Sprintf("user %s has now 1 strike", user.Username), nil
		} else {
			return fmt.Sprintf("user %s has now %d strikes", user.Username, user.Strikes), nil
		}
	}
}
