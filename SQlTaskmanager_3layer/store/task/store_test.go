package task

import (
	"SQLTaskmanager_3layer/models"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("Error Occured")
	}

	str := New(db)

	mock.ExpectExec("INSERT INTO tasks (Description, Status, UserID) VALUES (?, ?, ?)").WithArgs("task1", false, 101).WillReturnResult(sqlmock.NewResult(1, 1))

	newUser := models.Task{
		ID:          1,
		Description: "task1",
		Status:      false,
		UserID:      101,
	}

	returnedUser, err := str.Create(newUser, 101)

	if err != nil {
		t.Error("an error occured", err)
	}
	if returnedUser.Description != "task1" {
		t.Errorf("got user description %s, wanted task1", returnedUser.Description)
	}
	if newUser.UserID != 101 {
		t.Errorf("got user ID %d, wanted 101", newUser.UserID)
	}
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error("Error Occurred")
	}
	str := New(db)

	returnedRow := sqlmock.NewRows([]string{"ID", "Description", "Status", "UserID"}).AddRow(2, "task1", false, 101)
	mock.ExpectQuery("SELECT ID, Description, Status, UserID FROM tasks WHERE UserID = ?").WithArgs(2).WillReturnRows(returnedRow)

	ansRow, err := str.GetById(2)
	if err != nil {
		t.Error("an error occurred")
	}
	if len(ansRow) == 0 {
		t.Error("length should be greater than zero")
	}

	for i := range ansRow {
		if ansRow[i].Description == "task1" {
			if ansRow[i].UserID != 101 {
				t.Errorf("got user ID %d, wanted 101", ansRow[i].UserID)
			}
		}
	}
}
func TestDeleteTaskById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error("Error Occurred")
	}
	str := New(db)

	mock.ExpectExec("DELETE FROM tasks WHERE ID = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	userID, err := str.DeleteTaskById(1)
	if err != nil {
		t.Error("an error occured", err)
	}

	if userID != 1 {
		t.Errorf("got user ID %d, wanted 1", userID)
	}
}
func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error("Error Occurred")
	}
	str := New(db)

	updatedTask := models.Task{
		ID:          1,
		Description: "updated1",
		Status:      true,
		UserID:      101,
	}
	mock.ExpectExec("UPDATE tasks SET Description = ?, Status = ?, UserID = ? WHERE ID = ?").WithArgs("updated1", true, 101, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	ans, err := str.Update(updatedTask, 1)
	if err != nil {
		t.Error("an error occured", err)
	}

	if updatedTask.Description != ans.Description {
		t.Errorf("got description %s, wanted updated1", ans.Description)
	}

	if updatedTask.Status != ans.Status {
		t.Errorf("got status %v, wanted true", ans.Status)
	}

	if updatedTask.UserID != ans.UserID {
		t.Errorf("got user ID %d, wanted 101", ans.UserID)
	}
}
func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error("Error Occurred")
	}
	str := New(db)

	rows := sqlmock.NewRows([]string{"ID", "Description", "Status", "UserID"}).AddRow(1, "task1", false, 101).AddRow(2, "task2", false, 102)
	mock.ExpectQuery("SELECT ID, Description, Status, UserID FROM tasks").WillReturnRows(rows)
	ansUsers, err := str.GetAll()

	if err != nil {
		t.Error("an error occured", err)
	}
	if len(ansUsers) != 2 {
		t.Error("length should be greater than zero")
	}
	if ansUsers[0].Description != "task1" {
		t.Errorf("got user description %s, wanted task1", ansUsers[0].Description)
	}
	if ansUsers[1].Description != "task2" {
		t.Errorf("got user description %s, wanted task2", ansUsers[1].Description)
	}
	if ansUsers[0].UserID != 101 {
		t.Errorf("got user ID %d, wanted 101", ansUsers[0].UserID)
	}
	if ansUsers[1].UserID != 102 {
		t.Errorf("got user ID %d, wanted 101", ansUsers[1].UserID)
	}
}
