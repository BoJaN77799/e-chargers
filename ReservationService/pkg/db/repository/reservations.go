package repository

import (
	"errors"
	"fmt"
	"reservation_service/pkg/db"
	"reservation_service/pkg/models"
	"reservation_service/pkg/utils"
)

func CreateReservation(reservationDTO models.ReservationDTO) (models.Reservation, error) {

	var reservation models.Reservation

	err := utils.CheckReservationsInfo(&reservationDTO)
	if err != nil {
		return reservation, err
	}

	// verify user username and his vehicle id
	err = VerifyUserUsernameAndVehicle(reservationDTO.Username, reservationDTO.VehicleId)
	fmt.Println("aaaa")
	if err != nil {
		return reservation, err
	}

	// verify charger id and get charger capacity
	var capacity uint
	capacity, err = VerifyChargerId(reservationDTO.ChargerId)
	if err != nil {
		return reservation, err
	}

	reservation.Username = reservationDTO.Username
	reservation.ChargerId = reservationDTO.ChargerId
	reservation.VehicleId = reservationDTO.VehicleId
	reservation.DateFrom = reservationDTO.DateFrom
	reservation.DateTo = reservationDTO.DateFrom + reservationDTO.Duration*1000*60

	err = CheckIfReservationExist(&reservation)

	if err != nil {
		return reservation, err
	}

	err = CheckIfReservationOnOtherChargersExist(&reservation)

	if err != nil {
		return reservation, err
	}

	err = CheckChargerCapacity(&reservation, capacity)

	if err != nil {
		return reservation, err
	}

	if result := db.Db.Create(&reservation); result.Error != nil {
		return reservation, result.Error
	}

	return reservation, nil
}

func CheckIfReservationExist(reservation *models.Reservation) error {
	var reservationDB models.Reservation

	db.Db.Table("reservations").Where(
		"username = ? "+
			"AND charger_id = ? "+
			"AND vehicle_id = ?"+
			"AND NOT ((date_from >= ? AND date_from >= ?) OR (date_to <= ? AND date_to <= ?))",
		reservation.Username,
		reservation.ChargerId,
		reservation.VehicleId,
		reservation.DateFrom,
		reservation.DateTo,
		reservation.DateFrom,
		reservation.DateTo,
	).Find(&reservationDB)

	if reservationDB.ID != 0 {
		return errors.New("user has 2 reservation on same charger with same vehicle (charging periods overlaps)")
	}

	// TODO user with vehicle on another charger
	// TODO provjeri da li je korisnik rezervisao na nekom drugom mjestu u isto vrijeme

	return nil
}

func CheckIfReservationOnOtherChargersExist(reservation *models.Reservation) error {
	var reservationDB models.Reservation

	db.Db.Table("reservations").Where(
		"username = ? "+
			"AND vehicle_id = ?"+
			"AND NOT ((date_from >= ? AND date_from >= ?) OR (date_to <= ? AND date_to <= ?))",
		reservation.Username,
		reservation.VehicleId,
		reservation.DateFrom,
		reservation.DateTo,
		reservation.DateFrom,
		reservation.DateTo,
	).Find(&reservationDB)

	if reservationDB.ID != 0 {
		// TODO nadji tacno koji charger
		return errors.New("the user has already reserved the charging of this vehicle at other charger")
	}

	return nil
}

func CheckChargerCapacity(reservation *models.Reservation, chargerCapacity uint) error {
	var reservationsDB []models.Reservation

	db.Db.Table("reservations").Where(
		"charger_id = ? "+
			"AND NOT ((date_from >= ? AND date_from >= ?) OR (date_to <= ? AND date_to <= ?))",
		reservation.ChargerId,
		reservation.DateFrom,
		reservation.DateTo,
		reservation.DateFrom,
		reservation.DateTo,
	).Find(&reservationsDB)

	if uint(len(reservationsDB))+1 > chargerCapacity {
		// TODO vrati prijedlog prvog slobodnog termina i ponudi da zakaze tad
		return errors.New("there is no free slot on this charger right now")
	}
	return nil
}

func GetAllReservations() []models.Reservation {
	var reservations []models.Reservation

	db.Db.Table("reservations").Find(&reservations)

	return reservations
}

func GetAllReservationsFromUser(username string) []models.Reservation {
	var reservations []models.Reservation

	db.Db.Table("reservations").Where("username = ?", username).Find(&reservations)

	return reservations
}

func CancelReservation(username string, chargerId uint, vehicleId uint) error {
	var reservation models.Reservation

	db.Db.Table("reservations").Where("username = ? AND charger_id = ? AND vehicle_id = ?", username, chargerId, vehicleId).Find(&reservation)

	if reservation.ID == 0 {
		return errors.New("user with given username doesn't have any reservations on charger with given id with this vehicle")
	}

	db.Db.Delete(reservation)

	return nil
}
