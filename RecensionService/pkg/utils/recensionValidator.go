package utils

import (
	"errors"
	"recension_service/pkg/entities"
)

func CheckRecensionsInfo(recension *entities.RecensionDTO) error {
	if len(recension.Content) == 0 {
		return errors.New("recension content is missing")
	}

	if recension.Rate < 1 {
		return errors.New("recension rate is minimum 1")
	}

	if recension.Rate > 5 {
		return errors.New("recension rate is maximum 5")
	}

	return nil
}
