package models

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
