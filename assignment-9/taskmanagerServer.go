// package main
//
// import (
//
//	"fmt"
//	"net/http"
//	"strconv"
//	"strings"
//
// )
//
//	type Task struct {
//		ID          int
//		Description string
//		Status      bool
//	}
//
//	type TaskManager struct {
//		Tasks     []Task
//		GetNextID func() int
//	}
//
//	func IDGenerator() func() int {
//		id := 0
//		return func() int {
//			id++
//			return id
//		}
//	}
//
//	func AddTask(desc *string, tm *TaskManager) {
//		id := tm.GetNextID()
//		t := Task{
//			ID:          id,
//			Description: strings.TrimSpace(*desc),
//			Status:      false,
//		}
//		tm.Tasks = append(tm.Tasks, t)
//		fmt.Println("Task Added:", t.ID, "-", t.Description)
//	}
//
//	func CompleteTask(id int, tm *TaskManager) {
//		for i := range tm.Tasks {
//			if tm.Tasks[i].ID == id {
//				tm.Tasks[i].Status = true
//				fmt.Printf("Marked Task %d as Completed: %s\n", tm.Tasks[i].ID, tm.Tasks[i].Description)
//				return
//			}
//		}
//		fmt.Println("Task ID not found!")
//	}
//
// // Handler for Printing all Tasks which are not Completed
//
//	func (tm *TaskManager) TaskHandler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, "Printing All Tasks ")
//		for _, value := range tm.Tasks {
//			if value.Status == false { //printing tasks that are not completed
//				fmt.Fprintln(w, value.Description, "With ID : ", value.ID)
//			}
//		}
//	}
//
// // Finding particular Task using Task ID
//
//	func (tm *TaskManager) FindTaskHandler(w http.ResponseWriter, r *http.Request) {
//		reqidstr := r.PathValue("id")
//		reqid, err := strconv.Atoi(reqidstr)
//		if err != nil {
//			w.Write([]byte("Error Converting ID"))
//			w.WriteHeader(500)
//		}
//
//		for _, value := range tm.Tasks {
//			if value.ID == reqid {
//				fmt.Fprintln(w, "Task Found : ", value.Description, "ID : ", value.ID)
//			}
//		}
//
// }
//
// // Adding new Task through post request
//
//	func (tm *TaskManager) PostTaskhandler(w http.ResponseWriter, r *http.Request) {
//		description := r.PathValue("description")
//
//		newTask := Task{
//			ID:          tm.GetNextID(),
//			Description: description,
//			Status:      false,
//		}
//
//		tm.Tasks = append(tm.Tasks, newTask)
//		fmt.Fprintln(w, "New Task Added Successfully : ", newTask.Description)
//	}
//
//	func (tm *TaskManager) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
//		reqidstr := r.PathValue("id")
//		reqid, err := strconv.Atoi(reqidstr)
//		if err != nil {
//			w.Write([]byte("Error Converting ID"))
//			w.WriteHeader(500)
//		}
//
//		for i := 0; i < len(tm.Tasks); i++ {
//			if tm.Tasks[i].ID == reqid {
//				fmt.Fprint(w, "Marking Task with ID ", reqid, "As Completed", tm.Tasks[i].Description)
//				tm.Tasks[i].Status = true
//			}
//		}
//	}
//
//	func (tm *TaskManager) PutTaskHandler(w http.ResponseWriter, r *http.Request) {
//		if r.Method != http.MethodPut {
//			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
//			return
//		}
//
//		idStr := r.PathValue("id")
//		newDesc := r.PathValue("description")
//
//		id, err := strconv.Atoi(idStr)
//		if err != nil {
//			http.Error(w, "Invalid ID", http.StatusBadRequest)
//			return
//		}
//
//		for i := range tm.Tasks {
//			if tm.Tasks[i].ID == id {
//				tm.Tasks[i].Description = strings.TrimSpace(newDesc)
//				fmt.Fprintf(w, "Task %d updated successfully: %s\n", id, tm.Tasks[i].Description)
//				return
//			}
//		}
//
//		http.Error(w, "Task not found", http.StatusNotFound)
//	}
//
// var PendingTasks = []Task{}
//
//	func ListPendingTasks(tm *TaskManager) []Task {
//		fmt.Println("\nPending Tasks:")
//		for _, t := range tm.Tasks {
//			if !t.Status {
//				fmt.Printf("%d - %s\n", t.ID, t.Description)
//				PendingTasks = append(PendingTasks, t)
//			}
//		}
//		return PendingTasks
//	}
//
//	func (tm *TaskManager) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
//		if r.Method != http.MethodDelete {
//			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
//			return
//		}
//
//		idStr := r.PathValue("id")
//		id, err := strconv.Atoi(idStr)
//		if err != nil {
//			http.Error(w, "Invalid ID", http.StatusBadRequest)
//			return
//		}
//
//		for i, task := range tm.Tasks {
//			if task.ID == id {
//				// Remove task from slice
//				tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
//				fmt.Fprintf(w, "Task %d deleted successfully\n", id)
//				return
//			}
//		}
//
//		http.Error(w, "Task not found", http.StatusNotFound)
//	}
//
//	func main() {
//		//fmt.Println("Task Manager using WEB SERVER")
//		tm := TaskManager{Tasks: make([]Task, 0), GetNextID: IDGenerator()}
//
//		http.HandleFunc("/task", tm.TaskHandler)
//		http.HandleFunc("/task/{id}", tm.FindTaskHandler)
//		http.HandleFunc("/task/complete/{id}", tm.CompleteTaskHandler)
//		http.HandleFunc("/task/add/{description}", tm.PostTaskhandler)
//		http.HandleFunc("/task/update/{id}/{description}", tm.PutTaskHandler)
//		http.HandleFunc("/task/delete/{id}", tm.DeleteTaskHandler)
//
//		fmt.Print("Server Running")
//		http.ListenAndServe(":8080", nil)
//	}
package main

