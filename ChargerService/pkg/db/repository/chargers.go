package repository

import (
	"charger_service/pkg/db"
	"charger_service/pkg/entities"
	"charger_service/pkg/utils"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math"
)

const MinDistance = 500.0

func CreateCharger(charger entities.Charger) (entities.Charger, error) {

	err := utils.CheckChargersInfo(&charger)
	if err != nil {
		return charger, err
	}

	if result := db.Db.Create(&charger); result.Error != nil {
		return charger, result.Error
	}

	return charger, nil
}

func GetAllChargers() []entities.Charger {
	var chargers []entities.Charger

	db.Db.Preload("Address").Find(&chargers)

	return chargers
}

func SearchChargers(search entities.SearchDTO) []entities.Charger {

	var chargers []entities.Charger

	db.Db.Preload("Address").Where(
		"name like ?"+
			"AND NOT ((work_time_from >= ? AND work_time_from >= ?) OR (work_time_to <= ? AND work_time_to <= ?))"+
			"AND capacity <= ?"+
			"AND price_per_hour BETWEEN ? AND  ?"+
			"AND charging_speed_per_minute BETWEEN ? AND ?"+
			"AND plugs like ?",
		"%"+search.Name+"%",
		search.WorkTimeFrom,
		search.WorkTimeTo,
		search.WorkTimeFrom,
		search.WorkTimeTo,
		search.Capacity,
		search.PricePerHourFrom,
		search.PricePerHourTo,
		search.ChargingSpeedFrom,
		search.ChargingSpeedTo,
		"%"+search.Type+"%").Find(&chargers)

	return chargers
}

func GetChargerById(id uuid.UUID) (entities.Charger, error) {
	var charger entities.Charger
	err := db.Db.Preload("Address").Where("id = ?", id).Find(&charger).Error
	return charger, err
}

func GetClosestChargerToCoordinates(longitude float64, latitude float64) (*entities.Charger, error) {
	var chargers []entities.Charger

	var closestCharger *entities.Charger
	var minDistance = MinDistance
	chargerFound := false

	for _, charger := range chargers {
		distance := math.Sqrt(
			math.Pow(charger.Address.Longitude-longitude, 2) +
				math.Pow(charger.Address.Latitude-latitude, 2))

		if distance < minDistance {
			minDistance = distance
			closestCharger = &charger
			chargerFound = true
		}
	}

	if !chargerFound {
		return nil, fmt.Errorf("no charger found within %fm diameter", MinDistance)
	}

	return closestCharger, nil
}
