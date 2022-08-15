package repository

import (
	"charger_service/pkg/db"
	"charger_service/pkg/models"
	"charger_service/pkg/utils"
)

func CreateCharger(charger models.Charger) (models.Charger, error) {

	err := utils.CheckChargersInfo(&charger)
	if err != nil {
		return charger, err
	}

	if result := db.Db.Create(&charger); result.Error != nil {
		return charger, result.Error
	}

	return charger, nil
}

func GetAllChargers() []models.Charger {
	var chargers []models.Charger

	db.Db.Preload("Address").Find(&chargers)

	return chargers
}

func SearchChargers(search models.SearchDTO) []models.Charger {

	var chargers []models.Charger

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
