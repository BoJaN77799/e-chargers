package entities

import uuid "github.com/satori/go.uuid"

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserTokenState struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expiredAt"`
}

type VehicleDto struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	VehicleType string    `json:"vehicle_type"`
	Username    string    `json:"username"`
}

type UserReservationDTO struct {
	Username string `json:"username" gorm:"unique;not-null"`
	Vehicles []VehicleDto
}

type UserProfileDTO struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Vehicles    []VehicleDto
	Strikes     uint   `json:"strikes"`
	Banned      bool   `json:"banned"`
	BannedAt    uint64 `json:"banned_at"`
	BannedUntil uint64 `json:"banned_until"`
}

type UserReportDTO struct {
	Id          uuid.UUID `json:"id" `
	Username    string    `json:"username" `
	Email       string    `json:"email"  `
	Role        string    `json:"user_role"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Strikes     uint      `json:"strikes"`
	Banned      bool      `json:"banned"`
	BannedAt    uint64    `json:"banned_at"`
	BannedUntil uint64    `json:"banned_until"`
}
