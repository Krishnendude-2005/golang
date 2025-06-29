package user

import (
	"SQLTaskmanager_3layer/models"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	str := New(db)

	rows := sqlmock.NewRows([]string{"id", "task_name"}).AddRow(2, "task1")

	mock.ExpectQuery("SELECT ID, TaskName FROM users WHERE ID = ?").WithArgs(2).WillReturnRows(rows)

	user, err := str.GetById(2)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when getting user by ID", err)
	}

	if user.ID != 2 {
		t.Errorf("got user ID %d, wanted 1", user.ID)
	}

	if user.TaskName != "task1" {
		t.Errorf("got user task name %s, wanted task1", user.TaskName)
	}
}
func TestCreate(t *testing.T) {
	userDummy := models.User{
		ID:       2,
		TaskName: "task2",
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("unexpected error opening stub DB: %s", err)
	}

	str := New(db)

	mock.ExpectExec("INSERT INTO users (ID, TaskName) VALUES (?, ?)").
		WithArgs(2, "task2").
		WillReturnResult(sqlmock.NewResult(2, 1))

	ans, err := str.Create(userDummy)
	if err != nil {
		t.Fatalf("unexpected error on Create: %s", err)
	}

	if ans.TaskName != "task2" {
		t.Errorf("got TaskName %s, wanted task2", ans.TaskName)
	}
	if ans.ID != 2 {
		t.Errorf("got ID %d, wanted 2", ans.ID)
	}

}
