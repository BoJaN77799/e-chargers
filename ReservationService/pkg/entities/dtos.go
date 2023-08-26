package entities

import (
	uuid "github.com/satori/go.uuid"
)

type ReservationDTO struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	ChargerId   uuid.UUID `json:"charger_id"`
	ChargerName string    `json:"charger_name"`
	VehicleId   uuid.UUID `json:"vehicle_id"`
	VehicleName string    `json:"vehicle_name"`
	DateFrom    string    `json:"date_from"`
	DateTo      string    `json:"date_to"`
}

type VehicleDTO struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	VehicleType string `json:"vehicle_type"`
	Username    string `json:"username"`
}

type UserReservationDTO struct {
	Id          uuid.UUID `json:"username"`
	VehicleId   uuid.UUID `json:"vehicle_id"`
	VehicleName string    `json:"vehicle_name"`
}

type ChargerReservationDTO struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"charger_name"`
	Capacity uint      `json:"capacity"`
}
