package entities

import uuid "github.com/satori/go.uuid"

type LoginDTO struct {
	Email    string `json:"email"`
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
}

type UserReservationDTO struct {
	Id       string `json:"id"`
	Vehicles []VehicleDto
}

type UserProfileDTO struct {
	Id          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Strikes     uint      `json:"strikes"`
	Banned      bool      `json:"banned"`
	BannedAt    uint64    `json:"banned_at"`
	BannedUntil uint64    `json:"banned_until"`
	Vehicles    []VehicleDto
}

type UserReportDTO struct {
	Id          uuid.UUID `json:"id" `
	Email       string    `json:"email"  `
	Role        string    `json:"user_role"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Strikes     uint      `json:"strikes"`
	Banned      bool      `json:"banned"`
	BannedAt    uint64    `json:"banned_at"`
	BannedUntil uint64    `json:"banned_until"`
}
