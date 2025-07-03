package task_test

import (
	"context"
	"errors"
	"testing"

	"SQLTaskmanager_3layer/models"
	"SQLTaskmanager_3layer/service/task"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"

	"go.uber.org/mock/gomock"
)

func TestService_Create_Valid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := task.New(mockStore)

	// explicit context
	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	input := models.Task{Description: "dummy task1", Status: true, UserID: 42}
	expected := models.Task{ID: 11, Description: "dummy task1", Status: true, UserID: 42}

	mockStore.EXPECT().Create(ctx, input, 42).Return(expected, nil)

	got, err := svc.Create(ctx, input, 42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.ID != expected.ID {
		t.Errorf("expected ID %d, got %d", expected.ID, got.ID)
	}
}

func TestService_Create_Invalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := task.New(mockStore)

	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	input := models.Task{} // invalid
	_, err := svc.Create(ctx, input, 0)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := task.New(mockStore)

	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	expected := []models.Task{{ID: 101, Description: "dummy", Status: false, UserID: 9}}

	mockStore.EXPECT().GetById(ctx, 9).Return(expected, nil)

	tasks, err := svc.GetById(ctx, 9)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tasks) != 1 || tasks[0].ID != 101 {
		t.Errorf("unexpected tasks: %+v", tasks)
	}
}

func TestService_DeleteTaskById_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := task.New(mockStore)

	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	mockStore.EXPECT().DeleteTaskById(ctx, 5).Return(5, nil)

	id, err := svc.DeleteTaskById(ctx, 5)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	if id != 5 {
		t.Errorf("expected id 5, got %d", id)
	}
}

func TestService_DeleteTaskById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := task.New(mockStore)

	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	mockStore.EXPECT().DeleteTaskById(ctx, 7).Return(0, errors.New("some error"))

	_, err := svc.DeleteTaskById(ctx, 7)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestService_Update_Valid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := task.New(mockStore)

	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	input := models.Task{Description: "Updated", Status: false, UserID: 55}
	expected := models.Task{ID: 77, Description: "Updated", Status: false, UserID: 55}

	mockStore.EXPECT().Update(ctx, input, 77).Return(expected, nil)

	got, err := svc.Update(ctx, input, 77)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != 77 {
		t.Errorf("expected ID 77, got %d", got.ID)
	}
}

func TestService_Update_Invalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := task.New(mockStore)

	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	input := models.Task{} // invalid
	_, err := svc.Update(ctx, input, 0)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	svc := task.New(mockStore)

	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	expected := []models.Task{{ID: 1}, {ID: 2}}

	mockStore.EXPECT().GetAll(ctx).Return(expected, nil)

	got, err := svc.GetAll(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(got))
	}
}
