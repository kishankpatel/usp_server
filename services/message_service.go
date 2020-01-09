package services

import (
	"fmt"

	"github.com/kishankpatel/usp_server/models"
)

// MessageService interface declaration
type MessageService interface {
	ReceiveMessage(key string, msg string) (*models.Message, error)
	EncryptMessage(key string, msg string) (*models.Message, error)
}

type messageService struct {
	Message map[string]*models.Message
}

// NewMessageService method declaration
func NewMessageService() MessageService {
	return &messageService{
		Message: make(map[string]*models.Message),
	}
}

func (ms messageService) ReceiveMessage(key string, msg string) (*models.Message, error) {
	result, err := models.Decrypt(key, msg)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return result, nil
}

func (ms messageService) EncryptMessage(key string, msg string) (*models.Message, error) {
	result, err := models.Encrypt(key, msg)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return result, nil
}
