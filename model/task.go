package model

import (
	"time"
)

// TaskStatus is the status
type TaskStatus int

// Task model
type Task struct {
	ID               string     `json:"id,omitempty"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	Status           TaskStatus `json:"status"`
	CreationDate     time.Time  `json:"creationDate"`
	ModificationTime time.Time  `json:"modificationDate"`
}

const (
	// StatusTodo values
	StatusTodo TaskStatus = iota
	// StatusProgress value
	StatusProgress
	// StatusDone when is done
	StatusDone
)
