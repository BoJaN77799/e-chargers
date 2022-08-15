package models

type UserDTO struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddressDTO struct {
	Street     string  `json:"street"`
	City       string  `json:"city"`
	Country    string  `json:"country"`
	PostalCode uint    `json:"postal_code"`
	Longitude  float32 `json:"longitude"`
	Latitude   float32 `json:"latitude"`
}

type ChargerDTO struct {
	Name        string     `json:"name"`
	Address     AddressDTO `json:"address"`
	WorkTime    string     `json:"work_time"`
	Capacity    int        `json:"capacity"`
	Description string     `json:"description"`
	Rating      float32    `json:"rating"`
	Plugs       string     `json:"plugs"`
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
