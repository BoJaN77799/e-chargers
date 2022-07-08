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

	db.Db.Preload("Plugs").Preload("Address").Find(&chargers)

	return chargers
}

// func FindUserByUsernameAndPassword(chargername string, password string) (models.User, error) {
// 	var charger models.User

// 	db.Db.Table("chargers").Where("chargername = ?", chargername).First(&charger)

// 	if charger.Id == 0 {
// 		return charger, errors.New("invalid chargername")
// 	}

// 	if !utils.CheckPasswordHash(password, charger.Password) {
// 		return charger, errors.New(fmt.Sprintf("wrong password for chargername '%s'", charger.Username))
// 	}

// 	//if time.Now().Before(charger.BannedUntil) {
// 	//	return charger, errors.New("You are banned until: " + charger.BannedUntil.String())
// 	//}

// 	return charger, nil
// }
