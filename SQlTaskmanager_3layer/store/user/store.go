package user

import (
	"SQLTaskmanager_3layer/models"
	"database/sql"
)

type Store interface {
	Create(user models.User) (models.User, error)
	GetById(id int) (models.User, error)
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

func (s *store) Create(user models.User) (models.User, error) {
	query := "INSERT INTO users (TaskName) VALUES (?)"
	res, err := s.db.Exec(query, user.TaskName)
	if err != nil {
		return models.User{}, err
	}
	id, _ := res.LastInsertId()
	user.ID = int(id)
	return user, nil
}

func (s *store) GetById(id int) (models.User, error) {
	query := "SELECT ID, TaskName FROM users WHERE ID = ?"
	row := s.db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.TaskName)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
