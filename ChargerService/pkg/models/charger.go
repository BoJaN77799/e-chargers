package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type Address struct {
	gorm.Model
	Street     string  `json:"street" gorm:"not-null"`
	City       string  `json:"city" gorm:"not-null"`
	Country    string  `json:"country" gorm:"not-null"`
	PostalCode uint    `json:"postal_code" gorm:"not-null"`
	Longitude  float32 `json:"longitude" gorm:"not-null"`
	Latitude   float32 `json:"latitude" gorm:"not-null"`
	ChargerId  int
}

type Charger struct {
	gorm.Model
	Name                   string  `json:"name" gorm:"not-null"`
	Address                Address `json:"address" gorm:"foreignKey:Id"`
	WorkTimeFrom           int     `json:"work_time_from"  gorm:"not-null"`
	WorkTimeTo             int     `json:"work_time_to"  gorm:"not-null"`
	Capacity               uint    `json:"capacity" gorm:"not-null"`
	Description            string  `json:"description" gorm:"not null"`
	Rating                 float32 `json:"rating"`
	Plugs                  string  `json:"plugs"  gorm:"not-null"`
	PricePerHour           int     `json:"price_per_hour" gorm:"not-null"`
	ChargingSpeedPerMinute int     `json:"charging_speed" gorm:"not-null"`
	//Image string   `json:"image" gorm:"not null"`
}

func (address *Address) ToDTO() AddressDTO {
	return AddressDTO{
		Street:     address.Street,
		City:       address.City,
		Country:    address.Country,
		PostalCode: address.PostalCode,
		Latitude:   address.Latitude,
		Longitude:  address.Longitude,
	}
}

func (charger *Charger) ToDTO() ChargerDTO {
	return ChargerDTO{
		Name:                   charger.Name,
		Address:                charger.Address.ToDTO(),
		WorkTime:               fmt.Sprintf("%dh - %dh", charger.WorkTimeFrom, charger.WorkTimeTo),
		Capacity:               charger.Capacity,
		Description:            charger.Description,
		Rating:                 charger.Rating,
		Plugs:                  strings.Split(charger.Plugs, ","),
		PricePerHour:           fmt.Sprintf("%d €/Hour", charger.PricePerHour),
		ChargingSpeedPerMinute: fmt.Sprintf("%d kW/Min", charger.ChargingSpeedPerMinute),
	}
}

func (charger *Charger) ToReportDTO() ChargerReportDTO {
	return ChargerReportDTO{
		Name:                   charger.Name,
		Capacity:               charger.Capacity,
		Rating:                 charger.Rating,
		PricePerHour:           charger.PricePerHour,
		ChargingSpeedPerMinute: charger.ChargingSpeedPerMinute,
	}
}
