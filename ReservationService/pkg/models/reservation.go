package models

import (
	"github.com/jinzhu/gorm"
)

type Reservation struct {
	gorm.Model
	Username  string `json:"username" gorm:"not-null"`
	ChargerId uint   `json:"charger_id" gorm:"not-null"`
	VehicleId uint   `json:"vehicle_id" gorm:"not-null"`
	DateFrom  uint64 `json:"date_from"  gorm:"not-null"`
	DateTo    uint64 `json:"date_to" gorm:"not-null"`
}

func (reservation *Reservation) ToDTO() ReservationDTO {
	return ReservationDTO{
		Username:  reservation.Username,
		ChargerId: reservation.ChargerId,
		DateFrom:  reservation.DateFrom,
		Duration:  (reservation.DateTo - reservation.DateFrom) * 1000 * 60,
	}
}

type CancelReservation struct {
	Username  string `json:"username"`
	ChargerId uint   `json:"charger_id"`
	VehicleId uint   `json:"vehicle_id"`
}
