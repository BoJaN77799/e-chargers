package utils

import (
	"errors"
	"recension_service/pkg/models"
)

func CheckRecensionsInfo(recension *models.RecensionDTO) error {

	if len(recension.Username) == 0 {
		return errors.New("recension username is empty")
	}

	if recension.ChargerId <= 0 {
		return errors.New("recension chargerId is empty")
	}

	if recension.Date == 0 {
		return errors.New("date is missing")
	}

	if len(recension.Content) == 0 {
		return errors.New("recension content is missing")
	}

	if recension.Rate < 1 {
		return errors.New("recenzion rate is minimum 1")
	}

	if recension.Rate > 5 {
		return errors.New("recension rate is maximum 5")
	}

	return nil
}
