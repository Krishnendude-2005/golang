package task

import (
	"SQLTaskmanager_3layer/models"
	"database/sql"
)

type Store interface {
	Create(task models.Task, userID int) (models.Task, error)
	GetById(userID int) ([]models.Task, error)
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

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

func (s *store) GetById(userID int) ([]models.Task, error) {
	query := "SELECT ID, Description, Status, UserID FROM tasks WHERE UserID = ?"
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
