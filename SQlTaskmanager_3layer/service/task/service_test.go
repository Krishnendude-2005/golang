package task

import (
	"SQLTaskmanager_3layer/models"
	"testing"
)

type mockStore struct {
	// empty
}

func (m *mockStore) Create(task models.Task, userID int) (models.Task, error) {
	var taskCreated models.Task
	taskCreated.UserID = userID
	taskCreated.ID = 1001
	taskCreated.Description = "task description"
	taskCreated.Status = true
	return taskCreated, nil
}
func TestCreate(t *testing.T) {
	svc := New(&mockStore{})
	task := models.Task{
		UserID:      101,
		ID:          1001,
		Description: "task description",
		Status:      true,
	}

	returnedTask, err := svc.Create(task, 101)
	if err != nil {
		t.Error(err)
	}

	if returnedTask.UserID != 101 {
		t.Error("Expected 101, got ", returnedTask.UserID)
	}
	if returnedTask.ID != 1001 {
		t.Error("Expected 1001, got ", returnedTask.ID)
	}
	if returnedTask.Description != "task description" {
		t.Error("Expected task description, got ", returnedTask.Description)
	}
	if returnedTask.Status != true {
		t.Error("Expected true, got ", returnedTask.Status)
	}

}

func (m *mockStore) GetById(userID int) ([]models.Task, error) {
	//TODO implement me
	tasks := []models.Task{}

	if userID == 1 {
		task := models.Task{
			ID: 1,
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func Test_GetByID(t *testing.T) {
	svc := New(&mockStore{})

	res, err := svc.GetById(1)
	if res[0].ID != 1 || err != nil {
		t.Errorf("Somthing wrong")
	}
}

func (m *mockStore) Update(task models.Task, taskID int) (models.Task, error) {
	var fTask models.Task
	fTask.ID = taskID
	fTask.UserID = task.UserID
	fTask.Description = task.Description
	fTask.Status = task.Status

	return fTask, nil
}

func Test_Update(t *testing.T) {
	svc := New(&mockStore{})
	upDatedtask := models.Task{
		ID:          101,
		Description: "description",
		Status:      true,
		UserID:      1001,
	}
	res, _ := svc.Update(upDatedtask, 101)
	if res.Status != upDatedtask.Status {
		t.Errorf("Somthing wrong")
	}
	if res.UserID != upDatedtask.UserID {
		t.Errorf("Somthing wrong")
	}
}

func (m *mockStore) DeleteTaskById(id int) (int, error) {
	//svc := New(&mockStore{})
	var tasks []models.Task
	for i := 0; i < 3; i++ {
		var task models.Task
		task.ID = i
		task.UserID = i + 100
		task.Description = "task description"
		task.Status = true

		tasks = append(tasks, task)

	}

	for i := 0; i < len(tasks); i++ {
		if id == tasks[i].ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	return id, nil
}
func Test_DeleteTaskById(t *testing.T) {
	svc := New(&mockStore{})

	id, err := svc.DeleteTaskById(1)
	if err != nil {
		t.Error(err)
	}
	if id != 1 {
		t.Error("Expected 1, got ", id)
	}

}
func (m *mockStore) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	for i := 0; i < 3; i++ {
		var task models.Task
		task.ID = i
		task.UserID = i + 100
		task.Description = "task description"
		task.Status = true

		tasks = append(tasks, task)

	}
	return tasks, nil
}
func Test_GetAll(t *testing.T) {
	svc := New(&mockStore{})
	tasks, err := svc.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(tasks) != 3 {
		t.Error("Expected length 3, got ", len(tasks))
	}

	for i := 0; i < 3; i++ {
		if tasks[i].ID != i {
			t.Error("Expected ", tasks[i].ID, "got ", tasks[i].ID)
		}
		if tasks[i].UserID != i+100 {
			t.Error("Expected ", tasks[i].UserID, "got ", tasks[i].UserID)
		}
	}
}
