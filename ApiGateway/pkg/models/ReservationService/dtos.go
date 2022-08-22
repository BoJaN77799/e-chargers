package ReservationService

type ReservationDTO struct {
	Username  string `json:"username"`
	ChargerId uint   `json:"charger_id"`
	VehicleId uint   `json:"vehicle_id"`
	DateFrom  uint64 `json:"date_from"`
	Duration  uint64 `json:"duration"`
}

type CancelReservationDTO struct {
	Username  string `json:"username"`
	ChargerId uint   `json:"charger_id"`
	VehicleId uint   `json:"vehicle_id"`
}
