package entities

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"strings"
	"user_service/pkg/entities/base"
)

type UserRole int

const (
	Administrator  UserRole = 0
	RegisteredUser UserRole = 1
)

func (e UserRole) String() string {
	switch e {
	case Administrator:
		return "Administrator"
	case RegisteredUser:
		return "RegisteredUser"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type User struct {
	Id          uuid.UUID `json:"id" gorm:"primary_key;type:uuid"`
	Username    string    `json:"username" gorm:"unique;not-null"`
	Email       string    `json:"email"  gorm:"unique;not-null"`
	Role        UserRole  `json:"userRole" gorm:"not null"`
	Firstname   string    `json:"firstname" gorm:"not null"`
	Lastname    string    `json:"lastname" gorm:"not null"`
	Password    string    `json:"password" gorm:"not null"`
	Strikes     uint      `json:"strikes" gorm:"not-null"`
	Banned      bool      `json:"banned" gorm:"not-null"`
	BannedAt    uint64    `json:"banned_at"`
	BannedUntil uint64    `json:"banned_until"`
	Vehicles    []Vehicle
	base.Base
}

type Claims struct {
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Username string    `json:"username"`
	Id       uuid.UUID `json:"id"`
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

func StrToVehicleType(str string) VehicleType {
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
	Id          uuid.UUID   `json:"id" gorm:"primary_key;type:uuid"`
	Name        string      `json:"name" gorm:"unique;not null"`
	VehicleType VehicleType `json:"vehicle_type" gorm:"not null"`
	UserID      uuid.UUID   `json:"user_id" gorm:"type:uuid"`
	base.Base
}
