package repository

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"user_service/pkg/db"
	"user_service/pkg/entities"
	"user_service/pkg/utils"
)

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

func GetVehicleByIdAndUserId(id uuid.UUID, userId uuid.UUID) (entities.Vehicle, error) {
	var vehicle entities.Vehicle
	err := db.Db.Table("vehicles").Where("id = ? AND user_id = ?", id, userId).First(&vehicle).Error
	if err != nil {
		return vehicle, fmt.Errorf("there is no vehicle by search condition (id='%s' AND user_id='%s')", id, userId)
	}

	return vehicle, nil
}

func GetAllVehicles(userId uuid.UUID) []entities.Vehicle {
	var vehicles []entities.Vehicle
	db.Db.Table("vehicles").Where("user_id = ?", userId).Find(&vehicles)

	return vehicles
}

func DeleteVehicle(id uuid.UUID) error {
	var vehicle entities.Vehicle
	err := db.Db.Table("vehicles").Where("id = ?", id).Find(&vehicle).Error
	if err != nil {
		return errors.New("vehicle doesn't exist")
	}

	return db.Db.Delete(&vehicle).Error
}
