package main

import (
	"bytes"
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// Helper function for Creating DB Connection.
func GetTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/task_manager")
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("Test DB not reachable: %v", err)
	}

	return db
}

// Helper function for clearing the Previous tasks , to check functionality.
func ClearTaskTable(t *testing.T, db *sql.DB) {
	_, err := db.Exec("DELETE FROM tasks")
	if err != nil {
		t.Fatalf("Failed to clear tasks table: %v", err)
	}
}

// TaskManager - containing the DB connection & clearing previous data.
func GetTestTaskManager(t *testing.T) *TaskManager {
	db := GetTestDB(t)
	ClearTaskTable(t, db)

	return &TaskManager{DB: db}
}

// Helper function for inserting in DB.
func insertTestTask(t *testing.T, db *sql.DB, desc string, status bool) int64 {
	res, err := db.Exec("INSERT INTO tasks (Description, Status) VALUES (?, ?)", desc, status)
	if err != nil {
		t.Fatal("Insert failed:", err)
	}

	id, _ := res.LastInsertId()

	return id
}

// AddTaskHandler statusOK functionality check.
func TestAddTaskHandler_StatusOK(t *testing.T) {
	tm := GetTestTaskManager(t)
	body := `{"Description":"New Task","Status":false}`

	req := httptest.NewRequest(http.MethodPost, "/task/add", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	tm.AddTaskHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

// AddTaskHandler empty body handler check.
func TestAddTaskHandler_EmptyBody(t *testing.T) {
	tm := GetTestTaskManager(t)
	req := httptest.NewRequest(http.MethodPost, "/task/add", http.NoBody)
	w := httptest.NewRecorder()

	tm.AddTaskHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 BadRequest, got %d", w.Code)
	}
}

// AddTaskHandler invalid JSON body handler check.
func TestAddTaskHandler_InvalidJSON(t *testing.T) {
	tm := GetTestTaskManager(t)
	req := httptest.NewRequest(http.MethodPost, "/task/add", bytes.NewBufferString("bad json"))
	w := httptest.NewRecorder()

	tm.AddTaskHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 BadRequest, got %d", w.Code)
	}
}

// AddTaskHandler method-not-allowed handler check.
func TestAddTaskHandler_WrongMethod(t *testing.T) {
	tm := GetTestTaskManager(t)
	req := httptest.NewRequest(http.MethodGet, "/task/add", http.NoBody)
	w := httptest.NewRecorder()

	tm.AddTaskHandler(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405 MethodNotAllowed, got %d", w.Code)
	}
}

// FindTaskHandler status OK check.
func TestFindTaskHandler_StatusOK(t *testing.T) {
	tm := GetTestTaskManager(t)
	id := insertTestTask(t, tm.DB, "Find Me", false)

	req := httptest.NewRequest(http.MethodGet, "/task/find?id="+strconv.Itoa(int(id)), http.NoBody)
	w := httptest.NewRecorder()

	tm.FindTaskHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

// FindTaskHandler id not found handler check.
func TestFindTaskHandler_NotFound(t *testing.T) {
	tm := GetTestTaskManager(t)
	req := httptest.NewRequest(http.MethodGet, "/task/find?id=99999", http.NoBody)
	w := httptest.NewRecorder()

	tm.FindTaskHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404 NotFound, got %d", w.Code)
	}
}

// FindTaskHandler id empty handler check.
func TestFindTaskHandler_MissingID(t *testing.T) {
	tm := GetTestTaskManager(t)
	req := httptest.NewRequest(http.MethodGet, "/task/find", http.NoBody)
	w := httptest.NewRecorder()

	tm.FindTaskHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 BadRequest, got %d", w.Code)
	}
}

// CompleteTaskHandler check for status - 200 expected.
func TestCompleteTaskHandler(t *testing.T) {
	tm := GetTestTaskManager(t)
	id := insertTestTask(t, tm.DB, "Complete Me", false)

	req := httptest.NewRequest(http.MethodGet, "/task/complete?id="+strconv.Itoa(int(id)), http.NoBody)
	w := httptest.NewRecorder()

	tm.CompleteTaskHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

// CompleteTaskHandler method-not-allowed check.
func TestCompleteTaskHandler_WrongMethod(t *testing.T) {
	tm := GetTestTaskManager(t)
	req := httptest.NewRequest(http.MethodPost, "/task/complete?id=1", http.NoBody)
	w := httptest.NewRecorder()

	tm.CompleteTaskHandler(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405 MethodNotAllowed, got %d", w.Code)
	}
}

// UpdateTaskHandler check - expected 200.
func TestUpdateTaskHandler(t *testing.T) {
	tm := GetTestTaskManager(t)
	id := insertTestTask(t, tm.DB, "Old Desc", false)

	req := httptest.NewRequest(http.MethodPut, "/task/update?id="+strconv.Itoa(int(id))+"&desc=Updated", http.NoBody)
	w := httptest.NewRecorder()

	tm.UpdateTaskHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

// DeleteTaskHandler check - expected 200.
func TestDeleteTaskHandler(t *testing.T) {
	tm := GetTestTaskManager(t)
	id := insertTestTask(t, tm.DB, "Delete Me", false)

	req := httptest.NewRequest(http.MethodDelete, "/task/delete?id="+strconv.Itoa(int(id)), http.NoBody)
	w := httptest.NewRecorder()

	tm.DeleteTaskHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

// TaskHandler check - expected 200.
func TestTaskHandler(t *testing.T) {
	tm := GetTestTaskManager(t)
	insertTestTask(t, tm.DB, "Task 1", false)
	insertTestTask(t, tm.DB, "Task 2", false)

	req := httptest.NewRequest(http.MethodGet, "/task", http.NoBody)
	w := httptest.NewRecorder()

	tm.TaskHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}

	body, _ := io.ReadAll(w.Body)
	if !bytes.Contains(body, []byte("Pending Tasks")) {
		t.Error("Expected response to contain 'Pending Tasks'")
	}
}
