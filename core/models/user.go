package models

import "time"

// @Description User object
// @Schema
type User struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserName  string     `json:"username" validate:"required,min=3"`
	Password  string     `json:"password" validate:"required,min=6"`
	Name      string     `json:"name" validate:"required"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
