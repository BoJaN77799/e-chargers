package utils

import (
	"errors"
	"reservation_service/pkg/entities"
)

const (
	minReservationDuration = 15
	maxReservationDuration = 90
)

func CheckReservationsInfo(reservation *entities.ReservationDTO, dates DateFromTo) error {
	duration := dates.DateTo.Sub(dates.DateFrom)
	minutes := int(duration.Minutes())

	if minutes < minReservationDuration {
		return errors.New("minimum charging duration is 15 minutes")
	}

	if minutes > maxReservationDuration {
		return errors.New("maximum charging duration is 90 minutes")
	}

	return nil
}
