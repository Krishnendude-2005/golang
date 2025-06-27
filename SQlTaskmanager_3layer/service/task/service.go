package task

import (
	"SQLTaskmanager_3layer/models"
	storePkg "SQLTaskmanager_3layer/store/task"
)

type Service interface {
	Create(task models.Task, userID int) (models.Task, error)
	GetById(userID int) ([]models.Task, error)
}

type service struct {
	store storePkg.Store
}

func New(store storePkg.Store) Service {
	return &service{store: store}
}

func (s *service) Create(task models.Task, userID int) (models.Task, error) {
	if err := task.Validate(); err != nil {
		return models.Task{}, err
	}
	return s.store.Create(task, userID)
}

func (s *service) GetById(userID int) ([]models.Task, error) {
	return s.store.GetById(userID)
}
