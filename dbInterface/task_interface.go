package dbInterface

import (
	"database/sql"
	"fmt"
)

type TaskService interface {
	CreateTask(*Task) (int, error)
	GetTask(id string) (*Task, error)
	UpdateTask(*Task) error
	DeleteTask(id string) error
	//GetTasksByUser(userID string) ([]Task, error)
}

func (r *DbRepo) CreateTask(task *Task) (int, error) {
	query := `INSERT INTO tasks (name, description, user_id, status_id) VALUES (?, ?, ?, ?)`

	result, err := r.SqlConnection.Exec(query, task.TaskName, task.TaskDesc, task.UserID, task.StatusID)
	if err != nil {
		return 0, fmt.Errorf("failed to create task: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("could not retrieve task ID: %v", err)
	}

	return int(id), nil
}

func (r *DbRepo) GetTask(id string) (*Task, error) {
	query := `SELECT id, name, description, user_id, status_id FROM tasks WHERE id = ?`

	row := r.SqlConnection.QueryRow(query, id)

	var task Task
	err := row.Scan(&task.TaskID, &task.TaskName, &task.TaskDesc, &task.UserID, &task.StatusID)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task with ID %s not found", id)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get task: %v", err)
	}

	return &task, nil
}

func (r *DbRepo) UpdateTask(task *Task) error {
	query := `UPDATE tasks SET name = ?, description = ?, user_id = ?, status_id = ? WHERE id = ?`

	result, err := r.SqlConnection.Exec(query, task.TaskName, task.TaskDesc, task.UserID, task.StatusID, task.TaskID)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to verify update result: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no task found with ID %d", task.TaskID)
	}

	return nil
}

func (r *DbRepo) DeleteTask(id string) error {
	query := `DELETE FROM tasksWHERE id = ?`

	result, err := r.SqlConnection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not confirm deletion: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no task found with ID %s", id)
	}

	return nil
}
