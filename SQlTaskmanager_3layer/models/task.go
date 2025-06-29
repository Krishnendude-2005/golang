package models

import "fmt"

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	UserID      int    `json:"user_id"`
}

func (t *Task) Validate() error {
	if t.Description == "" {
		return fmt.Errorf("task description is required")
	}
	return nil
}
