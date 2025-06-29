package user

import (
	"SQLTaskmanager_3layer/models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockService struct {
	//empty
}

func (m *mockService) Create(user models.User) (models.User, error) {
	return models.User{
		ID:       user.ID,
		TaskName: user.TaskName,
	}, nil
}
func TestCreate(t *testing.T) {
	handler := New(&mockService{})
	task := models.User{
		ID:       101,
		TaskName: "New user",
	}
	taskJson, _ := json.Marshal(task)
	req := httptest.NewRequest(http.MethodPost, "/user/add", bytes.NewBuffer(taskJson))
	req.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()

	var respUser models.User

	handler.Create(writer, req)
	result := writer.Result()
	respBody, _ := io.ReadAll(result.Body)
	err := json.Unmarshal(respBody, &respUser)
	if err != nil {
		t.Error("Some error in unmarshalling")
	}
	if writer.Code != http.StatusOK {
		t.Error("Wrong response code")
	}
	if respUser.ID != task.ID {
		t.Error("Same ID expected")
	}
	if respUser.TaskName != task.TaskName {
		t.Error("Same TaskName expected")
	}
}
func (m *mockService) GetById(id int) (models.User, error) {
	if id == 101 {
		return models.User{
			ID:       101,
			TaskName: "New user",
		}, nil
	}
	return models.User{}, errors.New("user not found")
}
func TestGetByID(t *testing.T) {
	handler := New(&mockService{})

	req := httptest.NewRequest(http.MethodGet, "/user/find?id=101", nil)
	w := httptest.NewRecorder()

	handler.GetByID(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	var user models.User
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if user.ID != 101 || user.TaskName != "New user" {
		t.Errorf("Unexpected user data: %+v", user)
	}
}
