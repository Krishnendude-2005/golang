package models

import "fmt"

type User struct {
	ID       int    `json:"ID"`
	TaskName string `json:"TaskName"`
}

func (u User) Validate() error {
	if u.TaskName == "" {
		return fmt.Errorf("user task name is required")
	}
	return nil
}
