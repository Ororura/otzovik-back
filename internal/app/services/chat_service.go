package services

import (
	"otzovik-back/internal/domain"
	"otzovik-back/internal/domain/models"
)

type ChatService struct {
	repo domain.ChatRepository
}

func NewChatService(repo domain.ChatRepository) domain.ChatService { return &ChatService{repo: repo} }

func (c *ChatService) SaveMessage(message *models.Messages) error {
	return c.repo.SaveMessage(message)
}

func (c *ChatService) GetMessage(id *int) ([]*models.Messages, error) {
	return c.repo.GetMessage(id)
}
