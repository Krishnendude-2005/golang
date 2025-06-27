package models

import "fmt"

type User struct {
	ID       int    `json:"id"`
	TaskName string `json:"task_name"`
}

func (u User) Validate() error {
	if u.TaskName == "" {
		return fmt.Errorf("user task name is required")
	}
	return nil
}
