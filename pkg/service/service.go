package service

import (
	todo "example.com/m/v2"
	"example.com/m/v2/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}
type TodoList struct {
}

type TodoItem struct {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
