package dao

import (
	"encoding/json"
	"errors"
	"project/model"

	uuid "github.com/satori/go.uuid"
	redis "gopkg.in/redis.v5"
)

var _ TaskDao = (*TaskDAORedis)(nil)

type TaskDAORedis struct {
	redisCli *redis.Client
}

func NewTaskDAORedis(redisCli *redis.Client) TaskDao {
	return &TaskDAORedis{
		redisCli: redisCli,
	}
}

func (dao *TaskDAORedis) Get(id string) (*model.Task, error) {

	resTask, err := dao.redisCli.Get(id).Result()

	if err != nil {
		return nil, err
	} else if len(resTask) == 0 {
		return nil, errors.New("REDIS -> " + id + " does not exist")
	}

	task := model.Task{}

	err = json.Unmarshal([]byte(resTask), &task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (dao *TaskDAORedis) Upsert(task *model.Task) (*model.Task, error) {
	if task.ID == "" {
		task.ID = uuid.NewV4().String()
	}

	resTask, err := json.Marshal(task)
	if err != nil {
		return nil, err
	}

	status := dao.redisCli.Set(task.ID, string(resTask), 0)

	if status.Err() != nil {
		return nil, status.Err()
	}

	return task, nil
}

func (dao *TaskDAORedis) Delete(id string) error {

	result, err := dao.redisCli.Del(id).Result()

	if err != nil {
		return err
	} else if result == 0 {
		return errors.New("REDIS -> " + id + " does not exist")
	}

	return nil
}
