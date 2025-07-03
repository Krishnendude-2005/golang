package user

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

func (s *Service) Create(c *gofr.Context, user models.User) (models.User, error) {
	if err := user.Validate(); err != nil {
		return models.User{}, err
	}
	return s.store.Create(c, user)
}

func (s *Service) GetById(c *gofr.Context, id int) (models.User, error) {
	return s.store.GetById(c, id)
}
