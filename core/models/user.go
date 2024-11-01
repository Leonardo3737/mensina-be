package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"  validate:"required"`
	UserName string `json:"username"  validate:"required,min=3"`
	Password string `json:"password"  validate:"required,min=6"`
}
