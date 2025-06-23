package main

import (
	"testing"
)

// var taskdetails = &TaskManager{}.

func TestAddTask(t *testing.T) {
	manager := TaskManager{
		Tasks:     []Task{},
		GetNextID: IDGenerator(),
	}

	desc := "Test Task"
	AddTask(&desc, &manager)

	if len(manager.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(manager.Tasks))
	}

	task := manager.Tasks[0]
	if task.ID != 1 {
		t.Errorf("Expected task ID 1, got %d", task.ID)
	}

	if task.Description != desc {
		t.Errorf("Expected description 'Test Task', got '%s'", task.Description)
	}

	if task.Status != false {
		t.Errorf("Expected task status false, got true")
	}
}

func TestCompleteTask(t *testing.T) {
	manager := TaskManager{
		Tasks:     []Task{},
		GetNextID: IDGenerator(),
	}

	desc := "Completed Task"
	AddTask(&desc, &manager)

	CompleteTask(1, &manager)

	if manager.Tasks[0].Status != true {
		t.Errorf("Expected task status true, but found false")
	}
}

func TestListPendingTasks(t *testing.T) {
	PendingTasks = []Task{} // reset global

	manager := TaskManager{
		Tasks:     []Task{},
		GetNextID: IDGenerator(),
	}

	tasks := []string{"Task 1", "Task 2", "Task 3"}
	for _, task := range tasks {
		AddTask(&task, &manager)
	}

	CompleteTask(2, &manager)

	pending := ListPendingTasks(&manager)
	for _, task := range pending {
		if task.ID == 2 {
			t.Errorf("Expected task ID 2 to be completed, but found in pending list")
		}
	}
}
