package user_test

import (
	"SQLTaskmanager_3layer/handler/user"
	"SQLTaskmanager_3layer/models"
	"bytes"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
	gofrHTTP "gofr.dev/pkg/gofr/http"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	h := user.New(mockSvc)

	input := models.User{TaskName: "Hello"}
	expected := models.User{ID: 1, TaskName: "Hello"}

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	mockC, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	_ = ctx.Bind(&input)
	mockSvc.EXPECT().Create(ctx, input).Return(expected, nil)

	resp, err := h.Create(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := resp.(models.User)
	if got.ID != 1 {
		t.Errorf("expected ID 1, got %d", got.ID)
	}
}

func Test_Create_InvalidJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := user.New(NewMockService(ctrl))

	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader([]byte(`{bad json`)))
	req.Header.Set("Content-Type", "application/json")

	mockC, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	_, err := h.Create(ctx)
	if err == nil {
		t.Errorf("expected error for invalid JSON")
	}
}

func Test_GetById_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	h := user.New(mockSvc)

	expected := models.User{ID: 10, TaskName: "Test"}
	req := httptest.NewRequest(http.MethodGet, "/user/10", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "10"})

	mockC, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	mockSvc.EXPECT().GetById(ctx, 10).Return(expected, nil)

	resp, err := h.GetById(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := resp.(models.User)
	if got.ID != 10 {
		t.Errorf("expected ID 10, got %d", got.ID)
	}
}

func Test_GetById_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := user.New(NewMockService(ctrl))

	req := httptest.NewRequest(http.MethodGet, "/user/abc", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})

	mockC, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   gofrHTTP.NewRequest(req),
		Container: mockC,
	}

	_, err := h.GetById(ctx)
	if err == nil {
		t.Errorf("expected error for invalid ID")
	}
}
