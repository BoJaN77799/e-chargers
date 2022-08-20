package UserService

import "fmt"

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

type VehicleType int

const (
	CAR     VehicleType = 0
	BIKE    VehicleType = 1
	SCOOTER VehicleType = 2
)

func (e VehicleType) String() string {
	switch e {
	case CAR:
		return "CAR"
	case BIKE:
		return "BIKE"
	case SCOOTER:
		return "SCOOTER"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type VehicleDTO struct {
	Name        string      `json:"name"`
	VehicleType VehicleType `json:"vehicle_type"`
}
