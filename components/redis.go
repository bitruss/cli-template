package components

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/universe-30/CliAppTemplate/cli"
)

func InitRedis() (*redis.ClusterClient, error) {

	redis_addr, _redis_addr_err := cli.Config.GetString("redis_addr", "127.0.0.1")
	if _redis_addr_err != nil {
		return nil, errors.New("redis_addr [string] in config.json err," + _redis_addr_err.Error())
	}

	redis_username, redis_username_err := cli.Config.GetString("redis_username", "")
	if redis_username_err != nil {
		return nil, errors.New("redis_username [string] in config.json err," + redis_username_err.Error())
	}

	redis_password, redis_password_err := cli.Config.GetString("redis_password", "")
	if redis_password_err != nil {
		return nil, errors.New("redis_password [string] in config.json err," + redis_password_err.Error())
	}

	redis_port, redis_port_err := cli.Config.GetInt("redis_port", 6379)
	if redis_port_err != nil {
		return nil, errors.New("redis_port [int] in config.json err," + redis_port_err.Error())
	}

	Redis := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{redis_addr + ":" + strconv.Itoa(redis_port)},
		Username: redis_username,
		Password: redis_password,
	})

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return Redis, nil
}
