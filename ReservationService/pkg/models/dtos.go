package models

type ReservationDTO struct {
	Id          uint   `json:"id"`
	Username    string `json:"username"`
	ChargerId   uint   `json:"charger_id"`
	ChargerName string `json:"charger_name"`
	VehicleId   uint   `json:"vehicle_id"`
	VehicleName string `json:"vehicle_name"`
	DateFrom    uint64 `json:"date_from"`
	Duration    uint64 `json:"duration"`
}

type VehicleDTO struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	VehicleType string `json:"vehicle_type"`
	Username    string `json:"username"`
}

type UserReservationDTO struct {
	Username string `json:"username" gorm:"unique;not-null"`
	Vehicles []VehicleDTO
}

type ChargerReservationDTO struct {
	Id       uint   `json:"id"`
	Name     string `json:"charger_name"`
	Capacity uint   `json:"capacity"`
}
