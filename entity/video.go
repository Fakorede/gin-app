package entity

import "time"

type Video struct {
	ID	uint64	`json:"id" gorm:"primary_key;auto_increment"`
	Title	string	`json:"title" binding:"min=2,max=100" validate:"custom" gorm:"type:varchar(100)"`
	Description	string	`json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL	string	`json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author Person	`json:"author" binding:"required" gorm:"foreign_key:PersonID"`
	PersonID uint64 `json:"-"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}