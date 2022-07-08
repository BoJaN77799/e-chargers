package models

type UserRole int

const (
	Administrator    UserRole = 0
	UnregisteredUser UserRole = 1
	RegisteredUser   UserRole = 2
)

type User struct {
	Id        int      `json:"id" gorm:"primaryKey"`
	Username  string   `json:"username" gorm:"unique;not-null"`
	Password  string   `json:"password" gorm:"not null"`
	Email     string   `json:"email"  gorm:"unique;not-null"`
	Firstname string   `json:"firstname" gorm:"not null"`
	Lastname  string   `json:"lastname" gorm:"not null"`
	Role      UserRole `json:"userRole" gorm:"not null"`
}
