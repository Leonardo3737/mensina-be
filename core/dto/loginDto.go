package dto

type InputLoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type OutputToken struct {
	Token string `json:"token"`
}
