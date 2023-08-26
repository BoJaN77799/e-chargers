package repository

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"reservation_service/pkg/db"
	"reservation_service/pkg/entities"
	"reservation_service/pkg/utils"
	"time"
)

func CreateReservation(reservationDTO entities.ReservationDTO) (entities.Reservation, error) {

	dates, err := utils.ConvertDateFromAndDateTo(reservationDTO.DateFrom, reservationDTO.DateTo)
	if err != nil {
		return entities.Reservation{}, err
	}

	err = utils.CheckReservationsInfo(&reservationDTO, dates)
	if err != nil {
		return entities.Reservation{}, err
	}

	userWithVehicle, err := VerifyUserUsernameAndVehicle(reservationDTO.UserId, reservationDTO.VehicleId)
	if err != nil {
		return entities.Reservation{}, err
	}

	charger, err := GetChargerById(reservationDTO.ChargerId)
	if err != nil {
		return entities.Reservation{}, err
	}

	var reservation = entities.Reservation{
		UserId:      reservationDTO.UserId,
		ChargerId:   charger.Id,
		ChargerName: charger.Name,
		VehicleId:   userWithVehicle.VehicleId,
		VehicleName: userWithVehicle.VehicleName,
		DateFrom:    dates.DateFrom,
		DateTo:      dates.DateTo,
	}

	err = checkIfReservationExist(&reservation)
	if err != nil {
		return reservation, err
	}

	err = checkIfReservationOnOtherChargersExist(&reservation)
	if err != nil {
		return reservation, err
	}

	err = checkChargerCapacity(&reservation, charger.Capacity)
	if err != nil {
		return reservation, err
	}

	if result := db.Db.Create(&reservation); result.Error != nil {
		return reservation, result.Error
	}

	return reservation, nil
}

func checkIfReservationExist(reservation *entities.Reservation) error {
	var count int64

	db.Db.Table("reservations").
		Where("user_id = ?", reservation.UserId).
		Where("charger_id = ?", reservation.ChargerId).
		Where("vehicle_id = ?", reservation.VehicleId).
		Where("NOT (date_from >= ? OR date_to <= ?)", reservation.DateTo, reservation.DateFrom).
		Count(&count)

	if count > 0 {
		return errors.New("user has another reservation on the same charger with the same vehicle (charging periods overlap)")
	}
	return nil
}

func checkIfReservationOnOtherChargersExist(reservation *entities.Reservation) error {
	var count int64

	db.Db.Table("reservations").Where(
		"user_id = ? AND vehicle_id = ?"+
			"AND NOT (date_from >= ? OR date_to <= ?)",
		reservation.UserId,
		reservation.VehicleId,
		reservation.DateTo,
		reservation.DateFrom,
	).Count(&count)

	if count > 0 {
		return errors.New("the user has already reserved the charging of this vehicle at another charger")
	}

	return nil
}

func checkChargerCapacity(reservation *entities.Reservation, chargerCapacity uint) error {
	var count int64

	db.Db.Table("reservations").Where(
		"charger_id = ? AND NOT (date_from >= ? OR date_to <= ?)",
		reservation.ChargerId,
		reservation.DateTo,
		reservation.DateFrom,
	).Count(&count)

	if uint(count)+1 > chargerCapacity {
		return errors.New("there is no free slot on this charger right now")
	}
	return nil
}

func GetAllReservations() []entities.Reservation {
	var reservations []entities.Reservation
	db.Db.Table("reservations").Find(&reservations)
	return reservations
}

func GetAllReservationsFromUser(userId uuid.UUID) []entities.Reservation {
	var reservations []entities.Reservation
	db.Db.Table("reservations").Where("user_id = ?", userId).Find(&reservations)
	return reservations
}

func CancelReservation(id uuid.UUID) error {
	var reservation entities.Reservation

	err := db.Db.Table("reservations").Where("id = ?", id).Find(&reservation).Error

	if err != nil {
		return errors.New("user with given username doesn't have any reservations on charger with given id with this vehicle")
	}

	db.Db.Delete(reservation)

	return nil
}

func GetAllReservationsInPeriod(dateFrom, dateTo time.Time) []entities.Reservation {
	var reservations []entities.Reservation

	db.Db.Table("reservations").Where(
		"NOT ((date_from >= ? AND date_from >= ?) OR (date_to <= ? AND date_to <= ?))",
		dateFrom,
		dateTo,
		dateFrom,
		dateTo,
	).Find(&reservations)

	return reservations
}
