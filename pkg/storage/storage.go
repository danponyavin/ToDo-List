package storage

import (
	"ToDo-List/pkg/models"
	"database/sql"
	"errors"
)

var TaskNotFound = errors.New("Задача не найдена.")

type TodoStorage interface {
	Create(inp models.CreateTaskReq) (models.Task, error)
	GetAll() ([]models.Task, error)
	GetById(id int) (models.Task, error)
}

func (p *PostgresStorage) Create(inp models.CreateTaskReq) (models.Task, error) {
	var task models.Task
	err := p.db.QueryRow("INSERT INTO todo_items (title, description, due_date) VALUES ($1, $2, $3) "+
		"RETURNING id, title, description, due_date, created_at, updated_at", inp.Title, inp.Description, inp.DueDate).
		Scan(&task.Id, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (p *PostgresStorage) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := p.db.Select(&tasks, "SELECT * FROM todo_items")
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (p *PostgresStorage) GetById(id int) (models.Task, error) {
	var task models.Task
	err := p.db.Get(&task, "SELECT * FROM todo_items WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Task{}, TaskNotFound
		}
		return models.Task{}, err
	}

	return task, nil
}
