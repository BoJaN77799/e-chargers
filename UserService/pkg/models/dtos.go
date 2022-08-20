package models

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserTokenState struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expiredAt"`
}

type VehicleDTO struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	VehicleType string `json:"vehicle_type"`
	Username    string `json:"username"`
}

type UserReservationDTO struct {
	Username string `json:"username" gorm:"unique;not-null"`
	Vehicles []VehicleDTO
}
