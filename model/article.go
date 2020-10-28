package model

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	ID        uint   `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
