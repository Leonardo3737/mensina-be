package services

import "gorm.io/gorm"

func Paginate(page, perPage int, db *gorm.DB) *gorm.DB {
	if page < 1 {
		page = 1
	}
	if perPage <= 0 || perPage > 100 {
		perPage = 10
	}
	offset := (page - 1) * perPage
	return db.Offset(offset).Limit(perPage)
}
