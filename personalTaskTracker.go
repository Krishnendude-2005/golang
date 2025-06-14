package main

import (
	"bufio"
	"fmt"
	"os"
)

// GENERATING UNIQUE ID
// function for generating new ID's each time ---- using function closure
func idGenerator() func() int {
	id := 0

	return func() int {
		id++
		return id
	}
}

// calling it outside 1 time to start using it inside addTask function
var getNextId = idGenerator()

// TASK STRUCTURE
type task struct {
	id          int
	description string
	status      bool
}

// SLICE TO STORE TASKS-----globally defining it
var taskSlice []task

func main() {
	//using a slice to store the tasks --- using slice because unsure of its size
	taskSlice = make([]task, 0)

	fmt.Println("Enter No of Tasks to be tracked")
	var tasks int
	fmt.Scan(&tasks)

	for i := 0; i < tasks; i++ {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter a Task to Add ")
		input, _ := reader.ReadString('\n')

		addTask(&input)
	}

	listTasks()

	fmt.Println("Enter the task id to mark as completed")
	var completedTaskid int
	fmt.Scan(&completedTaskid)
	completeTask(completedTaskid)

	listTasks()
}

// ADDING A NEW TASK
// it adds a new task -- only takes the description -- id is generated and deafult status false
func addTask(description *string) {
	//generating id through getNextId function
	id := getNextId()

	//appending new task with existing
	taskSlice = append(taskSlice, task{id, *description, false})

	//giving output as what task is added
	fmt.Println("Task Added : ", id, "-", *description)
	fmt.Printf("\n")
}

// LISTING ALL PENDING TASKS
// using for loop , iterating over taskSlice and printing its values if its not completed ( status false )
func listTasks() {
	fmt.Println("Pending Tasks : ")
	for i := 0; i < len(taskSlice); i++ {
		if taskSlice[i].status == false {
			fmt.Println(taskSlice[i].id, " ", taskSlice[i].description)
			fmt.Printf("\n")
		}
	}
}

// SHOWING THE COMPLETED TASK
// iterating over taskSlice to find the desired task with same id and making its status as marked ( true )
func completeTask(id int) {
	for i := 0; i < len(taskSlice); i++ {
		if taskSlice[i].id == id {
			fmt.Println("Marking task", taskSlice[i].id, "as completed")
			taskSlice[i].status = true
			fmt.Printf("\n")
		}
	}
}
