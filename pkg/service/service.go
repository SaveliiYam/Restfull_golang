package service

import "example.com/m/v2/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
