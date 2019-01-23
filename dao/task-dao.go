package dao

import "project/model"

// TaskDao interfaces is used to know all the allowed methods
type TaskDao interface {
	Get(id string) (*model.Task, error)

	Upsert(task *model.Task) (*model.Task, error)

	Delete(id string) error
}
