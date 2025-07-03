package task_test

import (
	"context"
	"testing"

	"SQLTaskmanager_3layer/models"
	"SQLTaskmanager_3layer/store/task"

	"github.com/DATA-DOG/go-sqlmock"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
)

func TestTaskStore_Create(t *testing.T) {
	mockContainer, mocks := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	store := task.New()

	mocks.SQL.
		ExpectExec(`INSERT INTO tasks (Description, Status, UserID) VALUES (?, ?, ?)`).
		WithArgs("Test Task", true, 1).
		WillReturnResult(sqlmock.NewResult(101, 1))

	input := models.Task{Description: "Test Task", Status: true}
	result, err := store.Create(ctx, input, 1)

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if result.ID != 101 {
		t.Errorf("expected ID 101, got: %d", result.ID)
	}
	if result.UserID != 1 {
		t.Errorf("expected UserID 1, got: %d", result.UserID)
	}
}

func TestTaskStore_GetById(t *testing.T) {
	mockContainer, mocks := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	store := task.New()

	rows := sqlmock.NewRows([]string{"ID", "Description", "Status", "UserID"}).
		AddRow(1, "Task A", true, 2).
		AddRow(2, "Task B", false, 2)

	mocks.SQL.
		ExpectQuery(`SELECT ID, Description, Status, UserID FROM tasks WHERE UserID = ?`).
		WithArgs(2).
		WillReturnRows(rows)

	tasks, err := store.GetById(ctx, 2)

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got: %d", len(tasks))
	}
	if tasks[0].Description != "Task A" {
		t.Errorf("expected 'Task A', got: %s", tasks[0].Description)
	}
}

func TestTaskStore_DeleteTaskById(t *testing.T) {
	mockContainer, mocks := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	store := task.New()

	mocks.SQL.
		ExpectExec(`DELETE FROM tasks WHERE ID = ?`).
		WithArgs(3).
		WillReturnResult(sqlmock.NewResult(0, 1))

	affected, err := store.DeleteTaskById(ctx, 3)

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if affected != 1 {
		t.Errorf("expected 1 row affected, got: %d", affected)
	}
}

func TestTaskStore_Update(t *testing.T) {
	mockContainer, mocks := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	store := task.New()

	taskToUpdate := models.Task{
		Description: "Updated",
		Status:      true,
		UserID:      4,
	}

	mocks.SQL.
		ExpectExec(`UPDATE tasks SET Description = ?, Status = ?, UserID = ? WHERE ID = ?`).
		WithArgs("Updated", true, 4, 10).
		WillReturnResult(sqlmock.NewResult(0, 1))

	updatedTask, err := store.Update(ctx, taskToUpdate, 10)

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if updatedTask.ID != 10 {
		t.Errorf("expected ID 10, got: %d", updatedTask.ID)
	}
	if updatedTask.Description != "Updated" {
		t.Errorf("expected 'Updated', got: %s", updatedTask.Description)
	}
}

func TestTaskStore_GetAll(t *testing.T) {
	mockContainer, mocks := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	store := task.New()

	rows := sqlmock.NewRows([]string{"ID", "Description", "Status", "UserID"}).
		AddRow(1, "Task X", true, 5).
		AddRow(2, "Task Y", false, 6)

	mocks.SQL.
		ExpectQuery(`SELECT ID, Description, Status, UserID FROM tasks`).
		WillReturnRows(rows)

	all, err := store.GetAll(ctx)

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if len(all) != 2 {
		t.Errorf("expected 2 tasks, got: %d", len(all))
	}
	if all[0].UserID != 5 {
		t.Errorf("expected first task UserID 5, got: %d", all[0].UserID)
	}
}
