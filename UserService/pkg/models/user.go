package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"strings"
)

type UserRole int

const (
	Administrator    UserRole = 0
	UnregisteredUser UserRole = 1
	RegisteredUser   UserRole = 2
)

func (e UserRole) String() string {
	switch e {
	case Administrator:
		return "Administrator"
	case UnregisteredUser:
		return "UnregisteredUser"
	case RegisteredUser:
		return "RegisteredUser"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type User struct {
	gorm.Model
	Username  string   `json:"username" gorm:"unique;not-null"`
	Password  string   `json:"password" gorm:"not null"`
	Email     string   `json:"email"  gorm:"unique;not-null"`
	Firstname string   `json:"firstname" gorm:"not null"`
	Lastname  string   `json:"lastname" gorm:"not null"`
	Role      UserRole `json:"userRole" gorm:"not null"`
	Vehicles  []Vehicle
}

type Claims struct {
	Email    string `json:"email"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Id       uint   `json:"id"`
	jwt.StandardClaims
}

type VehicleType string

const (
	CAR     VehicleType = "CAR"
	BIKE    VehicleType = "BIKE"
	SCOOTER VehicleType = "SCOOTER"
)

var (
	capabilitiesMap = map[string]VehicleType{
		"car":     CAR,
		"bike":    BIKE,
		"scooter": SCOOTER,
	}
)

func ParseString(str string) VehicleType {
	c, _ := capabilitiesMap[strings.ToLower(str)]
	return c
}

func (e VehicleType) String() string {
	switch e {
	case CAR:
		return "CAR"
	case BIKE:
		return "BIKE"
	case SCOOTER:
		return "SCOOTER"
	default:
		return fmt.Sprintf("%s", string(e))
	}
}

type Vehicle struct {
	gorm.Model
	Name        string      `json:"name" gorm:"unique;not null"`
	VehicleType VehicleType `json:"vehicle_type" gorm:"not null"`
	UserID      uint
}

func (vehicle *Vehicle) ToDTO() VehicleDTO {
	return VehicleDTO{
		Name:        vehicle.Name,
		VehicleType: vehicle.VehicleType.String(),
	}
}
