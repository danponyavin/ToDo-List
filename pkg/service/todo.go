package service

import "ToDo-List/pkg/storage"

type TodoService struct {
	storage storage.TodoStorage
}

func NewToDoService(storage storage.TodoStorage) *TodoService {
	return &TodoService{storage}
}
