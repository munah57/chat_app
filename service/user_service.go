package service

import "real-chat/repository"

type UserService struct {
	Repo repository.UserRepository
}
