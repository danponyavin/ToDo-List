package service

import "ToDo-List/pkg/storage"

type IToDoService interface {
}

type Services struct {
	IToDoService
}

func NewServices(toDoStorage storage.TodoStorage) *Services {
	return &Services{
		IToDoService: NewToDoService(toDoStorage),
	}
}
