package dao_test

import (
	"fmt"
	"project/dao"
	"project/model"
	"testing"
	// "github.com/satori/go.uuid"
)

func TestTaskDaoMock(t *testing.T) {

	taskDaoMoc, err := dao.GetDao(dao.MockDAO)

	if err != nil {
		t.Error(err)
	}

	taskToSave := model.Task{
		Description: "Description mock",
	}

	taskSaved, err := taskDaoMoc.Upsert(&taskToSave)
	if err != nil {
		t.Error(err)
	} else {
		if taskSaved == taskSaved {
			fmt.Println(taskSaved)
			fmt.Println(taskToSave)
		}
	}
}
