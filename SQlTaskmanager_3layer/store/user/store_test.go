package user_test

import (
	"context"
	"testing"

	"SQLTaskmanager_3layer/models"
	"SQLTaskmanager_3layer/store/user"

	"github.com/DATA-DOG/go-sqlmock"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
)

func TestUserStore_Create(t *testing.T) {
	mockContainer, mocks := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	store := user.New()

	// ✅ Updated to expect ID and TaskName
	mocks.SQL.
		ExpectExec(`INSERT INTO users (ID, TaskName) VALUES (?, ?)`).
		WithArgs(101, "Test User").
		WillReturnResult(sqlmock.NewResult(0, 1)) // ID doesn't matter here

	// ✅ Include ID in test input
	input := models.User{ID: 101, TaskName: "Test User"}
	created, err := store.Create(ctx, input)

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if created.ID != 101 {
		t.Errorf("expected ID 101, got: %d", created.ID)
	}
	if created.TaskName != "Test User" {
		t.Errorf("expected TaskName 'Test User', got: %s", created.TaskName)
	}
}

func TestUserStore_GetById(t *testing.T) {
	mockContainer, mocks := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	store := user.New()

	row := sqlmock.NewRows([]string{"ID", "TaskName"}).
		AddRow(42, "User 42")

	mocks.SQL.
		ExpectQuery(`SELECT ID, TaskName FROM users WHERE ID = ?`).
		WithArgs(42).
		WillReturnRows(row)

	userData, err := store.GetById(ctx, 42)

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if userData.ID != 42 {
		t.Errorf("expected ID 42, got: %d", userData.ID)
	}
	if userData.TaskName != "User 42" {
		t.Errorf("expected TaskName 'User 42', got: %s", userData.TaskName)
	}
}
