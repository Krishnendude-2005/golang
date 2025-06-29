package task

import (
	"SQLTaskmanager_3layer/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockService struct {
	//empty
}

func (s *mockService) GetById(userID int) ([]models.Task, error) {
	return []models.Task{
		{ID: 1, Description: "Task 1", Status: false, UserID: userID},
		{ID: 2, Description: "Task 2", Status: true, UserID: userID},
	}, nil
}

func TestGetByID(t *testing.T) {
	handler := New(&mockService{})

	req := httptest.NewRequest(http.MethodGet, "/task?user_id=1001", nil)
	w := httptest.NewRecorder()
	handler.GetByUserID(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var tasks []models.Task
	err := json.Unmarshal(body, &tasks)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].UserID != 1001 || tasks[1].UserID != 1001 {
		t.Errorf("Expected user ID 1001, got %+v", tasks)
	}
}

func (s *mockService) Create(task models.Task, userID int) (models.Task, error) {
	var taskCreated models.Task
	taskCreated.UserID = userID
	taskCreated.ID = 1001
	taskCreated.Description = "task description"
	taskCreated.Status = true
	return taskCreated, nil
}
func TestCreate(t *testing.T) {
	handler := New(&mockService{})
	taskDummy := models.Task{
		Description: "task description",
		UserID:      1001,
		ID:          101,
		Status:      true,
	}

	jsonTaskDummy, _ := json.Marshal(taskDummy)
	req := httptest.NewRequest(http.MethodPost, "/task", bytes.NewBuffer(jsonTaskDummy))
	w := httptest.NewRecorder()
	handler.Create(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var taskCreated models.Task
	err := json.Unmarshal(body, &taskCreated)
	if err != nil {
		t.Errorf("Unmarshal failed: %v", err)
	}

	if taskCreated.UserID != taskDummy.UserID {
		t.Errorf("Create Task Failed")
	}
}

func (s *mockService) GetByUserID(userID int) ([]models.Task, error) {
	return []models.Task{
		{ID: 1, Description: "Task1", Status: false, UserID: userID},
		{ID: 2, Description: "Task2", Status: true, UserID: userID},
	}, nil

}

func (s *mockService) DeleteTaskById(id int) (int, error) {
	return id, nil
}
func TestDelete(t *testing.T) {
	handler := New(&mockService{})

	req := httptest.NewRequest(http.MethodDelete, "/task?id=1", nil)
	w := httptest.NewRecorder()
	handler.DeleteTaskById(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "Task deleted" {
		t.Errorf("Expected 'Task deleted', got '%s'", string(body))
	}
}

func (s *mockService) Update(task models.Task, taskID int) (models.Task, error) {
	task.ID = taskID
	return task, nil
}

func TestUpdate(t *testing.T) {
	handler := New(&mockService{})
	task := models.Task{
		Description: "Updated Task",
		Status:      true,
		UserID:      1001,
	}

	jsonTask, _ := json.Marshal(task)
	req := httptest.NewRequest(http.MethodPut, "/task?id=1", bytes.NewBuffer(jsonTask))
	w := httptest.NewRecorder()

	handler.Update(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var updated models.Task
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &updated)
	if err != nil {
		t.Errorf("Unmarshal failed: %v", err)
	}

	if updated.Description != "Updated Task" {
		t.Errorf("Expected description to be 'Updated Task', got '%s'", updated.Description)
	}
}

func (s *mockService) GetAll() ([]models.Task, error) {
	return []models.Task{
		{ID: 1, Description: "Task A", Status: true, UserID: 2001},
	}, nil
}

func TestGetAll(t *testing.T) {
	handler := New(&mockService{})
	req := httptest.NewRequest(http.MethodGet, "/task/all", nil)
	w := httptest.NewRecorder()

	handler.GetAll(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var tasks []models.Task
	if err := json.Unmarshal(body, &tasks); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(tasks) != 1 || tasks[0].Description != "Task A" {
		t.Errorf("Unexpected task data: %+v", tasks)
	}
}
