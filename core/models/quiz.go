package models

type Quiz struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Score int    `json:"score"`
	TagID uint   `json:"tagId"`
	Tag   Tag    `json:"tag,omitempty" gorm:"foreignKey:TagID"`
}
