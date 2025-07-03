package task_test

import (
	"SQLTaskmanager_3layer/handler/task"
	"SQLTaskmanager_3layer/models"

	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
	gofrHTTP "gofr.dev/pkg/gofr/http"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test: Create success.
func TestCreate_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSvc := NewMockService(ctrl)
	handler := task.New(mockSvc)

	input := models.Task{Description: "Test", Status: false, UserID: 1001}
	expected := models.Task{ID: 1, Description: "Test", Status: false, UserID: 1001}
	body, _ := json.Marshal(input)

	req := httptest.NewRequest(http.MethodPost, "/task", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	mockC, _ := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	_ = ctx.Bind(&input)
	mockSvc.EXPECT().Create(ctx, input, input.UserID).Return(expected, nil)

	resp, err := handler.Create(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, ok := resp.(models.Task)
	if !ok || got.ID != 1 {
		t.Errorf("expected ID 1, got %v", resp)
	}
}

func TestCreate_InvalidJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	handler := task.New(mockSvc)

	// Malformed JSON
	req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader("{bad json"))
	req.Header.Set("Content-Type", "application/json")

	mockC, _ := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	// No expectations because binding fails before calling service
	_, err := handler.Create(ctx)
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestCreate_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	handler := task.New(mockSvc)

	input := models.Task{Description: "Fail", Status: false, UserID: 1}
	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/task", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	mockC, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	_ = ctx.Bind(&input)
	mockSvc.EXPECT().Create(ctx, input, input.UserID).Return(models.Task{}, errors.New("fail"))

	_, err := handler.Create(ctx)
	if err == nil {
		t.Error("expected service error")
	}
}

func TestGetById_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	handler := task.New(mockSvc)

	expected := []models.Task{{ID: 1, Description: "Hello", UserID: 1001}}

	req := httptest.NewRequest(http.MethodGet, "/task/1001", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1001"})
	mockC, _ := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	mockSvc.EXPECT().GetById(ctx, 1001).Return(expected, nil)

	resp, err := handler.GetById(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, ok := resp.([]models.Task)
	if !ok || len(got) != 1 || got[0].UserID != 1001 {
		t.Error("unexpected result in GetById")
	}
}

func TestGetById_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler := task.New(NewMockService(ctrl))
	req := httptest.NewRequest(http.MethodGet, "/task/abc", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})

	mockC, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	_, err := handler.GetById(ctx)
	if err == nil {
		t.Error("expected error for invalid ID")
	}
}

func TestDelete_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	handler := task.New(mockSvc)

	req := httptest.NewRequest(http.MethodDelete, "/task/delete/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	mockC, _ := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	mockSvc.EXPECT().DeleteTaskById(ctx, 1).Return(1, nil)

	resp, err := handler.DeleteTaskById(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, ok := resp.(int)
	if !ok || got != 1 {
		t.Errorf("expected 1, got %v", resp)
	}
}

func TestDelete_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler := task.New(NewMockService(ctrl))
	req := httptest.NewRequest(http.MethodDelete, "/task/delete/abc", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})

	mockC, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	_, err := handler.DeleteTaskById(ctx)
	if err == nil {
		t.Error("expected error for invalid ID")
	}
}

// Test: Update success
func TestUpdate_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	handler := task.New(mockSvc)

	input := models.Task{Description: "Hello", Status: false, UserID: 1001}
	updated := models.Task{ID: 1, Description: "Updated", Status: true, UserID: 1001}
	body, _ := json.Marshal(input)

	// Use gorilla/mux to simulate URL vars correctly
	req := httptest.NewRequest(http.MethodPut, "/task/update/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": "1"}) // ✅ This is important

	mockC, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	_ = ctx.Bind(&input)

	// Match expected input values correctly
	mockSvc.EXPECT().Update(gomock.Any(), input, 1).Return(updated, nil)

	resp, err := handler.Update(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := resp.(models.Task)
	if got.Description != "Updated" {
		t.Error("task not updated")
	}
}

func TestUpdate_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler := task.New(NewMockService(ctrl))
	req := httptest.NewRequest(http.MethodPut, "/task/update/xyz", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "xyz"})

	mockC, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	_, err := handler.Update(ctx)
	if err == nil {
		t.Error("expected error for invalid ID")
	}
}

func TestGetAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	handler := task.New(mockSvc)

	expected := []models.Task{{ID: 1, UserID: 1001}}
	req := httptest.NewRequest(http.MethodGet, "/task/all", nil)

	mockC, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	mockSvc.EXPECT().GetAll(ctx).Return(expected, nil)

	resp, err := handler.GetAll(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, ok := resp.([]models.Task)
	if !ok || len(got) != 1 {
		t.Errorf("expected 1 task, got %v", resp)
	}
}
