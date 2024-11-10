package models

type Quiz struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	TagID uint   `json:"tagId"`
	Tag   Tag    `json:"tag,omitempty" gorm:"foreignKey:TagID"`
}
