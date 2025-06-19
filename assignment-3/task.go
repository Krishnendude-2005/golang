package main

import (
	"fmt"
	"strings"
)

type Task struct {
	ID          int
	Description string
	Status      bool
}

type TaskManager struct {
	Tasks     []Task
	GetNextID func() int
}

func IDGenerator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

func AddTask(desc *string, tm *TaskManager) {
	id := tm.GetNextID()
	t := Task{
		ID:          id,
		Description: strings.TrimSpace(*desc),
		Status:      false,
	}
	tm.Tasks = append(tm.Tasks, t)
	fmt.Println("Task Added:", t.ID, "-", t.Description)
}

func CompleteTask(id int, tm *TaskManager) {
	for i := range tm.Tasks {
		if tm.Tasks[i].ID == id {
			tm.Tasks[i].Status = true
			fmt.Printf("Marked Task %d as Completed: %s\n", tm.Tasks[i].ID, tm.Tasks[i].Description)
			return
		}
	}
	fmt.Println("Task ID not found!")
}

func ListPendingTasks(tm *TaskManager) []Task {
	var pending []Task
	for _, t := range tm.Tasks {
		if !t.Status {
			pending = append(pending, t)
		}
	}
	return pending
}
