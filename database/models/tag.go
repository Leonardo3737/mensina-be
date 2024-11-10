package models

type Tag struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
}
