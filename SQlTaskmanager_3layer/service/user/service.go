package user

import (
	"SQLTaskmanager_3layer/models"
	storePkg "SQLTaskmanager_3layer/store/user"
)

type Service interface {
	Create(user models.User) (models.User, error)
	GetById(id int) (models.User, error)
}

type service struct {
	store storePkg.Store
}

func New(store storePkg.Store) Service {
	return &service{store: store}
}

func (s *service) Create(user models.User) (models.User, error) {
	if err := user.Validate(); err != nil {
		return models.User{}, err
	}
	return s.store.Create(user)
}

func (s *service) GetById(id int) (models.User, error) {
	return s.store.GetById(id)
}
