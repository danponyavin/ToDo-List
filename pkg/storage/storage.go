package storage

import "ToDo-List/pkg/models"

type TodoStorage interface {
	CreateTask(inp models.CreateTaskReq) (models.Task, error)
}

func (p *PostgresStorage) CreateTask(inp models.CreateTaskReq) (models.Task, error) {
	var task models.Task
	err := p.db.QueryRow("INSERT INTO todo_items (title, description, due_date) VALUES ($1, $2, $3) "+
		"RETURNING id, title, description, due_date, created_at, updated_at", inp.Title, inp.Description, inp.DueDate).
		Scan(&task.Id, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}
