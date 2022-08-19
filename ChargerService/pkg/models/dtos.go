package models

type AddressDTO struct {
	Street     string  `json:"street"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	PostalCode uint    `json:"postal_code"`
	Longitude  float32 `json:"longitude"`
	Latitude   float32 `json:"latitude"`
}

type ChargerDTO struct {
	Name                   string     `json:"name"`
	Address                AddressDTO `json:"address"`
	WorkTime               string     `json:"work_time"`
	Capacity               uint       `json:"capacity"`
	Description            string     `json:"description"`
	Rating                 float32    `json:"rating"`
	Plugs                  []string   `json:"plugs"`
	ChargingSpeedPerMinute string     `json:"charging_speed"`
	PricePerHour           string     `json:"price_per_hour"`
	//Image string   `json:"image" gorm:"not null"`
}

type SearchDTO struct {
	Name              string `json:"name"`
	WorkTimeFrom      int    `json:"workTimeFrom"`
	WorkTimeTo        int    `json:"workTimeTo"`
	Capacity          int    `json:"capacity"`
	PricePerHourFrom  int    `json:"pricePerHourFrom"`
	PricePerHourTo    int    `json:"pricePerHourTo"`
	Type              string `json:"type"`
	ChargingSpeedFrom int    `json:"chargingSpeedFrom"`
	ChargingSpeedTo   int    `json:"chargingSpeedTo"`
}

type ChargerReportDTO struct {
	Name                   string  `json:"name"`
	Capacity               uint    `json:"capacity"`
	Rating                 float32 `json:"rating"`
	ChargingSpeedPerMinute int     `json:"charging_speed"`
	PricePerHour           int     `json:"price_per_hour"`
}
