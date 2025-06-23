package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var tm *TaskManager

func TestgetTaskHandler(t *testing.T) {

}
func TestaddTaskHandler(t *testing.T) {
	length := len(tm.Tasks)
	req := httptest.NewRequest("GET", "/task/add", nil)
	writer := httptest.NewRecorder()

	tm.addTaskHandler(writer, req)

	if len(tm.Tasks) != length+1 {
		t.Errorf("Expected task to be added")
	}
}
func TestupdateTaskHandler(t *testing.T) {
	//changedDesc := "Changed desc"
	length := len(tm.Tasks)
	req := httptest.NewRequest("PUT", "/task/update", nil)
	writer := httptest.NewRecorder()

	tm.updateTaskHandler(writer, req)

	if len(tm.Tasks) != length {
		t.Errorf("Expected the same length after alteration")
	}
}
func TestfindTaskHandler(t *testing.T) {
	id := 2

	req := httptest.NewRequest(http.MethodGet, "/task/find/2", nil)
	writer := httptest.NewRecorder()
}
