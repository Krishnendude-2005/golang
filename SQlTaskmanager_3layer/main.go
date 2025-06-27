package main

import (
	taskHandler "SQLTaskmanager_3layer/handler/task"
	userHandler "SQLTaskmanager_3layer/handler/user"
	taskService "SQLTaskmanager_3layer/service/task"
	userService "SQLTaskmanager_3layer/service/user"
	taskStore "SQLTaskmanager_3layer/store/task"
	userStore "SQLTaskmanager_3layer/store/user"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	// Connecting to MySQL
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/task_manager")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Task Layers
	tStore := taskStore.New(db)
	tService := taskService.New(tStore)
	tHandler := taskHandler.New(tService)

	// User Layers
	uStore := userStore.New(db)
	uService := userService.New(uStore)
	uHandler := userHandler.New(uService)

	// Task Routes
	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		tHandler.Create(w, r)
	})

	http.HandleFunc("/task/find", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		tHandler.GetById(w, r)
	})

	// User Routes
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		uHandler.Create(w, r)
	})

	http.HandleFunc("/user/find", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		uHandler.GetByID(w, r)
	})

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
