package models

import "fmt"

type Task struct {
	ID          int    `json:"ID"`
	Description string `json:"Description"`
	Status      bool   `json:"Status"`
	UserID      int    `json:"UserID"`
}

func (t *Task) Validate() error {
	if t.Description == "" {

		return fmt.Errorf("task description is required")
	}

	return nil
}
