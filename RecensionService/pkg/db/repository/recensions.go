package repository

import (
	"errors"
	"recension_service/pkg/db"
	"recension_service/pkg/models"
	"recension_service/pkg/utils"
	"time"
)

func CreateRecension(recensionDTO models.RecensionDTO) (models.Recension, error) {

	var recension models.Recension

	err := utils.CheckRecensionsInfo(&recensionDTO)
	if err != nil {
		return recension, err
	}

	err = VerifyUserUsername(recensionDTO.Username)

	if err != nil {
		return recension, err
	}

	err = VerifyCharger(recensionDTO.ChargerId)

	if err != nil {
		return recension, err
	}

	err = CheckIfRecensionExist(&recensionDTO)

	if err != nil {
		return recension, err
	}
	recension.Username = recensionDTO.Username
	recension.ChargerId = recensionDTO.ChargerId
	recension.Content = recensionDTO.Content
	recension.Rate = recensionDTO.Rate
	recension.Date = uint64(time.Now().UnixMilli())

	if result := db.Db.Create(&recension); result.Error != nil {
		return recension, result.Error
	}

	return recension, nil
}

func CheckIfRecensionExist(recension *models.RecensionDTO) error {
	var recensionDB models.Recension

	db.Db.Table("recensions").Where(
		"username = ? AND charger_id = ? ",
		recension.Username,
		recension.ChargerId,
	).Find(&recensionDB)

	if recensionDB.ID != 0 {
		return errors.New("user has 2 recensions on this charger")
	}
	return nil
}

func GetAllRecensions() []models.Recension {
	var recensions []models.Recension

	db.Db.Table("recensions").Find(&recensions)

	return recensions
}

func GetAllRecensionsFromUser(username string) []models.Recension {
	var recensions []models.Recension

	db.Db.Table("recensions").Where("username = ?", username).Find(&recensions)

	return recensions
}

func CancelRecension(username string, chargerId uint) error {
	var recension models.Recension

	db.Db.Table("recensions").Where("username = ? AND charger_id = ?", username, chargerId).Find(&recension)

	if recension.ID == 0 {
		return errors.New("user didn't give recension on this charger")
	}

	db.Db.Delete(recension)

	return nil
}

func GetAllRecensionsInPeriod(dateFromUInt64 uint64, dateToUInt64 uint64) []models.Recension {

	var recensions []models.Recension

	db.Db.Table("recensions").Where(
		"NOT ((date_from >= ? AND date_from >= ?) OR (date_to <= ? AND date_to <= ?))",
		dateFromUInt64,
		dateToUInt64,
		dateFromUInt64,
		dateToUInt64,
	).Find(&recensions)

	return recensions
}
