package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	PostID  string `gorm:"not null"`
	UserID  string `gorm:"not null"`
	Content string `gorm:"not null"`
	gorm.Model
}
