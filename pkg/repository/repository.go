package repository

type Authorization interface {
}
type TodoList struct {
}

type TodoItem struct {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}
