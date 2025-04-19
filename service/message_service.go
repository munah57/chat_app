package service

import "real-chat/repository"

type MessageService struct {
	Repo repository.MessageRepository
}
