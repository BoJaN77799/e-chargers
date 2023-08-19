package repository

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"recension_service/pkg/db"
	"recension_service/pkg/entities"
	"recension_service/pkg/utils"
)

func CreateRecension(recensionDTO entities.RecensionDTO) (entities.Recension, error) {
	err := utils.CheckRecensionsInfo(&recensionDTO)
	if err != nil {
		return entities.Recension{}, err
	}

	err = VerifyUserUsername(recensionDTO.UserId)
	if err != nil {
		return entities.Recension{}, err
	}

	err = VerifyCharger(recensionDTO.ChargerId)
	if err != nil {
		return entities.Recension{}, err
	}

	result, err := GetRecensionToxicity(recensionDTO.Content)
	if err != nil {
		return entities.Recension{}, err
	}

	var recension = entities.Recension{
		Id:        uuid.NewV4(),
		UserId:    recensionDTO.UserId,
		ChargerId: recensionDTO.ChargerId,
		Content:   recensionDTO.Content,
		Rate:      recensionDTO.Rate,
		Toxic:     result[1],
	}

	if result := db.Db.Create(&recension); result.Error != nil {
		return recension, result.Error
	}

	return recension, nil
}

func GetAllRecensions() []entities.Recension {
	var recensions []entities.Recension
	db.Db.Table("recensions").Find(&recensions)
	return recensions
}

func GetAllRecensionsFromUser(userId uuid.UUID) []entities.Recension {
	var recensions []entities.Recension
	db.Db.Table("recensions").Where("userId = ?", userId).Find(&recensions)
	return recensions
}

func CancelRecension(id uuid.UUID) error {
	recension, err := getRecensionByID(id)
	if err != nil {
		return err
	}

	return db.Db.Delete(recension).Error
}

func getRecensionByID(id uuid.UUID) (entities.Recension, error) {
	var recension entities.Recension
	err := db.Db.Table("recensions").Where("id = ?", id).Find(&recension).Error
	if err != nil {
		return entities.Recension{}, fmt.Errorf("recension doesn't exist by search condition id='%s'", id)
	}
	return recension, nil
}

func GetAllRecensionsOfCharger(chargerId uuid.UUID) []entities.Recension {
	var recensions []entities.Recension
	db.Db.Table("recensions").Where("charger_id = ? AND banned = false", chargerId).Find(&recensions)
	return recensions
}

func BanRecension(id uuid.UUID) error {
	recension := entities.Recension{
		Id:     id,
		Banned: true,
	}

	if err := db.Db.Update(&recension); err != nil {
		return fmt.Errorf("failed to make banned recension with id='%s'", id)
	}

	return nil
}
