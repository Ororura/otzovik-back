package models

type Review struct {
	ID         uint     `json:"-" gorm:"primaryKey"`
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	Rating     int      `json:"rating" binding:"required,min=1,max=10"`
	ImagePath  string   `json:"image_path"`
	UserID     uint     `json:"user_id" binding:"required"`
	CategoryID uint     `json:"category_id" binding:"required"`
	User       User     `json:"user" gorm:"foreignKey:UserID"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
}
