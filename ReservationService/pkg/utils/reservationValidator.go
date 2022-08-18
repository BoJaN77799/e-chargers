package utils

import (
	"errors"
	"reservation_service/pkg/models"
	"time"
)

func CheckReservationsInfo(reservation *models.ReservationDTO) error {

	if len(reservation.Username) == 0 {
		return errors.New("reservation username is empty")
	}

	if reservation.ChargerId <= 0 {
		return errors.New("reservation chargerId is empty")
	}

	if reservation.VehicleId <= 0 {
		return errors.New("reservation vehicleId is empty")
	}

	// can't make reservation in past and 1 and half hour from now
	if reservation.DateFrom < uint64(time.Now().UnixMilli()+5400000) {
		return errors.New("invalid reservation date (less than 1h 30min from now)")
	}

	if reservation.DateFrom == 0 {
		return errors.New("reservation date is empty")
	}

	if reservation.Duration < 15 {
		return errors.New("minimum charging duration is 15 minutes")
	}

	if reservation.Duration > 90 {
		return errors.New("maximum charging duration is 90 minutes")
	}

	return nil
}
