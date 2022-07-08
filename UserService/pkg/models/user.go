package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
	Id        int      `json:"id" gorm:"primaryKey"`
	Username  string   `json:"username" gorm:"unique;not-null"`
	Password  string   `json:"password" gorm:"not null"`
	Email     string   `json:"email"  gorm:"unique;not-null"`
	Firstname string   `json:"firstname" gorm:"not null"`
	Lastname  string   `json:"lastname" gorm:"not null"`
	Role      UserRole `json:"userRole" gorm:"not null"`
}

type Claims struct {
	Email    string `json:"email"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Id       int    `json:"id"`
	jwt.StandardClaims
}
