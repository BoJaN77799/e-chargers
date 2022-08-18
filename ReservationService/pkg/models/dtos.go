package models

type AddressDTO struct {
	Street     string  `json:"street"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	PostalCode uint    `json:"postal_code"`
	Longitude  float32 `json:"longitude"`
	Latitude   float32 `json:"latitude"`
}

type ReservationDTO struct {
	Username  string `json:"username"`
	ChargerId uint   `json:"charger_id"`
	VehicleId uint   `json:"vehicle_id"`
	DateFrom  uint64 `json:"date_from"`
	Duration  uint64 `json:"duration"`
}
