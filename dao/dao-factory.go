package dao

import (
	"errors"

	redis "gopkg.in/redis.v5"
)

type DBType int

const (
	RedisDAO DBType = iota
	MockDAO
)

func GetDao(daoType DBType) (TaskDao, error) {
	switch daoType {
	case RedisDAO:
		redisCli := initRedis()
		return NewTaskDAORedis(redisCli), nil
	case MockDAO:
		return NewTaskDAOMock(), nil
	default:
		return nil, errors.New("unknown DAO type")
	}
}

func initRedis() *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisCli.Ping().Result()
	if err != nil {
		panic(err)
	}
	return redisCli
}
