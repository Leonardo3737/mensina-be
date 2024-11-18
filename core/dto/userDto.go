package dto

type CreateUserDto struct {
	UserName string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=50"`
	Name     string `json:"name" binding:"required,min=2,max=50"`
}

type UpdateUserDto struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type OutputUserDto struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Name     string `json:"name"`
}
