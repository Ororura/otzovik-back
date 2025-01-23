package repositories

import (
	"gorm.io/gorm"
	"otzovik-back/internal/domain"
	"otzovik-back/internal/domain/models"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) domain.ChatRepository {
	return &ChatRepo{db: db}
}

func (c *ChatRepo) SaveMessage(message *models.Messages) error {
	return c.db.Create(message).Error
}

func (c *ChatRepo) GetMessage(id *int) ([]*models.Messages, error) {
	var messages []*models.Messages
	err := c.db.Where("id = ?", *id).Find(&messages).Error
	return messages, err
}
