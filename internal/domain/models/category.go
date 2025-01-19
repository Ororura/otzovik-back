package models

type Category struct {
	ID   uint   `json:"-" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
}
