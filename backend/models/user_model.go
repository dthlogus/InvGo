package models

type User struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"fullname" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Perfil   Perfil `json:"perfil"`
}
