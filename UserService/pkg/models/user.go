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
	Username    string   `json:"username" gorm:"unique;not-null"`
	Password    string   `json:"password" gorm:"not null"`
	Email       string   `json:"email"  gorm:"unique;not-null"`
	Firstname   string   `json:"firstname" gorm:"not null"`
	Lastname    string   `json:"lastname" gorm:"not null"`
	Role        UserRole `json:"userRole" gorm:"not null"`
	Vehicles    []Vehicle
	Strikes     uint   `json:"strikes" gorm:"not-null"`
	Banned      bool   `json:"banned" gorm:"not-null"`
	BannedAt    uint64 `json:"banned_at"`
	BannedUntil uint64 `json:"banned_until"`
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
	UserID      uint        `json:"user_id"`
}

func (vehicle *Vehicle) ToDTO() VehicleDTO {
	return VehicleDTO{
		Id:          vehicle.ID,
		Name:        vehicle.Name,
		VehicleType: vehicle.VehicleType.String(),
	}
}

func (user *User) ToDTO() UserReservationDTO {
	return UserReservationDTO{
		Username: user.Username,
		Vehicles: vehiclesToDto(user.Vehicles),
	}
}

func vehiclesToDto(vehicles []Vehicle) []VehicleDTO {
	var vehiclesDTO []VehicleDTO
	for _, vehicle := range vehicles {
		vehiclesDTO = append(vehiclesDTO, vehicle.ToDTO())
	}
	return vehiclesDTO
}

func (user *User) ToUserProfileDTO() UserProfileDTO {
	return UserProfileDTO{
		Email:     user.Email,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Strikes:   user.Strikes,
		Vehicles:  vehiclesToDto(user.Vehicles),
	}
}
