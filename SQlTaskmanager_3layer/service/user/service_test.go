package user_test

import (
	"context"
	"errors"
	"testing"

	"SQLTaskmanager_3layer/models"
	"SQLTaskmanager_3layer/service/user"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"

	"go.uber.org/mock/gomock"
)

func TestService_Create_Valid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := user.New(mockStore)

	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	input := models.User{ID: 1001, TaskName: "Test User"}
	expected := models.User{ID: 1001, TaskName: "Test User"}

	mockStore.EXPECT().Create(ctx, input).Return(expected, nil)

	got, err := svc.Create(ctx, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != expected.ID || got.TaskName != expected.TaskName {
		t.Errorf("expected %+v, got %+v", expected, got)
	}
}

func TestService_Create_Invalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := user.New(mockStore)

	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	input := models.User{} // TaskName is empty, should fail validation

	_, err := svc.Create(ctx, input)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}
}

func TestService_GetById_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := user.New(mockStore)

	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	expected := models.User{ID: 101, TaskName: "User 101"}
	mockStore.EXPECT().GetById(ctx, 101).Return(expected, nil)

	got, err := svc.GetById(ctx, 101)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != expected.ID || got.TaskName != expected.TaskName {
		t.Errorf("expected %+v, got %+v", expected, got)
	}
}

func TestService_GetById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := user.New(mockStore)

	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	mockStore.EXPECT().GetById(ctx, 999).Return(models.User{}, errors.New("not found"))

	_, err := svc.GetById(ctx, 999)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
