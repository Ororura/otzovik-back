package model

import "time"

type Review struct {
	ID         int       `json:"id"`
	Text       string    `json:"text"`
	DateCreate time.Time `json:"date_create"`
	DateUpdate time.Time `json:"date_update"`
	AuthorID   int       `json:"author_id"`
}
