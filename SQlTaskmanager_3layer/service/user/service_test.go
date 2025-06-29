package user

import (
	"SQLTaskmanager_3layer/models"
	"errors"
	"testing"
)

type mockStore struct {
	//empty
}

func (m *mockStore) Create(user models.User) (models.User, error) {
	var userMade models.User
	userMade.ID = user.ID
	userMade.TaskName = user.TaskName
	return userMade, nil
}
func Test_Create(t *testing.T) {
	svc := New(&mockStore{})
	newUser := models.User{
		ID:       1,
		TaskName: "task1",
	}
	respUser, err := svc.Create(newUser)
	if err != nil {
		t.Error(err)
	}
	if respUser.ID != 1 {
		t.Error("id is not 1")
	}
	if respUser.TaskName != "task1" {
		t.Error("task name is not task1")
	}
}
func (m *mockStore) GetById(id int) (models.User, error) {
	var gotUser models.User
	var users = []models.User{
		{
			ID:       1,
			TaskName: "New task 1",
		},
		{
			ID:       2,
			TaskName: "New task 2",
		},
	}

	for i := range users {
		if users[i].ID == id {
			gotUser = users[i]
			return gotUser, nil
		}
	}
	return models.User{}, errors.New("user not found")
}
func Test_GetById(t *testing.T) {
	svc := New(&mockStore{})

	respUser, err := svc.GetById(1)
	if err != nil {
		t.Error(err)
	}

	if respUser.ID != 1 {
		t.Error("id is not 1")
	}
}
