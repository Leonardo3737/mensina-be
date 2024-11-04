package dto

type CreateUserDto struct {
	UserName string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}

type UpdateUserDto struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
