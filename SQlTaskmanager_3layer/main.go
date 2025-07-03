package main

import (
	taskHandler "SQLTaskmanager_3layer/handler/task"
	userHandler "SQLTaskmanager_3layer/handler/user"
	taskService "SQLTaskmanager_3layer/service/task"
	userService "SQLTaskmanager_3layer/service/user"
	taskStore "SQLTaskmanager_3layer/store/task"
	userStore "SQLTaskmanager_3layer/store/user"
	"gofr.dev/examples/using-add-rest-handlers/migrations"
	"gofr.dev/pkg/gofr/datasource"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.Migrate(migrations.All())

	tStore := taskStore.New()
	tService := taskService.New(tStore)
	tHandler := taskHandler.New(tService)

	uStore := userStore.New()
	uService := userService.New(uStore)
	uHandler := userHandler.New(uService)

	// Routes
	app.POST("/task", tHandler.Create)
	app.GET("/task/find/{id}", tHandler.GetById)

	app.POST("/user", uHandler.Create)
	app.GET("/user/find/{id}", uHandler.GetById)
	app.GET("/mysql", MysqlHandler)
	app.Run()
}
func MysqlHandler(c *gofr.Context) (any, error) {
	var value int
	err := c.SQL.QueryRowContext(c, "select 2+2").Scan(&value)
	if err != nil {
		return nil, datasource.ErrorDB{Err: err, Message: "error from sql db"}
	}

	return value, nil
}
