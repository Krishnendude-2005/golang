package user

import (
	"SQLTaskmanager_3layer/models"
	"gofr.dev/pkg/gofr"
)

type Store interface {
	Create(c *gofr.Context, user models.User) (models.User, error)
	GetById(c *gofr.Context, id int) (models.User, error)
}

type store struct{}

func New() Store {
	return &store{}
}

// Create - Uses userID from input.
func (s *store) Create(c *gofr.Context, user models.User) (models.User, error) {
	query := "INSERT INTO users (ID, TaskName) VALUES (?, ?)"
	_, err := c.SQL.ExecContext(c, query, user.ID, user.TaskName)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *store) GetById(c *gofr.Context, id int) (models.User, error) {
	query := "SELECT ID, TaskName FROM users WHERE ID = ?"
	row := c.SQL.QueryRowContext(c, query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.TaskName)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
