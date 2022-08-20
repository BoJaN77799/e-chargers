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

	if user.ID == 0 {
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

func FindUserByUsername(username string) (models.User, error) {
	var user models.User

	db.Db.Table("users").Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return user, errors.New("invalid username")
	}

	return user, nil
}

func CheckUserOwnership(username string, vehicleId uint) (models.User, error) {
	var user models.User

	db.Db.Table("users").Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return user, errors.New("invalid username")
	}

	var vehicle models.Vehicle
	db.Db.Table("vehicles").Where("user_id = ? AND id = ?", user.ID, vehicleId).Find(&vehicle)

	if vehicle.ID == 0 {
		return user, errors.New("user with given username isn't owner of given vehicle")
	}

	user.Vehicles = append(user.Vehicles, vehicle)

	return user, nil
}

func CreateVehicle(vehicleDTO models.VehicleDTO) (models.Vehicle, error) {

	var err error

	var user models.User
	user, err = FindUserByUsername(vehicleDTO.Username)

	var vehicle models.Vehicle
	vehicle.Name = vehicleDTO.Name
	vehicle.VehicleType = models.ParseString(vehicleDTO.VehicleType)
	vehicle.UserID = user.ID

	if err != nil {
		return vehicle, errors.New("user with given username doesn't exist")
	}

	err = utils.CheckVehicleInfo(vehicle)
	if err != nil {
		return vehicle, err
	}

	if result := db.Db.Create(&vehicle); result.Error != nil {
		return vehicle, result.Error
	}

	return vehicle, nil
}

func GetAllVehicles(userId uint) []models.Vehicle {
	var vehicles []models.Vehicle

	db.Db.Table("vehicles").Where("user_id = ?", userId).Find(&vehicles)

	return vehicles
}

func DeleteVehicle(name string) error {

	var vehicle models.Vehicle

	db.Db.Table("vehicles").Where("name = ?", name).Find(&vehicle)

	if vehicle.ID <= 0 {
		return errors.New("vehicle with given name doesn't exist")
	}

	db.Db.Delete(&vehicle)

	return nil
}
