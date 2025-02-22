package repositories

import (
	"database/sql"
	"fmt"
	"otzovik-back/internal/domain"
	"otzovik-back/internal/domain/models"
)

type ChatRepo struct {
	db *sql.DB
}

func NewChatRepo(db *sql.DB) domain.ChatRepository {
	return &ChatRepo{db: db}
}

func (c *ChatRepo) SaveMessage(message *models.Messages) error {
	query:= "INSERT INTO messages(title, user_id) VALUES($1, $2)"
	err := c.db.QueryRow(query, message.Title, message.ID)
	if err != nil {
		return fmt.Errorf("ошибка при сохранении сообщения: %v", err)
	}

	return nil
}

func (c *ChatRepo) GetMessage(id *int) ([]*models.Messages, error) {
	var messages []*models.Messages
	query := "SELECT * FROM messages WHERE id = $1"

	rows, err := c.db.Query(query, id)

	if err != nil {
		return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
	}

	for rows.Next() {
		msg := new(models.Messages)
		if err := rows.Scan(&msg.ID, &msg.Title, &msg.CreatedAt, &msg.UpdatedAt, &msg.UserID); err != nil {
			return nil, fmt.Errorf("ошибка сканирования: %v", err)
		}

		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения строк: %v", err)
	}

	return messages, nil
}
