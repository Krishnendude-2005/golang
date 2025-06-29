package task

import (
	"SQLTaskmanager_3layer/models"
	"database/sql"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) *store {
	return &store{db: db}
}

// Create Task.
func (s *store) Create(task models.Task, userID int) (models.Task, error) {
	query := "INSERT INTO tasks (Description, Status, UserID) VALUES (?, ?, ?)"
	res, err := s.db.Exec(query, task.Description, task.Status, userID)
	if err != nil {
		return models.Task{}, err
	}
	id, _ := res.LastInsertId()
	task.ID = int(id)
	task.UserID = userID
	return task, nil
}

// GetById - get tasks by user ID.
func (s *store) GetById(userID int) ([]models.Task, error) {
	query := "SELECT ID, Description, Status, UserID FROM tasks WHERE UserID = ?"
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Status, &task.UserID); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

// DeleteTaskById - deletes by task ID
func (s *store) DeleteTaskById(id int) (int, error) {
	query := "DELETE FROM tasks WHERE ID = ?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update task by task ID.
func (s *store) Update(task models.Task, taskID int) (models.Task, error) {
	query := "UPDATE tasks SET Description = ?, Status = ?, UserID = ? WHERE ID = ?"
	_, err := s.db.Exec(query, task.Description, task.Status, task.UserID, taskID)
	if err != nil {
		return models.Task{}, err
	}
	task.ID = taskID
	return task, nil
}

// GetAll tasks.
func (s *store) GetAll() ([]models.Task, error) {
	query := "SELECT ID, Description, Status, UserID FROM tasks"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Status, &task.UserID); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
