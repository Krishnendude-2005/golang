package main

import (
	"assignment-9"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tm *assignment_9.TaskManager

func setup() {
	tm = &assignment_9.TaskManager{
		Tasks:     []assignment_9.Task{},
		GetNextID: assignment_9.IDGenerator(), // assuming this returns a func() int
	}
	tm.Tasks = append(tm.Tasks, assignment_9.Task{ID: 1, Description: "Task 1", Status: false})
	tm.Tasks = append(tm.Tasks, assignment_9.Task{ID: 2, Description: "Task 2", Status: false})
}

func TestTaskHandler(t *testing.T) {
	setup()

	req := httptest.NewRequest(http.MethodGet, "/task/add", nil)
	writer := httptest.NewRecorder()

	tm.AddTaskHandler(writer, req)
	resp := writer.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code , wanted %v", http.StatusOK)
	}
}
func TestAddTaskHandler(t *testing.T) {
	setup()
	length := len(tm.Tasks)
	req := httptest.NewRequest("GET", "/task/add", nil)
	writer := httptest.NewRecorder()

	tm.AddTaskHandler(writer, req)

	if len(tm.Tasks) != length+1 {
		t.Errorf("Expected task to be added")
	}
}
func TestUpdateTaskHandler(t *testing.T) {
	setup()
	//changedDesc := "Changed desc"
	length := len(tm.Tasks)
	req := httptest.NewRequest("PUT", "/task/update", nil)
	writer := httptest.NewRecorder()

	tm.UpdateTaskHandler(writer, req)

	if len(tm.Tasks) != length {
		t.Errorf("Expected the same length after alteration")
	}
}

func TestFindTaskHandler(t *testing.T) {
	setup()
	req := httptest.NewRequest(http.MethodGet, "/task/find/2", nil)
	writer := httptest.NewRecorder()

	tm.FindTaskHandler(writer, req)
	resp := writer.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200")
	}

}
func TestCompleteTaskHandler(t *testing.T) {
	setup()
	// Create a request to mark task ID 2 as completed
	req := httptest.NewRequest("PUT", "/task/complete?id=2", nil)
	writer := httptest.NewRecorder()

	tm.CompleteTaskHandler(writer, req)

	for _, val := range tm.Tasks {
		if val.ID == 2 {
			if val.Status != true {
				t.Errorf("expected the task status to be completed and marked true")
			}
		}
	}
}

func TestTaskManager_DeleteTaskHandler(t *testing.T) {
	setup()
	length := len(tm.Tasks)
	req := httptest.NewRequest(http.MethodDelete, "/task/delete/2", nil)
	writer := httptest.NewRecorder()

	tm.DeleteTaskHandler(writer, req)
	//resp := writer.Result()
	if len(tm.Tasks) != length-1 {
		t.Errorf("Expected 1 less  after deletion")
	}

}
