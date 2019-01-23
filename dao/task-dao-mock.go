package dao

import (
	"errors"
	"project/model"

	uuid "github.com/satori/go.uuid"
)

var _ TaskDao = (*TaskDaoMock)(nil)

type TaskDaoMock struct {
	storage map[string]*model.Task
}

func NewTaskDAOMock() TaskDao {
	daoMock := &TaskDaoMock{
		storage: make(map[string]*model.Task),
	}
	return daoMock
}

func (dao *TaskDaoMock) Get(id string) (*model.Task, error) {
	task, ok := dao.storage[id]

	if !ok {
		return nil, errors.New("Task not found with id -> " + id)
	}
	return task, nil
}

func (dao *TaskDaoMock) Upsert(task *model.Task) (*model.Task, error) {

	if task.ID == "" {
		task.ID = uuid.NewV4().String()
	}

	dao.storage[task.ID] = task
	return task, nil
}

func (dao *TaskDaoMock) Delete(id string) error {

	delete(dao.storage, id)

	return nil
}
