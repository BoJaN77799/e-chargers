package models

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	Street     string  `json:"street" gorm:"not-null"`
	City       string  `json:"city" gorm:"not-null"`
	Country    string  `json:"country" gorm:"not-null"`
	PostalCode uint    `json:"postal_code" gorm:"not-null"`
	Longitude  float32 `json:"longitude" gorm:"not-null"`
	Latitude   float32 `json:"latitude" gorm:"not-null"`
}

type Plug struct {
	gorm.Model
	PricePerHour           string `json:"price_per_hour" gorm:"not-null"`
	Type                   string `json:"type" gorm:"not-null"`
	ChargingSpeedPerMinute string `json:"charging_speed" gorm:"not-null"`
	//Image string   `json:"image" gorm:"not null"`
}

type Charger struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not-null"`
	Address     Address `json:"address" gorm:"foreignKey:Id"`
	WorkTime    string  `json:"work_time"  gorm:"not-null"`
	Capacity    int     `json:"capacity" gorm:"not-null"`
	Description string  `json:"description" gorm:"not null"`
	Rating      float32 `json:"rating"`
	Plugs       []Plug  `json:"plugs" gorm:"many2many:charger_plugs;"`
	//Image string   `json:"image" gorm:"not null"`
}

func (address *Address) ToDTO() AddressDTO {
	return AddressDTO{Street: address.Street, City: address.City, Country: address.Country, PostalCode: address.PostalCode, Latitude: address.Latitude, Longitude: address.Longitude}
}

func (charger *Charger) ToDTO() ChargerDTO {
	return ChargerDTO{Name: charger.Name, Address: charger.Address.ToDTO(), WorkTime: charger.WorkTime, Capacity: charger.Capacity, Description: charger.Description, Rating: charger.Rating, Plugs: plugsToDTO(charger.Plugs)}
}

func (plug *Plug) ToDTO() PlugDTO {
	return PlugDTO{PricePerHour: plug.PricePerHour, ChargingSpeedPerMinute: plug.ChargingSpeedPerMinute, Type: plug.Type}
}

func plugsToDTO(plugs []Plug) []PlugDTO {
	var plugsDTO []PlugDTO
	for _, plug := range plugs {
		plugsDTO = append(plugsDTO, plug.ToDTO())
	}
	return plugsDTO
}
