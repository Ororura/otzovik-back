package models

type User struct {
	ID       uint   `json:"-" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
