package models

type UserRole int

const (
	Administrator UserRole = iota
	UnregisteredUser
	RegisteredUser
)

type User struct {
	Id       int      `json:"id" gorm:"primaryKey"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Role     UserRole `json:"userRole"`
}
