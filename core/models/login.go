package models

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRes struct {
	Token string `json:"token"`
}