import (
	"fmt"
	"net/http"
	"strconv"
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

func (tm *TaskManager) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	desc := r.URL.Query().Get("desc")
	if desc == "" {
		http.Error(w, "Missing description", http.StatusBadRequest)
		return
	}

	t := Task{
		ID:          tm.GetNextID(),
		Description: strings.TrimSpace(desc),
		Status:      false,
	}
	tm.Tasks = append(tm.Tasks, t)
	fmt.Fprintln(w, "Added Task:", t.ID, "-", t.Description)
}

func (tm *TaskManager) TaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pending Tasks:")
	for _, t := range tm.Tasks {
		if !t.Status {
			fmt.Fprintf(w, "%d - %s\n", t.ID, t.Description)
		}
	}
}

func (tm *TaskManager) FindTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, t := range tm.Tasks {
		if t.ID == id {
			w.WriteHeader(http.StatusOK)
			val, er := w.Write([]byte("Task Found"))
			if er != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				fmt.Println(val)
			}
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func (tm *TaskManager) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, t := range tm.Tasks {
		if t.ID == id {
			tm.Tasks[i].Status = true
			fmt.Fprintf(w, "Task %d marked as completed\n", id)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func (tm *TaskManager) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT allowed", http.StatusMethodNotAllowed)
		return
	}
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/task/update/"), "/")
	if len(parts) < 2 {
		http.Error(w, "Use /task/update/{id}/{description}", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	newDesc := strings.Join(parts[1:], "/")
	for i, t := range tm.Tasks {
		if t.ID == id {
			tm.Tasks[i].Description = strings.TrimSpace(newDesc)
			fmt.Fprintf(w, "Task %d updated to: %s\n", id, newDesc)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func (tm *TaskManager) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/task/delete/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, t := range tm.Tasks {
		if t.ID == id {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			fmt.Fprintf(w, "Task %d deleted successfully\n", id)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func main() {
	tm := TaskManager{
		Tasks:     []Task{},
		GetNextID: IDGenerator(),
	}

	http.HandleFunc("/task", tm.TaskHandler)
	http.HandleFunc("/task/add", tm.AddTaskHandler)
	http.HandleFunc("/task/find", tm.FindTaskHandler)
	http.HandleFunc("/task/complete", tm.CompleteTaskHandler)
	http.HandleFunc("/task/update/", tm.UpdateTaskHandler) // dynamic path
	http.HandleFunc("/task/delete/", tm.DeleteTaskHandler) // dynamic path

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
