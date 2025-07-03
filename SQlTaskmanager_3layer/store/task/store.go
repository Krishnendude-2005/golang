package task

import (
	"SQLTaskmanager_3layer/models"
	"gofr.dev/pkg/gofr"
)

type Store struct{}

func New() *Store {
	return &Store{}
}

func (s *Store) Create(c *gofr.Context, task models.Task, userID int) (models.Task, error) {
	query := "INSERT INTO tasks (Description, Status, UserID) VALUES (?, ?, ?)"
	res, err := c.SQL.ExecContext(c, query, task.Description, task.Status, userID)
	if err != nil {
		return models.Task{}, err
	}

	id, _ := res.LastInsertId()
	task.ID = int(id)
	task.UserID = userID

	return task, nil
}

func (s *Store) GetById(c *gofr.Context, userID int) ([]models.Task, error) {
	query := "SELECT ID, Description, Status, UserID FROM tasks WHERE UserID = ?"
	rows, err := c.SQL.QueryContext(c, query, userID)
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
	return tasks, rows.Err()
}
func (s *Store) DeleteTaskById(c *gofr.Context, id int) (int, error) {
	query := "DELETE FROM tasks WHERE ID = ?"
	res, err := c.SQL.ExecContext(c, query, id)

	if err != nil {
		return 0, err
	}

	affected, _ := res.RowsAffected()

	return int(affected), nil
}
func (s *Store) Update(c *gofr.Context, task models.Task, taskID int) (models.Task, error) {
	query := "UPDATE tasks SET Description = ?, Status = ?, UserID = ? WHERE ID = ?"
	_, err := c.SQL.ExecContext(c, query, task.Description, task.Status, task.UserID, taskID)

	if err != nil {

		return models.Task{}, err
	}

	task.ID = taskID

	return task, nil
}

func (s *Store) GetAll(c *gofr.Context) ([]models.Task, error) {
	query := "SELECT ID, Description, Status, UserID FROM tasks"
	rows, err := c.SQL.QueryContext(c, query)

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
	return tasks, rows.Err()
}
