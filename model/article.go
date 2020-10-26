package model

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	ID      uint   `json:"id" gorm:"primaryKey"`
}
