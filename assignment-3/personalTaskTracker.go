package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// global ID generator
func idGenerator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

func main() {
	// task structure
	type task struct {
		id          int
		description string
		status      bool
	}

	// task manager -- instance of taskManager for passing it to functions
	type taskManager struct {
		tasks     []task
		getNextId func() int
	}

	// Initialize task manager locally-- making this instance to pass it to functions --actual value change
	manager := taskManager{
		tasks:     []task{},
		getNextId: idGenerator(),
	}

	// Add Task Function (local)
	addTask := func(desc *string, tm *taskManager) {
		id := tm.getNextId()
		t := task{
			id:          id,
			description: strings.TrimSpace(*desc),
			status:      false,
		}
		tm.tasks = append(tm.tasks, t)
		fmt.Println("Task Added:", t.id, "-", t.description)
	}

	// List Pending Tasks Function (local)
	listTasks := func(tm *taskManager) {
		fmt.Println("\nPending Tasks:")
		for _, t := range tm.tasks {
			if t.status == false {
				fmt.Printf("%d - %s\n", t.id, t.description)
			}
		}
	}

	// Complete Task Function (local)
	completeTask := func(id int, tm *taskManager) {
		for i := range tm.tasks {
			if tm.tasks[i].id == id {
				tm.tasks[i].status = true
				fmt.Printf("Marked Task %d as Completed: %s\n", tm.tasks[i].id, tm.tasks[i].description)
				return
			}
		}
		fmt.Println("Task ID not found!")
	}

	// reading tasks

	//giving options

	//var operation int

	fmt.Println("Enter number of tasks to add:")
	var taskLength int
	fmt.Scan(&taskLength)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < taskLength; i++ {
		fmt.Printf("Enter task no%d description: ", i+1)
		input, _ := reader.ReadString('\n')
		addTask(&input, &manager)
	}

	listTasks(&manager)

	fmt.Println("\nEnter the task ID to mark as completed:")
	var completedTaskID int
	fmt.Scan(&completedTaskID)

	completeTask(completedTaskID, &manager)

	listTasks(&manager)
}
