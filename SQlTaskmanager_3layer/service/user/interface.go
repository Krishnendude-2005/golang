package user

import (
	"SQLTaskmanager_3layer/models"
	"gofr.dev/pkg/gofr"
)

type Store interface {
	Create(c *gofr.Context, user models.User) (models.User, error)
	GetById(c *gofr.Context, id int) (models.User, error)
}
