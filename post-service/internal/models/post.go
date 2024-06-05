package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Title     string `gorm:"type:varchar(100);not null"`
	Content   string `gorm:"type:text;not null"`
	AuthorID  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
