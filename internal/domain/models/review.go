package models

type Review struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    Title     string `json:"title"`
    Content   string `json:"content"`
    Rating    int    `json:"rating"`
    ImagePath string `json:"image_path"`
    UserID    uint   `json:"user_id"`
    User      User   `json:"user" gorm:"foreignKey:UserID"`
} 