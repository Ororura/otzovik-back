package model

import "time"

type Review struct {
	ID         int       `json:"id"`
	Text       string    `json:"text"`
	DateCreate time.Time `gorm:"default:NOW()"`
	DateUpdate time.Time `gorm:"default:NOW()"`
	AuthorID   int       `json:"author_id"`
}
