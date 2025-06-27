package task

import (
	"SQLTaskmanager_3layer/models"
)

type Store interface {
	Create(task models.Task, userID int) (models.Task, error)
	GetById(userID int) ([]models.Task, error)
	DeleteTaskById(id int) error
	Update(task models.Task, taskID int) (models.Task, error)
	GetAll() ([]models.Task, error)
}

type service struct {
	store Store
}

func New(store Store) *service {
	return &service{store: store}
}

func (s *service) Create(task models.Task, userID int) (models.Task, error) {
	err := task.Validate()
	if err != nil {
		return models.Task{}, err
	}
	return s.store.Create(task, userID)
}

func (s *service) GetById(userID int) ([]models.Task, error) {
	return s.store.GetById(userID)
}

func (s *service) Delete(id int) error {
	return s.store.DeleteTaskById(id)
}

func (s *service) Update(task models.Task, taskID int) (models.Task, error) {
	err := task.Validate()
	if err != nil {
		return models.Task{}, err
	}
	return s.store.Update(task, taskID)
}

func (s *service) GetAll() ([]models.Task, error) {
	return s.store.GetAll()
}
