package models

type Review struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Rating       int    `json:"rating"`
	ImagePath    string `json:"image_path,omitempty"`    // omitempty — если поле пустое, оно не попадёт в JSON
	UserID       uint   `json:"user_id"`                 // Внешний ключ на User
	CategoryID   uint   `json:"category_id"`             // Внешний ключ на Category
	UserName     string `json:"user_name,omitempty"`     // Имя пользователя (из JOIN)
	CategoryName string `json:"category_name,omitempty"` // Имя категории (из JOIN)
	Review       *User
	Category     *Category
}
