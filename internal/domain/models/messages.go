package models

import "time"

type Messages struct {
	ID        uint
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
}
