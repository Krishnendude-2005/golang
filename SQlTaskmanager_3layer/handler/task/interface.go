// handler/task/interface.go

package task

import (
	"SQLTaskmanager_3layer/models"
	"gofr.dev/pkg/gofr"
)

type Service interface {
	Create(c *gofr.Context, task models.Task, userID int) (models.Task, error)
	GetById(c *gofr.Context, userID int) ([]models.Task, error)
	DeleteTaskById(c *gofr.Context, id int) (int, error)
	Update(c *gofr.Context, task models.Task, taskID int) (models.Task, error)
	GetAll(c *gofr.Context) ([]models.Task, error)
}
