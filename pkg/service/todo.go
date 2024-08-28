package service

import (
	"ToDo-List/pkg/models"
	"ToDo-List/pkg/storage"
	"log"
)

type TodoService struct {
	storage storage.TodoStorage
}

func NewToDoService(storage storage.TodoStorage) *TodoService {
	return &TodoService{storage}
}

func (t *TodoService) CreateTask(inp models.CreateTaskReq) (models.Task, error) {
	task, err := t.storage.CreateTask(inp)
	if err != nil {
		log.Println("CreateTask:", err)
		return models.Task{}, err
	}

	return task, nil
}
