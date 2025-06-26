package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID          int    `json:"ID"`
	Description string `json:"Description"`
	Status      bool   `json:"Status"`
}

type TaskManager struct {
	DB *sql.DB
}

func (tm *TaskManager) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var task Task
	if errorSms := json.Unmarshal(body, &task); errorSms != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	_, err = tm.DB.Exec("INSERT INTO tasks (Description, Status) VALUES (?, ?)", task.Description, task.Status)
	if err != nil {
		http.Error(w, "Database insert failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(task.Description))
}
func (tm *TaskManager) TaskHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := tm.DB.Query("SELECT ID, Description FROM tasks WHERE Status = FALSE")

	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, e := w.Write([]byte("Pending Tasks:\n"))
	if e != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}

	for rows.Next() {
		var id int

		var desc string

		_ = rows.Scan(&id, &desc)
		line := fmt.Sprintf("%d - %s\n", id, desc)
		_, err := w.Write([]byte(line))

		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)

			return
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

	var foundID int
	err = tm.DB.QueryRow("SELECT ID FROM tasks WHERE ID = ?", id).Scan(&foundID)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, errMsg := w.Write([]byte("Task Found\n"))

	if errMsg != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (tm *TaskManager) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	res, err := tm.DB.Exec("UPDATE tasks SET Status = TRUE WHERE ID = ?", id)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, errSms := w.Write([]byte("Task marked as completed\n"))

	if errSms != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (tm *TaskManager) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	desc := strings.TrimSpace(r.URL.Query().Get("desc"))

	if idStr == "" || desc == "" {
		http.Error(w, "Missing id or desc", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	res, err := tm.DB.Exec("UPDATE tasks SET Description = ? WHERE ID = ?", desc, id)
	if err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, errSms := w.Write([]byte("Task updated\n"))

	if errSms != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (tm *TaskManager) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, errSms := tm.DB.Exec("DELETE FROM tasks WHERE ID = ?", id)
	if errSms != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, errData := w.Write([]byte("Task deleted successfully\n"))

	if errData != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func main() {
	dbURL := "root:root@tcp(127.0.0.1:3306)/task_manager"
	db, err := sql.Open("mysql", dbURL)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}

	tm := &TaskManager{DB: db}

	http.HandleFunc("/task", tm.TaskHandler)
	http.HandleFunc("/task/add", tm.AddTaskHandler)
	http.HandleFunc("/task/find", tm.FindTaskHandler)
	http.HandleFunc("/task/complete", tm.CompleteTaskHandler)
	http.HandleFunc("/task/update", tm.UpdateTaskHandler)
	http.HandleFunc("/task/delete", tm.DeleteTaskHandler)

	log.Println("Server running at http://localhost:8080")

	server := http.Server{
		Addr:        ":8080",
		ReadTimeout: 3 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
