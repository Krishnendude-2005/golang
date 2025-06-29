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

	// Task Layers - store, service, handler
	tStore := taskStore.New(db)           // Store depends on DB.
	tService := taskService.New(tStore)   // Service depends on Store.
	tHandler := taskHandler.New(tService) // Handlers depends on Service.

	// User Layers - store, service, handler
	uStore := userStore.New(db)
	uService := userService.New(uStore)
	uHandler := userHandler.New(uService)

	// Task Handlers.
	http.HandleFunc("/task/add", tHandler.Create)
	http.HandleFunc("/task/user", tHandler.GetByUserID)
	http.HandleFunc("/task/delete", tHandler.DeleteTaskById)
	http.HandleFunc("/task/update", tHandler.Update)
	http.HandleFunc("/task/all", tHandler.GetAll)
	//User Handlers.
	http.HandleFunc("/user/add", uHandler.Create)
	http.HandleFunc("/user/find", uHandler.GetByID)
	// Server Running.
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
