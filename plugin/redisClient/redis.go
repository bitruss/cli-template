package redisClient

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/universe-30/CliAppTemplate/configuration"
)

var r *redis.ClusterClient

func GetSingleInstance() *redis.ClusterClient {
	return r
}

func Init() error {
	redis_addr, err := configuration.Config.GetString("redis_addr", "127.0.0.1")
	if err != nil {
		return errors.New("redis_addr [string] in config err," + err.Error())
	}

	redis_username, err := configuration.Config.GetString("redis_username", "")
	if err != nil {
		return errors.New("redis_username [string] in config err," + err.Error())
	}

	redis_password, err := configuration.Config.GetString("redis_password", "")
	if err != nil {
		return errors.New("redis_password [string] in config err," + err.Error())
	}

	redis_port, err := configuration.Config.GetInt("redis_port", 6379)
	if err != nil {
		return errors.New("redis_port [int] in config err," + err.Error())
	}

	r = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{redis_addr + ":" + strconv.Itoa(redis_port)},
		Username: redis_username,
		Password: redis_password,
	})

	_, err = r.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}
