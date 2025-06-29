package user

import (
	"SQLTaskmanager_3layer/models"
	"database/sql"
	"errors"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) Create(user models.User) (models.User, error) {
	var existingID int
	queryTwo := "SELECT ID FROM users WHERE ID = ?"
	err := s.db.QueryRow(queryTwo, user.ID).Scan(&existingID)

	if err == nil { // Means - there is a user in DB with that ID so we got valid Output. That means user already exists.
		return models.User{}, errors.New("user already exists")
	}

	query := "INSERT INTO users (ID, TaskName) VALUES (?, ?)"
	_, errSms := s.db.Exec(query, user.ID, user.TaskName)
	if errSms != nil {
		return models.User{}, err
	}
	//id, _ := res.LastInsertId()
	//user.ID = int(id)
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

//type mockStore struct {
//	// empty
//}
//
////func (m *mockStore) GetById(id int) (models.User, error) {
////	if id == 1 {
////		return models.User{
////			ID:       1,
////			TaskName: "New Task 1",
////		}, nil
////	}
////
////	if id == 2 {
////		return models.User{
////			ID:       2,
////			TaskName: "New Task 2",
////		}, nil
////	}
////	idStr := strconv.Itoa(id)
////	if idStr == "abc" {
////		return models.User{}, errors.New("id cannot be string")
////	}
////	return models.User{}, errors.New("failed")
////}
