package dto

type InputLoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type OutputToken struct {
	Token string `json:"token"`
}
