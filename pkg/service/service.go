package service

import (
	"ToDo-List/pkg/models"
	"ToDo-List/pkg/storage"
)

type IToDoService interface {
	CreateTask(inp models.CreateTaskReq) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTask(id int) (models.Task, error)
	UpdateTask(id int, inp models.CreateTaskReq) (models.Task, error)
	DeleteTask(id int) error
}

type Services struct {
	IToDoService
}

func NewServices(toDoStorage storage.TodoStorage) *Services {
	return &Services{
		IToDoService: NewToDoService(toDoStorage),
	}
}
