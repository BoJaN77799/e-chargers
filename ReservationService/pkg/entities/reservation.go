package entities

import (
	uuid "github.com/satori/go.uuid"
	"reservation_service/pkg/entities/base"
	"time"
)

type Reservation struct {
	Id          uuid.UUID `json:"id" gorm:"primary_key;type:uuid"`
	UserId      uuid.UUID `json:"userId" gorm:"not-null"`
	ChargerId   uuid.UUID `json:"charger_id" gorm:"not-null"`
	ChargerName string    `json:"charger_name" gorm:"not-null"`
	VehicleId   uuid.UUID `json:"vehicle_id" gorm:"not-null"`
	VehicleName string    `json:"vehicle_name" gorm:"not-null"`
	DateFrom    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DateTo      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	base.Base
}

func (reservation *Reservation) ToDTO() ReservationDTO {
	return ReservationDTO{
		Id:          reservation.Id,
		UserId:      reservation.UserId,
		ChargerId:   reservation.ChargerId,
		DateFrom:    reservation.DateFrom.Format(time.RFC3339),
		DateTo:      reservation.DateFrom.Format(time.RFC3339),
		ChargerName: reservation.ChargerName,
		VehicleName: reservation.VehicleName,
		VehicleId:   reservation.VehicleId,
	}
}
