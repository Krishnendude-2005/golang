package task

import (
	"SQLTaskmanager_3layer/models"
	"gofr.dev/pkg/gofr"
)

type Service struct {
	store Store
}

func New(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) Create(c *gofr.Context, task models.Task, userID int) (models.Task, error) {
	if err := task.Validate(); err != nil {
		return models.Task{}, err
	}
	return s.store.Create(c, task, userID)
}

func (s *Service) GetById(c *gofr.Context, userID int) ([]models.Task, error) {
	return s.store.GetById(c, userID)
}

func (s *Service) DeleteTaskById(c *gofr.Context, id int) (int, error) {
	return s.store.DeleteTaskById(c, id)
}

func (s *Service) Update(c *gofr.Context, task models.Task, taskID int) (models.Task, error) {
	if err := task.Validate(); err != nil {
		return models.Task{}, err
	}
	return s.store.Update(c, task, taskID)
}

func (s *Service) GetAll(c *gofr.Context) ([]models.Task, error) {
	return s.store.GetAll(c)
}
