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
	task, err := t.storage.Create(inp)
	if err != nil {
		log.Println("Service. CreateTask:", err)
		return models.Task{}, err
	}

	return task, nil
}

func (t *TodoService) GetAllTasks() ([]models.Task, error) {
	tasks, err := t.storage.GetAll()
	if err != nil {
		log.Println("Service. GetAll:", err)
		return []models.Task{}, err
	}

	return tasks, nil
}

func (t *TodoService) GetTask(id int) (models.Task, error) {
	task, err := t.storage.GetById(id)
	if err != nil {
		log.Println("Service. GetTask:", err)
		return models.Task{}, err
	}

	return task, nil
}

func (t *TodoService) UpdateTask(id int, inp models.CreateTaskReq) (models.Task, error) {
	task, err := t.storage.UpdateTask(id, inp)
	if err != nil {
		log.Println("Service. UpdateTask:", err)
		return models.Task{}, err
	}

	return task, nil
}

func (t *TodoService) DeleteTask(id int) error {
	err := t.storage.DeleteTask(id)
	if err != nil {
		log.Println("Service. DeleteTask:", err)
		return err
	}

	return nil
}
