package models

type User struct {
	Id           string
	Email        string `json:"email"`
	Name         string `json:"name"` // giving the name of the field in the json ouptut
	Age          int64  `json:"age"`
	PasswordHash string `json:"password_hash"`
}

type NewUser struct {
	Name     string `json:"name" validate:"required,max=60"`
	Email    string `json:"email" validate:"required,email,max=60"`
	Age      int64  `json:"age" validate:"required,min=18"`
	Password string `json:"password" validate:"required,max=60"`
}
